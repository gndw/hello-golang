package fxb

import "hello-golang/v1/functions/builder/model"

// Useless function, but now is just for tidying up
func ProvideServices(fs ...interface{}) (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		app.Fx.Providers = append(app.Fx.Providers, fs...)
		return nil
	}
}

// Useless function, but now is just for tidying up
func ProvideFunctions(fs ...interface{}) (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		app.Fx.Providers = append(app.Fx.Providers, fs...)
		return nil
	}
}

// Useless function, but now is just for tidying up
func ProvideHandlers(fs ...interface{}) (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		app.Fx.Providers = append(app.Fx.Providers, fs...)
		return nil
	}
}

// Useless function, but now is just for tidying up
func ProvideStartups(fs ...interface{}) (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		app.Fx.Startups = append(app.Fx.Startups, fs...)
		return nil
	}
}