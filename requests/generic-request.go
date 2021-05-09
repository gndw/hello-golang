package requests

import (
	"hello-golang/errors"
)

type Validation func() *errors.Error

type GenericRequest struct {
	Data interface{} `json:"data"`
}
