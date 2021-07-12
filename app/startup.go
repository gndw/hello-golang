package app

import (
	"hello-golang/v1/applogic/authentication"
)

type Application struct {
	Authentication authentication.Authentication
}

func (a *Application) Startup () (funcs []interface{}) {
	return []interface{}{
		authentication.New,
	}

}