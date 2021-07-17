package builder

import (
	"go.uber.org/dig"
	"go.uber.org/fx"
)


func ConfigApp(providers []interface{}, startups []interface{}) (fxapp *fx.App, err error) {

	container := dig.New()

	var invokes []fx.Option
	invokes = append(invokes, fx.Invoke(func(lc fx.Lifecycle, sd fx.Shutdowner) {
		container.Provide(func() fx.Lifecycle {
			return lc
		})
		container.Provide(func() fx.Shutdowner {
			return sd
		})
	}))

	
	for _, f := range providers {
		container.Provide(f)
	}
	
	for _, f := range startups {
		invokes = append(invokes, fx.Invoke(func() {
			container.Invoke(f)
		}))
	}
	
	return fx.New(invokes...), nil
}