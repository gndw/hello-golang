package http_gorilla_mux

import (
	"hello-golang/v1/functions/stderror"
	"log"

	"github.com/gorilla/mux"
)

type HttpGorillaMux struct {
	Router *mux.Router
}

func (x *HttpGorillaMux) Startup() (err *stderror.Error) {
	log.Println("Starting Gorilla Mux ...")
	x.Router = mux.NewRouter()
	return
}