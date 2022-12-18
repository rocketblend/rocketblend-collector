package config

import (
	"github.com/go-playground/validator"
	"github.com/rocketblend/rocketblend/pkg/core/runtime"
	"github.com/spf13/viper"
)

type (
	Collection struct {
		Name        string             `yaml:"name" validate:"required"`
		Description string             `yaml:"description"`
		Args        string             `yaml:"args"`
		Includes    []string           `yaml:"includes"`
		Packages    []string           `yaml:"packages"`
		Platforms   []runtime.Platform `yaml:"platforms" validate:"required"`
	}

	Config struct {
		Library     string       `yaml:"library" validate:"required"`
		Collections []Collection `yaml:"collections" validate:"required"`
	}
)

func Load() (config *Config, err error) {
	viper.SetConfigType("yaml")      // Set the config type to YAML
	viper.SetConfigName("collector") // Set the name of the configuration file

	viper.AddConfigPath(".") // Look for the configuration file in the current

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, err
	}

	return config, err
}
