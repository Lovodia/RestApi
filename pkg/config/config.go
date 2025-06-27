package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port string `yaml:"port" env:"SERVER_PORT"`
}

type LoggerConfig struct {
	Level string `yaml:"level" env:"LOGGER_LEVEL"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	Logger LoggerConfig `yaml:"logger"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// Переопределяем параметры из переменных окружения, если они указаны
	if port := os.Getenv("SERVER_PORT"); port != "" {
		cfg.Server.Port = port
	}
	if level := os.Getenv("LOGGER_LEVEL"); level != "" {
		cfg.Logger.Level = level
	}

	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
	if cfg.Logger.Level == "" {
		cfg.Logger.Level = "info"
	}

	return &cfg, nil
}
