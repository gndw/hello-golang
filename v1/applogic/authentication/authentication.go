package authentication

import (
	"hello-golang/handlers"
	"hello-golang/v1/services/httpserver"
	"log"
)


type Authentication struct {
	HttpServer httpserver.IHttpServer
}

func New (httpserver httpserver.IHttpServer) (auth *Authentication, err error) {
	auth = &Authentication{}
	auth.HttpServer = httpserver

	log.Println("hello")

	auth.HttpServer.Router().HandleFunc("/", handlers.HealthHandler).Methods("GET")

	return auth, nil

}