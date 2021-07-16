package main

import (
	"hello-golang/v1/functions/auth"
	"hello-golang/v1/functions/health"
	"hello-golang/v1/handlers/httphandler"
	"hello-golang/v1/packages/fxwrap"
	"hello-golang/v1/packages/gochi"
	"hello-golang/v1/startups/httpsups"
)


func main() {

	app := fxwrap.Setup(
		[]interface{}{
			// Set all Providers Here
			// Set Functions
			auth.GetFunction,
			health.GetFunction,
			// Set Handlers
			httphandler.GetHealthHandler,
			httphandler.GetAuthHandler,
			// Set Services
			gochi.GetHttpServer,
		},
		[]interface{}{
			// Set all Startups Here
			httpsups.StartHealthSystem,
		},
	)

	app.Run()
}
