package helpers

import (
	"hello-golang/errors"
	"hello-golang/responses"
	"io"
	"io/ioutil"

	"net/http"
)

func ReadBodyFromHTTPRequest(r *http.Request) (string, *errors.Error) {
	bytes, read_error := ioutil.ReadAll(r.Body)
	if read_error != nil {
		return "", errors.CreateErrorFromPrimitiveError(read_error)
	}

	return string(bytes), nil
}

func SendingBadRequestResponse(rw http.ResponseWriter, err *errors.Error) {
	response := &responses.GenericResponse{Error: err}
	rw.WriteHeader(http.StatusBadRequest)
	io.WriteString(rw, StructToSafeJSONString(response))
}

func SendingOKResponse(rw http.ResponseWriter, data interface{}) {
	response := &responses.GenericResponse{Data: data}
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, StructToSafeJSONString(response))
}

func SendingUnauthorizedResponse(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusUnauthorized)
}
