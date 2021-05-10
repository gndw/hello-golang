package auth

import (
	"hello-golang/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Token string `json:"token"`
}

type UserClaims struct {
	Userid string `json:"userid"`
	jwt.StandardClaims
}

func Login(username string, password string) (*Auth, *errors.Error) {

	if username != "username" || password != "password" {
		return nil, errors.CreateError(errors.WrongUserAndPassword)
	}

	claims := &UserClaims{
		Userid: "123",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Unix() + 15000,
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, signed_error := unsignedToken.SignedString([]byte("secret"))
	if signed_error != nil {
		return nil, errors.CreateErrorFromPrimitiveError(signed_error)
	}

	return &Auth{Token: signedToken}, nil
}
