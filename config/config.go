package config

type Config struct {
	DB   DBConfig
	HTTP HTTPConfig
}

func LoadConfig() *Config {
	return &Config{
		DB:   LoadDBConfig(),
		HTTP: LoadHTTPConfig(),
	}
}
