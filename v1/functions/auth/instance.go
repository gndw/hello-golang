package auth

import "log"

type Function struct {
}

func GetFunction() (f *Function, err error) {
	auth := &Function{}
	return auth, nil
}

func (x *Function) Login(usrname string, pwd string) (err error) {
	log.Println("Login...")
	return nil
}