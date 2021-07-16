package httpserver

type Interface interface {
	AddHttpHandler(req AddRequest) error
}