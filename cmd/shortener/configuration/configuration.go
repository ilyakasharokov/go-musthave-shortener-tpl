// Конфигурация
package configuration

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

var paramNames = map[string]string{
	"BASE_URL":          "b",
	"SERVER_ADDRESS":    "a",
	"FILE_STORAGE_PATH": "f",
	"ENABLE_HTTPS":      "s",
	"CONFIG":            "c",
	"TRUSTED_SUBNET":    "t",
	"ENABLE_GRPC":       "g",
}

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:""`
	Database        string `env:"DATABASE_DSN"`
	EnableHTTPS     bool   `env:"ENABLE_HTTPS"`
	Config          string `env:"CONFIG"`
	TrustedSubnet   string `env:"TRUSTED_SUBNET"`
	EnableGRPC      bool   `env:"ENABLE_GRPC"`
}

func New() Config {
	// Parse environment
	var cEnv Config
	err := env.Parse(&cEnv)
	if err != nil {
		panic(err)
	}
	// file config
	var c Config
	fcfg := flag.String(paramNames["CONFIG"], "", "")
	if *fcfg != "" {
		cEnv.Config = *fcfg
	}
	if cEnv.Config != "" {
		c, _ = getConfigFromFIle(cEnv.Config)
	}

	c.EnableHTTPS = cEnv.EnableHTTPS
	c.EnableGRPC = cEnv.EnableGRPC
	if cEnv.Database != "" {
		c.Database = cEnv.Database
	}
	if c.ServerAddress == "" || cEnv.ServerAddress != ":8080" {
		c.ServerAddress = cEnv.ServerAddress
	}
	if cEnv.BaseURL != "" {
		c.BaseURL = cEnv.BaseURL
	}
	if cEnv.FileStoragePath != "" {
		c.FileStoragePath = cEnv.FileStoragePath
	}
	if cEnv.TrustedSubnet != "" {
		c.TrustedSubnet = cEnv.TrustedSubnet
	}
	bu := flag.String(paramNames["BASE_URL"], "", "")
	sa := flag.String(paramNames["SERVER_ADDRESS"], "", "")
	fs := flag.String(paramNames["FILE_STORAGE_PATH"], "", "")
	db := flag.String(paramNames["DATABASE_DSN"], "", "")
	tls := flag.Bool(paramNames["ENABLE_HTTPS"], false, "")
	ts := flag.String(paramNames["TRUSTED_SUBNET"], "", "")
	grpc := flag.Bool(paramNames["ENABLE_GRPC"], false, "")
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
	if *db != "" {
		c.Database = *db
	}
	if tls != nil {
		c.EnableHTTPS = *tls
	}
	if *ts != "" {
		c.TrustedSubnet = *ts
	}
	if grpc != nil {
		c.EnableHTTPS = *grpc
	}
	return c
}
