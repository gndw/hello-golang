package gochi

import (
	"hello-golang/v1/services/http/model"
	"hello-golang/v1/services/http/router"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type Instance struct {
	router *chi.Mux
}

func GetHttpRouter(lc fx.Lifecycle) (router.Interface, error) {

	ins := &Instance{}
	ins.router = chi.NewRouter()

	return ins, nil
}

func (x *Instance) AddHttpHandler(req model.AddRequest) (err error) {
	x.router.MethodFunc(req.Method, req.Endpoint, req.Handler)
	return nil
}

func (x *Instance) GetHandler() (handler http.Handler, err error) {
	return x.router, nil
}