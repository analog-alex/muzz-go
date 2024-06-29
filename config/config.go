package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port       string `env:"PORT, default=8080"`
	JwtSecret  string `env:"JWT_SECRET, default=secret_signing_key"`
	DbHost     string `env:"DB_HOST, default=localhost"`
	DbPort     string `env:"DB_PORT, default=5432"`
	DbUser     string `env:"DB_USER, default=user"`
	DbPassword string `env:"DB_PASSWORD, default=password"`
	DbName     string `env:"DB_NAME, default=mzz"`
}

var config Config

func init() {
	config = Config{}
	ctx := context.Background()
	if err := envconfig.Process(ctx, &config); err != nil {
		panic(err)
	}
}

func GetApplicationConfig() Config {
	return config
}
