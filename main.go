package main

import (
	"hello-golang/v1/domains/functions/auth"
	"hello-golang/v1/domains/functions/health"
	"hello-golang/v1/domains/handlers/httphandler"
	"hello-golang/v1/domains/startups/httpsups"
	"hello-golang/v1/functions/builder"
	"hello-golang/v1/packages/gochi"
	"os"
)


func main() {

	app, err := builder.CreateApp(
		builder.ProvideServices(
			gochi.GetHttpServer,
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
		os.Exit(1)
	}

	app.Run()
}
