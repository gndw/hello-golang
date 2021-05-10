package middlewares

import (
	"context"
	"hello-golang/constants"
	"hello-golang/helpers"
	"net/http"
)

func ReadBodyFromHTTPRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		body_string, body_error := helpers.ReadBodyFromHTTPRequest(r)
		if body_error != nil {
			helpers.SendingBadRequestResponse(rw, body_error)
			return
		}

		ctx := context.WithValue(r.Context(), constants.Context_Body, body_string)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
