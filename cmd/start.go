package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	controllers "github.com/waymobetta/go-coindrop-api/controllers"
	"github.com/waymobetta/go-coindrop-api/db"
	routehandlers "github.com/waymobetta/go-coindrop-api/handlers"
	"github.com/waymobetta/go-coindrop-api/router"
)

func main() {
	cors := handlers.CORS(
		handlers.AllowedHeaders(
			[]string{"X-Requested-With", "Content-Type", "Authorization"},
		),
		handlers.AllowedMethods(
			[]string{"GET", "POST", "OPTIONS"},
		),
		handlers.AllowedOrigins([]string{"*"}),
	)

	dbport, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	dbs := db.NewDB(&db.Config{
		Host:    os.Getenv("POSTGRES_HOST"),
		Port:    dbport,
		User:    os.Getenv("POSTGRES_USER"),
		Pass:    os.Getenv("POSTGRES_PASS"),
		Dbname:  os.Getenv("POSTGRES_DBNAME"),
		SSLMode: os.Getenv("POSTGRES_SSL_MODE"),
	})

	hdlrs := routehandlers.NewHandlers(&routehandlers.Config{
		DB: dbs,
	})
	rtr := router.NewRouter(&router.Config{
		Region:     os.Getenv("AWS_COINDROP_COGNITO_REGION"),
		UserPoolID: os.Getenv("AWS_COINDROP_COGNITO_USER_POOL_ID"),
		Handlers:   hdlrs,
	})

	rootHandler := cors(rtr)

	port := "5000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	host := fmt.Sprintf("0.0.0.0:%s", port)
	//log.Fatal(http.ListenAndServe(host, rootHandler))

	// TODO: merge current routes with goa routes

	// Create service
	service := goa.New("coindrop")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "user" controller
	c := controllers.NewUserController(service, dbs)
	app.MountUserController(service, c)

	// goa handler
	goaHandler := service.Server.Handler

	rootMux := http.NewServeMux()

	// merge handlers
	rootMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		goaRoutesRegex := regexp.MustCompile(`v1/users`)
		// Update regex to include any base goa routes in order to properly forward to goa handler
		isGoaRoute := goaRoutesRegex.Match([]byte(r.URL.Path))

		if isGoaRoute {
			goaHandler.ServeHTTP(w, r)
		} else {
			rootHandler.ServeHTTP(w, r)
		}
	})

	log.Printf("[cmd] listening on %s\n", host)

	// Start service
	panic(http.ListenAndServe(host, rootMux))
}
