package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Node struct {
		Port int    `toml:"port"`
		Name string `toml:"name"`
	} `toml:"node"`
	Bootnodes []string `toml:"bootnodes"`
}

func LoadConfig() (*Config, error) {
	var path string = "./config/config.toml"
	var cfg Config

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("Config file not found: %s", err)
	}

	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
