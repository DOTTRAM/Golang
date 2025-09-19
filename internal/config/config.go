package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPort string `env:"HTTP_PORT" envDefault:"8080"`

	DB struct {
		Host     string `env:"HOST" envDefault:"localhost"`
		Port     int    `env:"PORT" envDefault:"5432"`
		User     string `env:"USER" envDefault:"admin"`                                // Измените на admin
		Password string `env:"PASSWORD" envDefault:"1wFmQXTBXb/riob7SVbFGlVr1k0ZyC/V"` // Ваш пароль
		Name     string `env:"NAME" envDefault:"main"`
	} `envPrefix:"DB_"`
}

func Load() (*Config, error) {
	if _, err := os.Stat("./.env"); err == nil {
		if err := godotenv.Load("./.env"); err != nil {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("error parsing environment variables: %w", err)
	}
	return cfg, nil
}
