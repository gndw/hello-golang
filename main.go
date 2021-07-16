package main

import (
	"hello-golang/v1/domains/functions/auth"
	"hello-golang/v1/domains/functions/health"
	"hello-golang/v1/domains/handlers/httphandler"
	"hello-golang/v1/domains/startups/httpsups"
	"hello-golang/v1/packages/fxwrap"
	"hello-golang/v1/packages/gochi"
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
