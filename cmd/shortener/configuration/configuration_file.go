package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ConfigFile struct {
	ServerAddress   string `json:"server_address"`
	BaseURL         string `json:"base_url"`
	FileStoragePath string `json:"file_storage_path"`
	DatabaseDSN     string `json:"database_dsn"`
	EnableHTTPS     bool   `json:"enable_https"`
}

func getConfigFromFIle(fileName string) (Config, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return Config{}, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return Config{}, err
	}
	cfg := ConfigFile{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	return Config{
		ServerAddress: cfg.ServerAddress,
		BaseURL:       cfg.BaseURL,
		FileStoragePath:      cfg.FileStoragePath,
		EnableHTTPS:   cfg.EnableHTTPS,
		Database: cfg.DatabaseDSN,
	}, nil
}