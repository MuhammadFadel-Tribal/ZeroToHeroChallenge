package config

import (
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"go.uber.org/zap"
)

// Configurations Application wide configurations
type Configurations struct {
	Database DatabaseConfigurations `koanf:"database"`
	Server   ServerConfigurations   `koanf:"server"`
	Kafka    KafkaConfigurations    `koanf:"kafka"`
}

type ServerConfigurations struct {
	Port int `koanf:"port"`
}

// DatabaseConfigurations database configurations
type DatabaseConfigurations struct {
	Dsn  string `koanf:"dsn"`
	Pool int    `koanf:"pool"`
}

// KafkaConfigurations Kafka general configurations
type KafkaConfigurations struct {
	Broker   string `koanf:"broker"`
	Topic    string `koanf:"topic"`
	MaxBytes int    `koanf:"maxBytes"`
}

// LoadConfig Loads configurations depending upon the environment
func LoadConfig(logger *zap.SugaredLogger) *Configurations {
	k := koanf.New(".")
	err := k.Load(file.Provider("resources/config.yml"), yaml.Parser())
	if err != nil {
		logger.Fatalf("Failed to locate configurations. %v", err)
	}

	// Searches for env variables and will transform them into koanf format
	// e.g. SERVER_PORT variable will be server.port: value
	err = k.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil)
	if err != nil {
		logger.Fatalf("Failed to replace environment variables. %v", err)
	}

	var configuration Configurations

	err = k.Unmarshal("", &configuration)
	if err != nil {
		logger.Fatalf("Failed to load configurations. %v", err)
	}

	return &configuration
}
