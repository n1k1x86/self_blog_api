package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BlogDBConfig `yaml:"db_config"`
}

type BlogDBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Server   string `yaml:"server"`
	DBName   string `yaml:"db_name"`
}

func ReadConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	yaml.Unmarshal(data, cfg)
	return cfg, nil
}
