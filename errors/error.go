package errors

import (
	"bytes"
	"container/list"
	"fmt"
)

type Error struct {
	Code    string     `json:"code"`
	Message string     `json:"message"`
	Details *list.List `json:"details"`
}

func (e Error) Error() string {
	var buffer bytes.Buffer
	buffer.WriteString(e.Message)

	if e.Details != nil {
		for d := e.Details.Front(); d != nil; d = d.Next() {
			buffer.WriteString("\n")
			buffer.WriteString(d.Value.(error).Error())
		}
	}

	return buffer.String()
}

func (e *Error) AddDetail(de *Error) {
	if e.Details == nil {
		e.Details = list.New()
	}

	e.Details.PushBack(de)
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
