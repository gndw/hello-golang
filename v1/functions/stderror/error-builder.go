package stderror

import "fmt"

const (
	InternalServerError = 10000
	FailedToParseToJSON = 10001
	AuthorizationFailed = 10002

	WrongUserAndPassword = 11000
)

func CreateError(code int) *Error {
	switch code {
	case InternalServerError:
		return &Error{Code: fmt.Sprint(InternalServerError), Message: "Internal Server Error"}
	case FailedToParseToJSON:
		return &Error{Code: fmt.Sprint(FailedToParseToJSON), Message: "Object is failed to parse to JSON. Contact Developer!"}
	case AuthorizationFailed:
		return &Error{Code: fmt.Sprint(AuthorizationFailed), Message: "Token Authorization Failed"}
	case WrongUserAndPassword:
		return &Error{Code: fmt.Sprint(WrongUserAndPassword), Message: "Wrong Username or Password"}
	default:
		return &Error{Code: "Unknown Error", Message: "Unknown Error"}
	}
}

func CreateErrorFromPrimitiveError(e error) *Error {
	return &Error{Code: "primitive", Message: e.Error()}
}
