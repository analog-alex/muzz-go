package server

import (
	"context"
	"fmt"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port      string `env:"PORT, default=8080"`
	JwtSecret string `env:"JWT_SECRET, default=ultra-secret-key"`
}

var config Config

func init() {
	config = Config{}
	ctx := context.Background()
	if err := envconfig.Process(ctx, &config); err != nil {
		panic(err)
	}

	fmt.Println(config.Port)
	fmt.Println("Config loaded")
}

func GetApplicationConfig() Config {
	return config
}
