package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	DefaultConfigPath   = "/etc/config/genshin-artifact-db/config.yaml"
	DefaultPort         = ":8080"
	DefaultDataFilePath = "/var/lib/genshin-artifact-db/artifacts.json"
)

type Config struct {
	Port         string `yaml:"port"`
	DataFilePath string `yaml:"data_file_path"`
}

func DefaultConfig() *Config {
	return &Config{
		Port:         DefaultPort,
		DataFilePath: DefaultDataFilePath,
	}
}

func LoadConfig(path string) (*Config, error) {
	cfg := DefaultConfig()

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	if cfg.Port == "" {
		cfg.Port = DefaultPort
	}
	if cfg.DataFilePath == "" {
		cfg.DataFilePath = DefaultDataFilePath
	}

	return cfg, nil
}
