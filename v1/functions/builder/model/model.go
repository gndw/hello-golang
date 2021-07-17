package model

import "hello-golang/v1/functions/builder/fxb/model"

type App struct {
	Fx model.FxModel
}

func (a *App) Run() {
	a.Fx.Instance.Run()
}

type BuilderOption func(*App) error