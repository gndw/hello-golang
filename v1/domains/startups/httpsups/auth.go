package httpsups

import (
	"hello-golang/v1/domains/handlers/httphandler"
	"hello-golang/v1/services/http/model"
	"hello-golang/v1/services/http/server"
	"hello-golang/v1/services/log"
)


func StartAuthSystem(log log.Interface, server *server.Service, handler *httphandler.AuthHandler) {

	log.Infoln("starting Auth System...")

	server.AddHttpHandler(
		model.AddRequest{
			Method: "POST",
			Endpoint: "/login",
			Handler: handler.LoginHandler,
		},
	)
		
}