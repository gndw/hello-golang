package httphandler

import (
	"net/http"

	"github.com/go-chi/render"
)

type GenericRequest struct {
	Data interface{} `json:"data"`
}

type GenericResponse struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

// Function to Sending HTTP Response.
// Set this function to be executed at defer.
// Important notes : you need to pass data interface{} as POINTER of your actual data!
// Eventough your data is already a pointer lol
func SendingResponse(rw http.ResponseWriter, r *http.Request, err *error, data interface{}) (func()) {
	return func () {
		if err != nil && (*err) != nil {
			render.Status(r, 400)
			render.JSON(rw, r, GenericResponse{Error: (*err).Error()})
		} else {
			render.JSON(rw, r, GenericResponse{Data: data})
		}
	}
}