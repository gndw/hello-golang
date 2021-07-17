package main

import (
	"hello-golang/v1/domains/functions/auth"
	"hello-golang/v1/domains/functions/health"
	"hello-golang/v1/domains/handlers/httphandler"
	"hello-golang/v1/domains/startups/httpsups"
	"hello-golang/v1/functions/builder"
	"hello-golang/v1/packages/gochi"
	"hello-golang/v1/services/http/server"
	"log"
)


func main() {

	app, err := builder.CreateApp(
		builder.ProvideServices(
			gochi.GetHttpRouter,
			server.GetService,
		),
		builder.ProvideFunctions(
			auth.GetFunction,
			health.GetFunction,
		),
		builder.ProvideHandlers(
			httphandler.GetAuthHandler,
			httphandler.GetHealthHandler,
		),
		builder.ProvideStartups(
			httpsups.StartHealthSystem,
		),
	)

	if (err != nil) {
		log.Fatal(err)
	}

	app.Run()
}
