package fxb

import (
	"hello-golang/v1/functions/builder/model"
	"hello-golang/v1/packages/logrus"
	"hello-golang/v1/services/log"

	glog "log"

	"go.uber.org/fx"
)

func ProvideGenerics(fs ...interface{}) (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		for _, f := range fs {
			app.Fx.Container.Provide(f)
		}
		return nil
	}
}

// Useless function, but now is just for tidying up
func ProvideServices(fs ...interface{}) (opt model.BuilderOption) {
	return ProvideGenerics(fs...)
}

// Useless function, but now is just for tidying up
func ProvideFunctions(fs ...interface{}) (opt model.BuilderOption) {
	return ProvideGenerics(fs...)
}

// Useless function, but now is just for tidying up
func ProvideHandlers(fs ...interface{}) (opt model.BuilderOption) {
	return ProvideGenerics(fs...)
}

func ProvideStartups(fs ...interface{}) (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		for _, f := range fs {
			app.Fx.Options = append(app.Fx.Options, fx.Invoke(func (sd fx.Shutdowner) {
				err := app.Fx.Container.Invoke(f)
				if (err != nil) {
					logerr := app.Fx.Container.Invoke(func (logg log.Interface){ logg.Errorf("failed to invoke startup. err: %s", err) })
					if (logerr != nil) {
						glog.Printf("failed to invoke startup. err: %s", err)
					}
					sd.Shutdown()
				}
			}))
		}
		return nil
	}
}

func OverrideFxLogger() (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		logrus, err := logrus.GetLog()
		if (err != nil) {
			return
		}
		app.Fx.Options = append(app.Fx.Options, fx.Logger(logrus))
		return nil
	}
}

func UseDefaultLogger() (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		app.Fx.Container.Provide(logrus.GetLog)
		return nil
	}
}