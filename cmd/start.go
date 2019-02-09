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
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	controllers "github.com/waymobetta/go-coindrop-api/controllers"
	"github.com/waymobetta/go-coindrop-api/db"
	routehandlers "github.com/waymobetta/go-coindrop-api/handlers"
	mw "github.com/waymobetta/go-coindrop-api/middleware"
	"github.com/waymobetta/go-coindrop-api/router"
)

func main() {
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

	rootHandler := handlers.CORS(
		handlers.AllowedHeaders(
			[]string{"X-Requested-With", "Content-Type", "Authorization"},
		),
		handlers.AllowedMethods(
			[]string{"GET", "POST", "OPTIONS"},
		),
		handlers.AllowedOrigins([]string{"*"}),
	)(rtr)

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

	// Mount "user" controller
	userCtrlr := controllers.NewUserController(service, dbs)
	app.MountUserController(service, userCtrlr)

	walletCtrlr := controllers.NewWalletController(service, dbs)
	app.MountWalletController(service, walletCtrlr)

	tasksCtrlr := controllers.NewTasksController(service, dbs)
	app.MountTasksController(service, tasksCtrlr)

	// goa handler
	goaHandler := mw.RateLimitHandler(service.Server.Handler)

	rootMux := http.NewServeMux()

	// NOTE: merging current routes with goa routes
	rootMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// limit payload sizes
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		goaRoutesRegex := regexp.MustCompile(`v1/(users|wallets|tasks)`)
		// Update regex to include any base goa routes in order to properly forward to goa handler
		isGoaRoute := goaRoutesRegex.Match([]byte(r.URL.Path))

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
	panic(http.ListenAndServe(host, cors.Default().Handler(rootMux)))
}
