package errors

import "fmt"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	InternalServerError = 10000
)

func CreateError(code int) *Error {
	switch code {
	case InternalServerError:
		return &Error{Code: fmt.Sprint(InternalServerError), Message: "Internal Server Error"}
	default:
		return &Error{Code: "Unknown Error", Message: "Unknown Error"}
	}
}
