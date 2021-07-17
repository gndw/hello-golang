package builder

import (
	"go.uber.org/fx"
)

type App struct {
	providers []interface{}
	startups []interface{}
	instance *fx.App
}

func (a *App) Run () {
	a.instance.Run()
}

func CreateApp(options ...BuilderOption) (app *App, err error) {

	app = &App{}

	for _, option := range options {
		err = option(app)
		if err != nil {
			return
		}
	}

	fxapp, err := ConfigApp(app.providers, app.startups)
	if (err != nil) {
		return nil, err
	}

	app.instance = fxapp
	return app, nil
}

