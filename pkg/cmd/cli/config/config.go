package config

import (
	"reflect"

	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"github.com/rocketblend/rocketblend/pkg/core/runtime"
	"github.com/spf13/viper"
)

type (
	Collection struct {
		Name        string             `mapstructure:"name" validate:"required"`
		Description string             `mapstructure:"description"`
		Args        string             `mapstructure:"args"`
		Includes    []string           `mapstructure:"includes"`
		Packages    []string           `mapstructure:"packages"`
		Platforms   []runtime.Platform `mapstructure:"platforms" validate:"required"`
	}

	Config struct {
		Library     string       `mapstructure:"library" validate:"required"`
		Collections []Collection `mapstructure:"collections" validate:"required"`
	}
)

func PlatformHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		// Check that the data is string
		if f.Kind() != reflect.String {
			return data, nil
		}

		// Check that the target type is our custom type
		if t != reflect.TypeOf(runtime.Platform(0)) {
			return data, nil
		}

		// Return the parsed value
		p := runtime.Platform(0)
		return p.FromString(data.(string)), nil
	}
}

func Load() (config *Config, err error) {
	viper.SetConfigType("yaml")      // Set the config type to YAML
	viper.SetConfigName("collector") // Set the name of the configuration file

	viper.AddConfigPath(".") // Look for the configuration file in the current

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config, viper.DecodeHook(PlatformHookFunc()))

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, err
	}

	return config, err
}
