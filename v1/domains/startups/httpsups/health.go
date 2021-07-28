package httpsups

import (
	"hello-golang/v1/domains/handlers/httphandler"
	"hello-golang/v1/services/http/model"
	"hello-golang/v1/services/http/server"
	"hello-golang/v1/services/log"
)


func StartHealthSystem(log log.Interface, server *server.Service, handler *httphandler.HealthHandler) {

	log.Infoln("starting Health System...")

	server.AddHttpHandler(
		model.AddRequest{
			Method: "GET",
			Endpoint: "/",
			Handler: handler.ServerHealthHandler,
		},
	)
		
}