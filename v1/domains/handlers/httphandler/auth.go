package httphandler

import (
	"hello-golang/v1/domains/managers/auth"
	"net/http"

	"github.com/go-chi/render"
)

type AuthHandler struct {
	manager *auth.Manager
}

func GetAuthHandler(manager *auth.Manager) (auth *AuthHandler, err error) {
	h := &AuthHandler{
		manager: manager,
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

	loginResponse, err = h.manager.Login(loginRequest)
}