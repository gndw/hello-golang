package httphandler

type GenericResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}