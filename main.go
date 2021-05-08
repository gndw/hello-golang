package main

import (
	"hello-golang/middlewares"
	"log"
	"net/http"

	"hello-golang/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", controllers.HealthController)
	r.Use(middlewares.ContentTypeJSONMiddleware)
	log.Fatal(http.ListenAndServe(":3000", r))
}
