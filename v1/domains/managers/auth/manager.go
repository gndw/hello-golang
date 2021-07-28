package auth

import "hello-golang/v1/services/log"

type Manager struct {
	log log.Interface
}

func GetManager(log log.Interface) (f *Manager, err error) {
	auth := &Manager{
		log: log,
	}
	return auth, nil
}

func (x *Manager) Login(req *LoginRequest) (response *LoginResponse, err error) {
	x.log.Warningf("Some user try to login with username %v and password %v", req.Username, req.Password)
	return &LoginResponse{
		Userid: "xxx",
		Username: req.Username,
	}, nil
}