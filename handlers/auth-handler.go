package handlers

import (
	"hello-golang/helpers"
	"hello-golang/middlewares"
	"hello-golang/requests"
	"hello-golang/services/auth"
	"net/http"
)

func LoginHandler(rw http.ResponseWriter, r *http.Request) {

	request := &requests.LoginRequest{}
	structify_error := helpers.JSONStringToStruct(
		r.Context().Value(middlewares.Body).(string), request,
		requests.GenerateValidationForLoginData(&request.Data))
	if structify_error != nil {
		helpers.SendingBadRequestResponse(rw, structify_error)
		return
	}

	auth, auth_err := auth.Login(request.Data.Username, request.Data.Password)
	if auth_err != nil {
		helpers.SendingBadRequestResponse(rw, auth_err)
		return
	} else {
		helpers.SendingOKRequestResponse(rw, auth)
		return
	}

}
