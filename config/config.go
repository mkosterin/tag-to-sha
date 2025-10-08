package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Registries map[string]RegistryConfig `yaml:",inline"`
}

type RegistryConfig struct {
	AuthReq bool   `yaml:"authReq"`
	AuthUrl string `yaml:"authUrl,omitempty"`
	Token   string `yaml:"token,omitempty"`
}

func NewConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
