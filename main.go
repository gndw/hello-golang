package main

import (
	"hello-golang/v1/domains/functions/auth"
	"hello-golang/v1/domains/functions/health"
	"hello-golang/v1/domains/handlers/httphandler"
	"hello-golang/v1/domains/startups/httpsups"
	"hello-golang/v1/functions/builder"
	"hello-golang/v1/functions/builder/fxb"
	"hello-golang/v1/packages/factory"
	"hello-golang/v1/services/http/server"
	"log"
)


func main() {

	app, err := builder.CreateApp(
		fxb.ProvideServices(
			factory.GetRouter(),
			server.GetService,
		),
		fxb.ProvideFunctions(
			auth.GetFunction,
			health.GetFunction,
		),
		fxb.ProvideHandlers(
			httphandler.GetAuthHandler,
			httphandler.GetHealthHandler,
		),
		fxb.ProvideStartups(
			httpsups.StartHealthSystem,
			httpsups.StartAuthSystem,
		),
	)

	if (err != nil) {
		log.Fatal(err)
	}

	app.Run()
}
