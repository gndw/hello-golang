package health

type Function struct {
}

func GetFunction() (f *Function, err error) {
	health := &Function{}
	return health, nil
}

func (f *Function) GetServerHealth() (health *Server, err error) {
	return &Server{Status: true}, nil
}