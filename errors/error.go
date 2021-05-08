package errors

import (
	"bytes"
	"fmt"
)

type Error struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Details []*Error `json:"details"`
}

func (e Error) Error() string {
	var buffer bytes.Buffer
	buffer.WriteString(e.Message)

	if e.Details != nil {
		buffer.WriteString(" >> details:")
		for _, value := range e.Details {
			buffer.WriteString(" | ")
			buffer.WriteString(value.Error())
		}
	}

	return buffer.String()
}

func (e *Error) AddDetail(de *Error) {

	if e.Details == nil {
		e.Details = make([]*Error, 1)
		e.Details[0] = de
	} else {
		e.Details = append(e.Details, de)
	}

}

const (
	InternalServerError = 10000
	FailedToParseToJSON = 10001
)

func CreateError(code int) *Error {
	switch code {
	case InternalServerError:
		return &Error{Code: fmt.Sprint(InternalServerError), Message: "Internal Server Error"}
	case FailedToParseToJSON:
		return &Error{Code: fmt.Sprint(FailedToParseToJSON), Message: "Object is failed to parse to JSON. Contact Developer!"}
	default:
		return &Error{Code: "Unknown Error", Message: "Unknown Error"}
	}
}

func CreateErrorFromPrimitiveError(e error) *Error {
	return &Error{Code: "primitive", Message: e.Error()}
}
