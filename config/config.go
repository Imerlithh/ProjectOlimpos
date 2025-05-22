package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Type     string `yaml:"type"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Extra    string `yaml:"extra"`
	} `yaml:"database"`
}

var Current Config

func Load() (Config, error) {
	configPath := filepath.Join("config", "config.yaml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("config.yaml okunamadı: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("config.yaml parse edilemedi: %w", err)
	}

	return cfg, nil
}

func BuildConnectionString(cfg Config) string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?%s",
		cfg.Database.Type,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.Extra,
	)
}
