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

		// Appending Lifecycle & Shutdowner from fx into our Container
		// to make it accessible to other providers
		app.Fx.Options = append(app.Fx.Options, fx.Invoke(func(lc fx.Lifecycle, sd fx.Shutdowner) {
			app.Fx.Container.Provide(func() fx.Lifecycle {
				return lc
			})
			app.Fx.Container.Provide(func() fx.Shutdowner {
				return sd
			})
		}))

		// Appending onerror object
		// This object handle functions that need to be done, incase of providing error / panic
		app.Fx.Container.Provide(func () *OnError {
			return &OnError{}
		})

		return nil

	}
}

func CreateFxApp() (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		app.Fx.Instance = fx.New(app.Fx.Options...)
		return nil
	}
}

type OnError struct {
	Funcs []func()
}

func (oe *OnError) Append (f func())  {
	oe.Funcs = append(oe.Funcs, f)
}

func (oe *OnError) Execute ()  {
	for _, fu := range oe.Funcs {
		fu()
	}
}