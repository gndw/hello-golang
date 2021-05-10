package handlers

import (
	"hello-golang/constants"
	"hello-golang/helpers"
	"net/http"
)

func DataSelfHandler(rw http.ResponseWriter, r *http.Request) {

	userid := r.Context().Value(constants.Context_Userid).(string)
	helpers.SendingOKResponse(rw, userid)

}
