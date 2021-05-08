package health

import (
	"hello-golang/errors"
)

type Health struct {
	ServerStatus bool `json:"server_status"`
}

func CheckHealth() (*Health, *errors.Error) {
	return &Health{ServerStatus: true}, nil
}
