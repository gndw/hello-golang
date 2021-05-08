package responses

import (
	"hello-golang/errors"
)

type GenericResponse struct {
	Error *errors.Error `json:"error"`
	Data  interface{}   `json:"data"`
}
