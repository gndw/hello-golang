package model

import (
	"go.uber.org/dig"
	"go.uber.org/fx"
)


type FxModel struct {
	Container *dig.Container
	Instance  *fx.App
	Options  []fx.Option
}