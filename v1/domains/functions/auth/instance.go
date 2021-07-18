package auth

import "hello-golang/v1/services/log"

type Function struct {
	log log.Interface
}

func GetFunction(log log.Interface) (f *Function, err error) {
	auth := &Function{
		log: log,
	}
	return auth, nil
}

func (x *Function) Login(req *LoginRequest) (response *LoginResponse, err error) {
	x.log.Warningf("Some user try to login with username %v and password %v", req.Username, req.Password)
	return &LoginResponse{
		Userid: "xxx",
		Username: req.Username,
	}, nil
}