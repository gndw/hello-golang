package authentication

import (
	"hello-golang/v1/functions/stderror"
	"hello-golang/v1/services/httpserver"
)


type Authentication struct {
}

func (x *Authentication) Startup(httpServer httpserver.IHttpServer) (err *stderror.Error) {
	return
}