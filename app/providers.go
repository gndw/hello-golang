package app

import (
	"hello-golang/v1/repositories/http_gorilla_mux"

	"go.uber.org/dig"
)


func SetupProviders (container *dig.Container) {
	container.Provide(http_gorilla_mux.New)
}