package controllers

import (
	"hello-golang/helpers"
	"hello-golang/responses"
	"hello-golang/services/health"
	"io"
	"net/http"
)

func HealthController(rw http.ResponseWriter, r *http.Request) {

	response := &responses.GenericResponse{}

	health, health_err := health.CheckHealth()
	if health_err == nil {
		response.Data = health
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		response.Error = health_err
	}

	str := helpers.StructToSafeJSONString(response)
	io.WriteString(rw, str)

}
