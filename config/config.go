package config

import (
	"fmt"
	"log"
	"os"

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

var cfg Config

func Load() Config {
	if cfg.Database.Type != "" {
		return cfg // cache
	}

	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("config.yaml açılamadı: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("config.yaml parse edilemedi: %v", err)
	}

	return cfg
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
