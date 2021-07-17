package server

import (
	"context"
	"fmt"
	"hello-golang/v1/services/http/model"
	"hello-golang/v1/services/http/router"
	"net/http"

	"go.uber.org/fx"
)

type Service struct {
	server *http.Server
	router router.Interface
}

func GetService(lc fx.Lifecycle, shutdowner fx.Shutdowner, router router.Interface) (service *Service, err error) {
	
	handler, err := router.GetHandler()
	if (err != nil) {
		return
	}

	service = &Service{
		router: router,
		server: &http.Server{Addr: "0.0.0.0:3000", Handler: handler},
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := service.server.ListenAndServe()
				if err != nil && err != http.ErrServerClosed {
					fmt.Println("server failed to start. err: ", err)
					shutdowner.Shutdown()
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return service.server.Shutdown(ctx)
		},
	})

	return service, nil
}

func (s *Service) AddHttpHandler(req model.AddRequest) (err error) {
	return s.router.AddHttpHandler(req)
}