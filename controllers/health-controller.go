package controllers

import (
	"hello-golang/helpers"
	"hello-golang/responses"
	"hello-golang/services/health"
	"io"
	"net/http"
)

func HealthController(rw http.ResponseWriter, r *http.Request) {

	health, health_err := health.CheckHealth()
	if health_err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	response := &responses.GenericResponse{Error: health_err, Data: health}

	finalResponse := helpers.StructToSafeJSONString(response)
	io.WriteString(rw, finalResponse)

}
