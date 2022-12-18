package config

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"github.com/rocketblend/rocketblend/pkg/core/runtime"
	"github.com/spf13/viper"
)

type (
	Collection struct {
		Name        string             `mapstructure:"name" validate:"required"`
		Description string             `mapstructure:"description" default:""`
		Args        string             `mapstructure:"args" default:""`
		Packages    []string           `mapstructure:"packages"`
		Platforms   []runtime.Platform `mapstructure:"platforms" validate:"required"`
	}

	Collector struct {
		Parallelism int    `mapstructure:"parallelism" default:"2"`
		Delay       string `mapstructure:"delay" default:"5s"`
		Agent       string `mapstructure:"agent" default:"random"`
		Proxy       string `mapstructure:"proxy"`
	}

	Config struct {
		Library     string       `mapstructure:"library" validate:"required"`
		Collector   Collector    `mapstructure:"collector"`
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
	v := viper.New()

	v.SetConfigName("collector") // Set the name of the configuration file
	v.AddConfigPath(".")         // Look for the configuration file in the current
	v.SetConfigType("yml")       // Set the config type to YAML

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.BindEnv("collector.proxy")

	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(&config, viper.DecodeHook(PlatformHookFunc()))

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, err
	}

	return config, err
}
