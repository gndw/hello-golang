package requests

import (
	gerrors "errors"
	"hello-golang/errors"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Data LoginData `json:"Data"`
}

func GenerateValidationForLoginData(logindata *LoginData) *[]Validation {

	// Here you can check and sanitize your request data

	return &[]Validation{
		func() *errors.Error {
			if logindata.Username == "" {
				return errors.CreateErrorFromPrimitiveError(gerrors.New("username must not be empty"))
			} else {
				return nil
			}
		},
		func() *errors.Error {
			if logindata.Password == "" {
				return errors.CreateErrorFromPrimitiveError(gerrors.New("password must not be empty"))
			} else {
				return nil
			}
		}}
}
