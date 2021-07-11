package app

import (
	"hello-golang/v1/applogic/authentication"
	"hello-golang/v1/repositories/http_gorilla_mux"
	"log"
)


func (a *App) New() {
	log.Println("Creating Application...")
	a.Services = Services{
		HttpServer: &http_gorilla_mux.HttpGorillaMux{},
	}
	a.AppLogic = AppLogic{
		Authentication: authentication.Authentication{},
	}

	a.Services.HttpServer.Startup()
	a.AppLogic.Authentication.Startup(a.Services.HttpServer)
}
