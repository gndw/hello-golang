package httphandler

import (
	"hello-golang/v1/domains/functions/health"
	"hello-golang/v1/helpers"
	"net/http"
)


type HealthHandler struct {
	f *health.Function
}

func GetHealthHandler(f *health.Function) (handler *HealthHandler, err error) {
	h := &HealthHandler{
		f: f,
	}
	return h, nil
}

func (h *HealthHandler) Handler(rw http.ResponseWriter, r *http.Request) {
	
	stat, err := h.f.Get()
	if (err != nil) {
		helpers.SendingBadRequestResponse(rw, err)
		return
	}

	helpers.SendingOKResponse(rw, stat)
}