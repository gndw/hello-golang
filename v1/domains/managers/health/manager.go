package health

type Manager struct {
}

func GetManager() (f *Manager, err error) {
	health := &Manager{}
	return health, nil
}

func (f *Manager) GetServerHealth() (health *ServerHealthResponse, err error) {
	return &ServerHealthResponse{Status: true}, nil
}