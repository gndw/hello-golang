package helpers

import (
	"io"
	"io/ioutil"

	"net/http"
)

type GenericResponse struct{
	Data interface{} `json:"data"`
	Error string `json:"error"`
}

func ReadBodyFromHTTPRequest(r *http.Request) (body string, err error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func SendingBadRequestResponse(rw http.ResponseWriter, err error) {
	response := &GenericResponse{Error: err.Error()}
	rw.WriteHeader(http.StatusBadRequest)
	io.WriteString(rw, StructToSafeJSONString(response))
}

func SendingOKResponse(rw http.ResponseWriter, data interface{}) {
	response := &GenericResponse{Data: data}
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, StructToSafeJSONString(response))
}

func SendingUnauthorizedResponse(rw http.ResponseWriter, details error) {
	response := &GenericResponse{Error: "unauthorized"}
	rw.WriteHeader(http.StatusUnauthorized)
	io.WriteString(rw, StructToSafeJSONString(response))
}
