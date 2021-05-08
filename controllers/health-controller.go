package controllers

import (
	"hello-golang/helpers"
	"hello-golang/responses"
	"hello-golang/services/health"
	"io"
	"net/http"
)

func HealthController(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	health, health_err := health.CheckHealth()
	if health_err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	response := &responses.GenericResponse{Error: health_err, Data: health}

	finalResponse, fr_err := helpers.StructToJSON(response)
	if fr_err != nil {
		io.WriteString(rw, fr_err.Error())
	} else {
		io.WriteString(rw, finalResponse)
	}
}
