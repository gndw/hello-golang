package httphandler

type AuthHandler struct {
}

func GetAuthHandler() (auth *AuthHandler, err error) {
	h := &AuthHandler{}
	return h, nil
}