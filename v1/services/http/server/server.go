package server

import (
	"context"
	"fmt"
	"hello-golang/v1/services/http/model"
	"hello-golang/v1/services/http/router"
	"hello-golang/v1/services/log"
	"net/http"

	"go.uber.org/fx"
)

type Service struct {
	server *http.Server
	router router.Interface
}

func GetService(lc fx.Lifecycle, shutdowner fx.Shutdowner, router router.Interface, log log.Interface) (service *Service, err error) {
	
	handler, err := router.GetHandler()
	if (err != nil) {
		return
	}

	port := 3000

	service = &Service{
		router: router,
		server: &http.Server{Addr: fmt.Sprintf("0.0.0.0:%v", port) , Handler: handler},
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) (err error) {
			go func() {
				err := service.server.ListenAndServe()
				if err != nil && err != http.ErrServerClosed {
					log.Errorf("server failed to start. err: ", err)
					shutdowner.Shutdown()
				}
				}()
			log.Infof("server is started in port %v...", port)
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