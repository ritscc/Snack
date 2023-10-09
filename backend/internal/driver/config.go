package driver

import "github.com/caarlos0/env/v9"

type Config struct {
	ServerConfig ServerConfig
}

type ServerConfig struct {
	Host string `env:"SERVER_HOST" envDefault:"localhost"`
	Port int    `env:"SERVER_PORT" envDefault:"8080"`
}

func LoadConfig(config *Config) error {
	if err := env.Parse(config); err != nil {
		return err
	}

	return nil
}
