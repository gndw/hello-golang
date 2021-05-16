package requests

import (
	gerrors "errors"
	"hello-golang/errors"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterData struct {
	Email string `json:"email"`
	LoginData
}

type LoginRequest struct {
	Data LoginData `json:"Data"`
}

type RegisterRequest struct {
	Data RegisterData `json:"Data"`
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

func GenerateValidationForRegisterData(registerdata *RegisterData) *[]Validation {

	// Here you can check and sanitize your request data

	loginDataValidations := GenerateValidationForLoginData(&registerdata.LoginData)
	registerDataValidations := []Validation{
		func() *errors.Error {
			if registerdata.Email == "" {
				return errors.CreateErrorFromPrimitiveError(gerrors.New("email must not be empty"))
			} else {
				return nil
			}
		}}

	x := append(registerDataValidations, *loginDataValidations...)
	return &x
}
