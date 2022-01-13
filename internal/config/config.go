package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

const (
	DefaultConfigFile = "config.json"
)

// Open returns the configuration
func Open(configFileName string) (*Config, error) {

	bytearray, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	err = json.Unmarshal(bytearray, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
