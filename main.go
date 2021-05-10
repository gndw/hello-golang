package main

import (
	"hello-golang/middlewares"
	"log"
	"net/http"

	"hello-golang/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	r.HandleFunc("/auth/login", handlers.LoginHandler).Methods("POST")

	dataroute := r.PathPrefix("/data").Subrouter()
	dataroute.Use(middlewares.AuthorizationMiddleware)
	dataroute.HandleFunc("/self", handlers.DataSelfHandler).Methods("POST")

	r.Use(middlewares.ContentTypeJSONMiddleware)
	r.Use(middlewares.ReadBodyFromHTTPRequestMiddleware)

	log.Fatal(http.ListenAndServe(":3000", r))
}
