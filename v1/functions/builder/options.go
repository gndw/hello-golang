package builder

type BuilderOption func(*App) error

// Useless function, but now is just for tidying up
func ProvideServices(fs ...interface{}) (opt BuilderOption) {
	return func(app *App) (err error) {
		app.providers = append(app.providers, fs...)
		return nil
	}
}

// Useless function, but now is just for tidying up
func ProvideFunctions(fs ...interface{}) (opt BuilderOption) {
	return func(app *App) (err error) {
		app.providers = append(app.providers, fs...)
		return nil
	}
}

// Useless function, but now is just for tidying up
func ProvideHandlers(fs ...interface{}) (opt BuilderOption) {
	return func(app *App) (err error) {
		app.providers = append(app.providers, fs...)
		return nil
	}
}

// Useless function, but now is just for tidying up
func ProvideStartups(fs ...interface{}) (opt BuilderOption) {
	return func(app *App) (err error) {
		app.startups = append(app.startups, fs...)
		return nil
	}
}