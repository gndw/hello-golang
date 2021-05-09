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

	return &Auth{Token: "token-example"}, nil
}
