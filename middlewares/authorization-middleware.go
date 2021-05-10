package middlewares

import (
	"context"
	"hello-golang/constants"
	"hello-golang/helpers"
	"hello-golang/services/auth"
	"net/http"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authstring := r.Header.Get("Authorization")
		runes := []rune(authstring)
		if string(runes[:7]) != "Bearer " {
			helpers.SendingUnauthorizedResponse(rw)
			return
		}

		token := string(runes[7:])
		claims, validate_error := auth.ValidateToken(token)
		if validate_error != nil {
			helpers.SendingUnauthorizedResponse(rw, validate_error)
			return
		}

		ctx := context.WithValue(r.Context(), constants.Context_Userid, claims.Userid)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
