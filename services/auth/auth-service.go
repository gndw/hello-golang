package auth

import (
	"hello-golang/errors"
)

type Auth struct {
	Token string `json:"token"`
}

func Login(username string, password string) (*Auth, *errors.Error) {

	if username != "username" || password != "password" {
		return nil, errors.CreateError(errors.WrongUserAndPassword)
	}

	signedToken, token_error := GenerateToken("123")
	if token_error != nil {
		return nil, token_error
	}

	return &Auth{Token: signedToken}, nil
}
