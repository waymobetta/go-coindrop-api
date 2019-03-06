package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	auth "github.com/waymobetta/go-coindrop-api/auth"
	controllers "github.com/waymobetta/go-coindrop-api/controllers"
	"github.com/waymobetta/go-coindrop-api/db"
	mw "github.com/waymobetta/go-coindrop-api/middleware"
)

func main() {
	log.SetReportCaller(true)

	dbport := 5000
	var err error
	if os.Getenv("POSTGRES_PORT") != "" {
		dbport, err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		if err != nil {
			log.Fatal(err)
		}
	}

	dbs := db.NewDB(&db.Config{
		Host:    os.Getenv("POSTGRES_HOST"),
		Port:    dbport,
		User:    os.Getenv("POSTGRES_USER"),
		Pass:    os.Getenv("POSTGRES_PASS"),
		Dbname:  os.Getenv("POSTGRES_DBNAME"),
		SSLMode: os.Getenv("POSTGRES_SSL_MODE"),
	})

	cognitoRegion := os.Getenv("AWS_COINDROP_COGNITO_REGION")
	cognitoUserPoolID := os.Getenv("AWS_COINDROP_COGNITO_USER_POOL_ID")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{
			"Range",
			"Accept",
			"Content-Type",
			"Authorization",
			"X-CA-Session",
			"X-Requested-With",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"UPDATE",
			"DELETE",
			"PATCH",
			"OPTIONS",
			"HEAD",
		},
		AllowCredentials: true,
		Debug:            true,
	})

	port := "5000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	host := fmt.Sprintf("0.0.0.0:%s", port)

	// Create goa service
	service := goa.New("coindrop")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	app.UseJWTAuthMiddleware(service, mw.Auth(auth.NewAuth(&auth.Config{
		CognitoRegion:     cognitoRegion,
		CognitoUserPoolID: cognitoUserPoolID,
	}), dbs))

	// Mount controllers
	healthcheckCtrlr := controllers.NewHealthcheckController(service, dbs)
	app.MountHealthcheckController(service, healthcheckCtrlr)

	usersCtrlr := controllers.NewUsersController(service, dbs)
	app.MountUsersController(service, usersCtrlr)

	walletsCtrlr := controllers.NewWalletsController(service, dbs)
	app.MountWalletsController(service, walletsCtrlr)

	tasksCtrlr := controllers.NewTasksController(service, dbs)
	app.MountTasksController(service, tasksCtrlr)

	resultsCtrlr := controllers.NewResultsController(service, dbs)
	app.MountResultsController(service, resultsCtrlr)

	quizzesCtrlr := controllers.NewQuizzesController(service, dbs)
	app.MountQuizzesController(service, quizzesCtrlr)

	redditCtrlr := controllers.NewRedditController(service, dbs)
	app.MountRedditController(service, redditCtrlr)

	verifyRedditCtrlr := controllers.NewVerifyredditController(service, dbs)
	app.MountVerifyredditController(service, verifyRedditCtrlr)

	redditHarvestCtrlr := controllers.NewRedditharvestController(service, dbs)
	app.MountRedditharvestController(service, redditHarvestCtrlr)

	webhooksCtrlr := controllers.NewWebhooksController(service, dbs)
	app.MountWebhooksController(service, webhooksCtrlr)

	// goa handler
	goaHandler := c.Handler(mw.RateLimitHandler(service.Server.Handler))

	rootMux := http.NewServeMux()

	rootMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// limit payload sizes
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		if strings.HasPrefix(r.URL.Path, "/documentation") {
			http.FileServer(http.Dir("./web")).ServeHTTP(w, r)
		} else {
			goaHandler.ServeHTTP(w, r)
		}
	})

	log.Printf("[cmd] listening on %s\n", host)

	// Start service
	panic(http.ListenAndServe(host, rootMux))
}
