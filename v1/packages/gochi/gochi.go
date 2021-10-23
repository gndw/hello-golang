package gochi

import (
	"hello-golang/v1/services/http/model"
	"hello-golang/v1/services/http/router"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Instance struct {
	router *chi.Mux
}

func GetRouter() (router.Interface, error) {

	ins := &Instance{}
	ins.router = chi.NewRouter()
	ins.router.Use(middleware.Recoverer)

	return ins, nil
}

func (x *Instance) AddHttpHandler(req model.AddRequest) (err error) {
	x.router.MethodFunc(req.Method, req.Endpoint, req.Handler)
	return nil
}

func (x *Instance) GetHandler() (handler http.Handler, err error) {
	return x.router, nil
}