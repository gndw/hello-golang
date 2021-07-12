package httpserver

import (
	"hello-golang/v1/functions/stderror"

	"github.com/gorilla/mux"
)


type IHttpServer interface {
	AddHttpRouter() *stderror.Error
	Router() *mux.Router
}