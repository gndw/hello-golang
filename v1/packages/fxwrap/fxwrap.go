package fxwrap

import (
	"go.uber.org/dig"
	"go.uber.org/fx"
)

func Setup(providers []interface{}, startups []interface{}) (app *fx.App) {
	container := dig.New()
	for _, p := range providers {
		container.Provide(p)
	}

	var invokes []fx.Option
	invokes = append(invokes, fx.Invoke(func(lc fx.Lifecycle) {
		container.Provide(func() fx.Lifecycle {
			return lc
		})
	}))
	for _, s := range startups {
		invokes = append(invokes, fx.Invoke(func() {
			container.Invoke(s)
		}))
	}

	app = fx.New(invokes...)
	return app
}