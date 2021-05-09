package middlewares

import (
	"context"
	"hello-golang/helpers"
	"net/http"
)

const (
	Body key = iota
)

func ReadBodyFromHTTPRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		body_string, body_error := helpers.ReadBodyFromHTTPRequest(r)
		if body_error != nil {
			helpers.SendingBadRequestResponse(rw, body_error)
			return
		}

		ctx := context.WithValue(r.Context(), Body, body_string)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
