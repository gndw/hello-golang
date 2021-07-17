package fxb

import (
	"hello-golang/v1/functions/builder/model"

	"go.uber.org/dig"
	"go.uber.org/fx"
)

func CreateContainer() (opt model.BuilderOption) {
	return func(app *model.App) (err error) {

		// Dependencies Injection Container
		// Hold all provider funcs
		app.Fx.Container = dig.New()

		// Appending Lifecycle & Shutdowner into Container
		// to make it accessible to other providers
		app.Fx.Options = append(app.Fx.Options, fx.Invoke(func(lc fx.Lifecycle, sd fx.Shutdowner) {
			app.Fx.Container.Provide(func() fx.Lifecycle {
				return lc
			})
			app.Fx.Container.Provide(func() fx.Shutdowner {
				return sd
			})
		}))

		return nil

	}
}

func CreateFxApp() (opt model.BuilderOption) {

	return func(app *model.App) (err error) {

		// Create the App
		app.Fx.Instance = fx.New(app.Fx.Options...)
		return nil
	}
}