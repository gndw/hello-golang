package middlewares

import (
	"context"
	"hello-golang/constants"
	"hello-golang/helpers"
	"hello-golang/services/auth"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authstring := r.Header.Get("Authorization")
		runes := []rune(authstring)
		if string(runes[:7]) != "Bearer " {
			helpers.SendingUnauthorizedResponse(rw)
			return
		}

		tokenstring := string(runes[7:])
		token, parse_error := jwt.ParseWithClaims(
			tokenstring,
			&auth.UserClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			},
		)
		if parse_error != nil {
			helpers.SendingUnauthorizedResponse(rw)
			return
		}

		claims, ok := token.Claims.(*auth.UserClaims)
		if !ok {
			helpers.SendingUnauthorizedResponse(rw)
			return
		}

		if claims.ExpiresAt < time.Now().UTC().Unix() {
			helpers.SendingUnauthorizedResponse(rw)
			return
		}

		ctx := context.WithValue(r.Context(), constants.Context_Userid, claims.Userid)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
