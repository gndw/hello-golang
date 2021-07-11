package httpserver

import "hello-golang/v1/functions/stderror"

type IHttpServer interface {
	Startup() *stderror.Error
}