package httpsups

import (
	"hello-golang/v1/domains/handlers/httphandler"
	"hello-golang/v1/services/http/model"
	"hello-golang/v1/services/http/server"
	"log"
)


func StartHealthSystem(server *server.Service, handler *httphandler.HealthHandler) {

	log.Println("Starting Health System...")

	server.AddHttpHandler(
		model.AddRequest{
			Method: "GET",
			Endpoint: "/",
			Handler: handler.Handler,
		},
	)
		
}