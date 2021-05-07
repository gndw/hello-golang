package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HealthHandler(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "health")
}
