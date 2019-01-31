package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
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
	log.Printf("[cmd] listening on %s\n", host)
	log.Fatal(http.ListenAndServe(host, rootHandler))
}
