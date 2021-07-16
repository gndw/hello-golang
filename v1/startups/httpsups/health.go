package httpsups

import (
	"hello-golang/v1/handlers/httphandler"
	"hello-golang/v1/services/httpserver"
	"log"
)


func StartHealthSystem(server httpserver.Interface, handler *httphandler.HealthHandler) {

	log.Println("Starting Health System...")

	server.AddHttpHandler(
		httpserver.AddRequest{
			Method: "GET",
			Endpoint: "/",
			Handler: handler.Handler,
		},
	)
		
}