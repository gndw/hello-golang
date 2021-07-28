package httphandler

import (
	"hello-golang/v1/domains/managers/health"
	"net/http"
)


type HealthHandler struct {
	manager *health.Manager
}

func GetHealthHandler(manager *health.Manager) (handler *HealthHandler, err error) {
	h := &HealthHandler{
		manager: manager,
	}
	return h, nil
}

func (h *HealthHandler) ServerHealthHandler(rw http.ResponseWriter, r *http.Request) {
	
	var (
		err error
		serverHealthResponse = &health.ServerHealthResponse{}
	)

	defer SendingResponse(rw, r, &err, &serverHealthResponse) ()
	serverHealthResponse, err = h.manager.GetServerHealth()
}