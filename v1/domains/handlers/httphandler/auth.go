package httphandler

import (
	"hello-golang/v1/domains/functions/auth"
	"net/http"

	"github.com/go-chi/render"
)

type AuthHandler struct {
	f *auth.Function
}

func GetAuthHandler(f *auth.Function) (auth *AuthHandler, err error) {
	h := &AuthHandler{
		f: f,
	}
	return h, nil
}

func (h *AuthHandler) LoginHandler(rw http.ResponseWriter, r *http.Request) {

	var (
		err error
		loginRequest = &auth.LoginRequest{}
		loginResponse = &auth.LoginResponse{}
	)

	defer SendingResponse(rw, r, &err, &loginResponse) ()
	
	err = render.DecodeJSON(r.Body, &GenericRequest{ Data: loginRequest })
	if err != nil {
		return
	}

	loginResponse, err = h.f.Login(loginRequest)
}