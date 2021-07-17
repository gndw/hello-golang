package logb

import (
	"hello-golang/v1/functions/builder/model"
	"hello-golang/v1/packages/logrus"
	"hello-golang/v1/services/log"

	glog "log"

	"go.uber.org/fx"
)

func DisableOverrideFxLogger() (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		app.Log.IsOverrideFxLogger = &model.NullableBoolean{
			Value: false,
		}
		return nil
	}
}

func OverrideFxLogger() (opt model.BuilderOption) {
	return func(app *model.App) (err error) {

		if app.Log.IsOverrideFxLogger != nil && !app.Log.IsOverrideFxLogger.Value {
			dbgstr := "skip default override fx logger"
			logerr := app.Fx.Container.Invoke(func (logg log.Interface){ logg.Debugln(dbgstr) })
			if (logerr != nil) {
				glog.Println(dbgstr)
			}
			return
		}

		logrus, err := logrus.GetLog()
		if err != nil {
			return
		}
		app.Fx.Options = append(app.Fx.Options, fx.Logger(logrus))
		return nil
	}
}

func DisableDefaultLogger() (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		app.Log.UseDefaultLogger = &model.NullableBoolean{
			Value: false,
		}
		return nil
	}
}

func UseDefaultLogger() (opt model.BuilderOption) {
	return func(app *model.App) (err error) {

		if app.Log.UseDefaultLogger != nil && !app.Log.UseDefaultLogger.Value {
			glog.Println("skip using default logger")
			return
		}

		app.Fx.Container.Provide(logrus.GetLog)
		return nil
	}
}