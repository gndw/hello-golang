package http_gorilla_mux

import (
	"context"
	"hello-golang/v1/functions/stderror"
	"hello-golang/v1/services/httpserver"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

type HttpGorillaMux struct {
	R *mux.Router
}

func New(lc fx.Lifecycle) (httpserver.IHttpServer, error) {
	log.Println("createing router")
	server := &HttpGorillaMux{}
	server.R = mux.NewRouter()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			http.ListenAndServe(":3000", server.R)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return server, nil
}

func (x *HttpGorillaMux) AddHttpRouter() (err *stderror.Error) {
	log.Println("Adding Http Router ...")
	return nil
}

func (x *HttpGorillaMux) Router() (*mux.Router) {
	return x.R
}