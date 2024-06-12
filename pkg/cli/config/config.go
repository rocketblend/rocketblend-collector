package config

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/rocketblend/rocketblend/pkg/runtime"
	"github.com/spf13/viper"
)

type (
	Collection struct {
		Name        string             `mapstructure:"name" validate:"required"`
		Description string             `mapstructure:"description"`
		Args        string             `mapstructure:"args"`
		Platforms   []runtime.Platform `mapstructure:"platforms" validate:"required"`
	}

	Collector struct {
		Parallelism int    `mapstructure:"parallelism"`
		Delay       string `mapstructure:"delay"`
		Agent       string `mapstructure:"agent"`
		Proxy       string `mapstructure:"proxy"`
	}

	Config struct {
		Library     string        `mapstructure:"library" validate:"required"`
		OutputDir   string        `mapstructure:"outputDir"`
		Collector   *Collector    `mapstructure:"collector"`
		Collections []*Collection `mapstructure:"collections" validate:"required"`
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
		return runtime.PlatformFromString(data.(string)), nil
	}
}

func Load() (config *Config, err error) {
	v := viper.New()

	v.SetDefault("collector.parallelism", 2)
	v.SetDefault("collector.delay", "10s")
	v.SetDefault("collector.agent", "random")
	v.SetDefault("collector.outputDir", ".")

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
