package configuration

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

var paramNames = map[string]string{
	"BASE_URL":          "b",
	"SERVER_ADDRESS":    "a",
	"FILE_STORAGE_PATH": "f",
}

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:"http://localhost:8080"`
	BaseURL         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:""`
}

func New() Config {
	var c Config
	// Parse environment
	err := env.Parse(&c)
	if err != nil {
		panic(err)
	}
	bu := flag.String(paramNames["BASE_URL"], "", "")
	sa := flag.String(paramNames["SERVER_ADDRESS"], "", "")
	fs := flag.String(paramNames["FILE_STORAGE_PATH"], "", "")
	flag.Parse()
	if *bu != "" {
		c.BaseURL = *bu
	}
	if *sa != "" {
		c.ServerAddress = *sa
	}
	if *fs != "" {
		c.FileStoragePath = *fs
	}
	return c
}