package builder

import (
	"hello-golang/v1/functions/builder/fxb"
	"hello-golang/v1/functions/builder/logb"
	"hello-golang/v1/functions/builder/model"
)


func CreateApp(bopts ...model.BuilderOption) (app *model.App, err error) {

	app = &model.App{}
	options := []model.BuilderOption {}

	// Create Container must be executed first
	options = append(options, fxb.CreateContainer())

	// Fill in options from Builder
	// All of this are user specs, inputted from main.go / app entry point
	options = append(options, bopts...)

	// Fill in Default Options
	// All default option will check app status before making any changes
	// to makesure not overriding user specific option
	options = append(options, logb.UseDefaultLogger())
	options = append(options, logb.OverrideFxLogger())

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
