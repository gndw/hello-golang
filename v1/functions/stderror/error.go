package stderror

import (
	"bytes"
)

type Error struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Details []*Error `json:"details"`
}

func (e *Error) Error() string {
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
