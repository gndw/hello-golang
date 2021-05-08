package health

type Health struct {
	ServerStatus bool `json:"server_status"`
}

func CheckHealth() *Health {
	return &Health{ServerStatus: true}
}
