package builder

import (
	fxfunc "hello-golang/v1/functions/fx"

	"go.uber.org/fx"
)

type App struct {
	Instance *fx.App
}

func (a *App) Run () {
	a.Instance.Run()
}

func CreateApp(options ...*BuilderOption) (app *App, err error) {

	app = &App{}

	var providers []interface{}
	var startups []interface{}

	for _, option := range options {
		switch option.optype {
		case FxIsProvider: {
			providers = append(providers, option.list...)
			break;
		}
		case FxIsStartup: {
			startups = append(startups, option.list...)
			break;
		}
		}
	}

	fxapp, err := fxfunc.ConfigApp(providers, startups)
	if (err != nil) {
		return nil, err
	}

	app.Instance = fxapp
	return app, nil
}

// Useless function, but now is just for tidying up
func ProvideServices(fs ...interface{}) (opt *BuilderOption) {

	return &BuilderOption{
		optype: FxIsProvider,
		list: fs,
	}

}

// Useless function, but now is just for tidying up
func ProvideFunctions(fs ...interface{}) (opt *BuilderOption) {

	return &BuilderOption{
		optype: FxIsProvider,
		list: fs,
	}

}

// Useless function, but now is just for tidying up
func ProvideHandlers(fs ...interface{}) (opt *BuilderOption) {

	return &BuilderOption{
		optype: FxIsProvider,
		list: fs,
	}

}

// Useless function, but now is just for tidying up
func ProvideStartups(fs ...interface{}) (opt *BuilderOption) {

	return &BuilderOption{
		optype: FxIsStartup,
		list: fs,
	}

}