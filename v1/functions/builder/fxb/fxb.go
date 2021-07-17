package fxb

import (
	"hello-golang/v1/functions/builder/model"

	"go.uber.org/dig"
	"go.uber.org/fx"
)


func CreateFxApp() (opt model.BuilderOption) {

	return func(app *model.App) (err error) {

		container := dig.New()

		var invokes []fx.Option
		invokes = append(invokes, fx.Invoke(func(lc fx.Lifecycle, sd fx.Shutdowner) {
			container.Provide(func() fx.Lifecycle {
				return lc
			})
			container.Provide(func() fx.Shutdowner {
				return sd
			})
		}))

		
		for _, f := range app.Fx.Providers {
			container.Provide(f)
		}
		
		for _, f := range app.Fx.Startups {
			invokes = append(invokes, fx.Invoke(func() {
				container.Invoke(f)
			}))
		}
		
		app.Fx.Instance = fx.New(invokes...)
		return nil
	}
}