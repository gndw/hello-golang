package router

import (
	"hello-golang/v1/services/http/model"
	"net/http"
)


type Interface interface {
	AddHttpHandler(req model.AddRequest) (err error)
	GetHandler() (handler http.Handler, err error)
}