package model

import "net/http"

type AddRequest struct {
	Method   string
	Endpoint string
	Handler  func(rw http.ResponseWriter, r *http.Request)
}