package startup

import "hello-golang/v1/functions/stderror"

type IStartup interface {
	Startup() *stderror.Error
}