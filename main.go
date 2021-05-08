package main

import (
	"log"
	"net/http"

	"hello-golang/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", controllers.HealthController)
	log.Fatal(http.ListenAndServe(":3000", r))
}
