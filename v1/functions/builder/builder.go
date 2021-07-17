package builder

import (
	"hello-golang/v1/functions/builder/fxb"
	"hello-golang/v1/functions/builder/model"
)


func CreateApp(bopts ...model.BuilderOption) (app *model.App, err error) {

	app = &model.App{}
	options := []model.BuilderOption {}

	// Create Container must be executed first
	options = append(options, fxb.CreateContainer())

	// Fill in options from Builder
	options = append(options, bopts...)

	// Fill in Default Options
	options = append(options, fxb.OverrideFxLogger())
	options = append(options, fxb.UseDefaultLogger())

	// Create FX App must be executed last
	options = append(options, fxb.CreateFxApp())

	for _, option := range options {
		err = option(app)
		if err != nil {
			return
		}
	}

	return app, nil
}
