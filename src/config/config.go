package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Filepaths []string `json:"filepaths"`
}

func LoadConfig(path string) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

func SaveConfig(path string, config *Config) (*Config, error) {
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return nil, err
	}
	absPath, _ := filepath.Abs(path)
	fmt.Println("Saving to:", absPath)
	return config, os.WriteFile(path, b, 0644)
}
