package controllers

import (
	"encoding/json"
	"hello-golang/services/health"
	"io"
	"net/http"
)

type HealthResponse struct {
}

func HealthController(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-type", "application/json")
	h := health.CheckHealth()
	hobj, _ := json.Marshal(h)
	io.WriteString(rw, string(hobj))
}
