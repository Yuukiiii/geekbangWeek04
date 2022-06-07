package configs

type Config struct {
	Port    int    `json:"port"`
	Message string `json:"message"`
}

var GlobalConfig = new(Config)
