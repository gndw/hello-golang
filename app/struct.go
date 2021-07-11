package app

import (
	"hello-golang/v1/applogic/authentication"
	"hello-golang/v1/services/httpserver"
)


type App struct {
	Services Services
	AppLogic AppLogic
}

type Services struct {
	HttpServer httpserver.IHttpServer
}

type AppLogic struct {
	Authentication authentication.Authentication
}