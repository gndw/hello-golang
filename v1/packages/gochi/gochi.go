package gochi

import (
	"context"
	"hello-golang/v1/services/httpserver"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type Instance struct {
	router *chi.Mux
	server *http.Server
}

func GetHttpServer(lc fx.Lifecycle) (httpserver.Interface, error) {

	ins := &Instance{}
	ins.router = chi.NewRouter()
	ins.server = &http.Server{Addr: "0.0.0.0:3000", Handler: ins.router}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go ins.server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return ins, nil
}

func (x *Instance) AddHttpHandler(req httpserver.AddRequest) (err error) {

	switch req.Method {
	case "GET":
		x.router.Get(req.Endpoint, req.Handler)
	}
	return nil
}