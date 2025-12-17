package serviceutils

type LogConfig struct {
	Level string `json:"level"`
}

type ServiceConfig struct {
	Log         LogConfig `json:"log"`
	ServiceName string    `json:"service_name"`
}
