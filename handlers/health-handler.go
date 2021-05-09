package handlers

import (
	"hello-golang/helpers"
	"hello-golang/services/health"
	"net/http"
)

func HealthHandler(rw http.ResponseWriter, r *http.Request) {

	health, health_err := health.CheckHealth()
	if health_err != nil {
		helpers.SendingBadRequestResponse(rw, health_err)
		return
	} else {
		helpers.SendingOKRequestResponse(rw, health)
		return
	}
}
