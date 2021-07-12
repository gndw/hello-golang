package main

import (
	"hello-golang/app"

	"go.uber.org/dig"
	"go.uber.org/fx"
)


func main() {

	container := dig.New()
	app.SetupProviders(container)

	appl := app.Application{}
	startups := appl.Startup()
	
	var invokes []fx.Option
	invokes = append(invokes, fx.Invoke(func (lc fx.Lifecycle)  {	
		container.Provide(func() fx.Lifecycle {
			return lc
		})
	}),)
	for _, s := range startups {
		invokes = append(invokes, fx.Invoke(func ()  {
			container.Invoke(s)
		}))
	}

	app := fx.New(invokes...)
	app.Run()
}
