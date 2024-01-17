package config

import (
	"io"

	"gopkg.in/yaml.v2"
)

func LoadConfig(reader io.Reader) (*Config, error) {
	buf, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	cfg := Config{}
	if err := yaml.Unmarshal(buf, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
