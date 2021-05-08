package controllers

import (
	"encoding/json"
	"hello-golang/responses"
	"hello-golang/services/health"
	"io"
	"net/http"
)

func HealthController(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-type", "application/json")
	health, err := health.CheckHealth()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	response := responses.GenericResponse{Error: err, Data: health}

	hobj, _ := json.Marshal(response)
	var res string = string(hobj)

	io.WriteString(rw, res)
}
