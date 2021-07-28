package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
}