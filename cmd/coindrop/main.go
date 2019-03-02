package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
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
	routehandlers "github.com/waymobetta/go-coindrop-api/handlers"
	mw "github.com/waymobetta/go-coindrop-api/middleware"
	"github.com/waymobetta/go-coindrop-api/router"
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

	hdlrs := routehandlers.NewHandlers(&routehandlers.Config{
		DB: dbs,
	})
	rtr := router.NewRouter(&router.Config{
		Region:     cognitoRegion,
		UserPoolID: cognitoUserPoolID,
		Handlers:   hdlrs,
	})

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

	rootHandler := c.Handler(rtr)

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
	})))

	// Mount controllers
	healthcheckCtrlr := controllers.NewHealthcheckController(service, dbs)
	app.MountHealthcheckController(service, healthcheckCtrlr)

	userCtrlr := controllers.NewUserController(service, dbs)
	app.MountUserController(service, userCtrlr)

	walletCtrlr := controllers.NewWalletController(service, dbs)
	app.MountWalletController(service, walletCtrlr)

	tasksCtrlr := controllers.NewTasksController(service, dbs)
	app.MountTasksController(service, tasksCtrlr)

	resultsCtrlr := controllers.NewResultsController(service, dbs)
	app.MountResultsController(service, resultsCtrlr)

	quizCtrlr := controllers.NewQuizController(service, dbs)
	app.MountQuizController(service, quizCtrlr)

	redditCtrlr := controllers.NewRedditController(service, dbs)
	app.MountRedditController(service, redditCtrlr)

	verifyRedditCtrlr := controllers.NewVerifyredditController(service, dbs)
	app.MountVerifyredditController(service, verifyRedditCtrlr)

	// goa handler
	goaHandler := c.Handler(mw.RateLimitHandler(service.Server.Handler))

	rootMux := http.NewServeMux()

	// NOTE: merging current routes with goa routes
	rootMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// limit payload sizes
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		// Update regex to include any base goa routes in order to properly forward to goa handler
		goaRoutesRegex := regexp.MustCompile(`v1/(health|users|wallets|tasks|quiz|results|social|verify)`)
		isGoaRoute := goaRoutesRegex.Match([]byte(strings.ToLower(r.URL.Path)))

		if isGoaRoute {
			goaHandler.ServeHTTP(w, r)
		} else if strings.HasPrefix(r.URL.Path, "/documentation") {
			http.FileServer(http.Dir("./web")).ServeHTTP(w, r)
		} else {
			rootHandler.ServeHTTP(w, r)
		}
	})

	log.Printf("[cmd] listening on %s\n", host)

	// Start service
	panic(http.ListenAndServe(host, rootMux))
}
