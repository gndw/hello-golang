package model

import (
	"go.uber.org/dig"
	"go.uber.org/fx"
)

type App struct {
	Fx struct {
		Container *dig.Container
		Instance  *fx.App
		Options   []fx.Option
	}
	Log struct {
		UseDefaultLogger *NullableBoolean
		IsOverrideFxLogger *NullableBoolean
	}
}

func (a *App) Run() {
	a.Fx.Instance.Run()
}
