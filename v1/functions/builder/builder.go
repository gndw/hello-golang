package builder

import (
	"hello-golang/v1/functions/builder/fxb"
	"hello-golang/v1/functions/builder/model"
)


func CreateApp(options ...model.BuilderOption) (app *model.App, err error) {

	app = &model.App{}

	// Create FX App is executed at the end
	options = append(options, fxb.CreateFxApp())

	for _, option := range options {
		err = option(app)
		if err != nil {
			return
		}
	}

	return app, nil
}
