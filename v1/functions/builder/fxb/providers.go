package fxb

import (
	"hello-golang/v1/functions/builder/model"
	"hello-golang/v1/services/log"

	glog "log"

	"go.uber.org/fx"
)

func ProvideGenerics(fs ...interface{}) (opt model.BuilderOption) {
	return func(app *model.App) (err error) {
		for _, f := range fs {
			err = app.Fx.Container.Provide(f)
			if (err != nil) {
				return
			}
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
			
			// this cmd line is used to store memory of f to be invoked in fx.Invoke
			// if not, then it will be replaced by next iteration
			savedFunc := f

			app.Fx.Options = append(app.Fx.Options, fx.Invoke(func (sd fx.Shutdowner) {

				defer func ()  {
					if r := recover(); r != nil {
						app.Fx.Container.Invoke(func (oe *OnError)  {
							oe.Execute()
						})
						logerr := app.Fx.Container.Invoke(func (logg log.Interface){ logg.Errorf("panic! : %s", r) })
						if (logerr != nil) {
							glog.Printf("panic! : %s", r)
						}
					}
				}()

				err := app.Fx.Container.Invoke(savedFunc)
				if (err != nil) {
					logerr := app.Fx.Container.Invoke(func (logg log.Interface){ logg.Errorf("failed to invoke startup. err: %s", err) })
					if (logerr != nil) {
						glog.Printf("failed to invoke startup. err: %s", err)
					}
					app.Fx.Container.Invoke(func (oe *OnError)  {
						oe.Execute()
					})
					sd.Shutdown()
				}
			}))
		}
		return nil
	}
}