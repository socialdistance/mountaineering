package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	HTTP    HttpConf
	Storage StorageConf
	Logger  LoggerConf
	Fs      FileServerConf
}

type FileServerConf struct {
	Host string `json:"host_fs"`
	Port string `json:"port_fs"`
}

type HttpConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type StorageConf struct {
	Dsn string `json:"dsn"`
}

type LoggerConf struct{}

func NewConfig() Config {
	return Config{}
}

func LoadConfig(path string) (*Config, error) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("invalid config file %s: %w", path, err)
	}

	config := NewConfig()
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("invalid unmarshall config %w", err)
	}

	return &config, nil
}
