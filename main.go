package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var Router = mux.NewRouter().StrictSlash(true)

func main() {
	log.Println("Started!")
	corsRouter := cors.AllowAll().Handler(Router)
	loggedRouter := handlers.LoggingHandler(os.Stdout, corsRouter)
	log.Fatal(http.ListenAndServe(":3684", loggedRouter))
}
