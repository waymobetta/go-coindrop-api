package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/waymobetta/go-coindrop-api/router"
)

func main() {
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router.NewRouter())))
}
