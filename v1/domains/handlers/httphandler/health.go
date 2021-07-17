package httphandler

import (
	"hello-golang/v1/domains/functions/health"
	"net/http"

	"github.com/go-chi/render"
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
		render.Status(r, 400)
		render.JSON(rw, r, GenericResponse{Error: err.Error()})
		return
	}

	render.JSON(rw, r, GenericResponse{Data: stat})
}