package model

import "go.uber.org/fx"

type FxModel struct {
	Providers []interface{}
	Startups  []interface{}
	Instance  *fx.App
}