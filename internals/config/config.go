package config

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	IP   string
	Port string
}

func NewConfig() *Config {
	serverConfig := ServerConfig{
		IP:   "127.0.0.1",
		Port: "8000",
	}

	config := &Config{
		Server: serverConfig,
	}

	return config
}
