package config

import "github.com/spf13/viper"

type (
	Collection struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Includes    []string `json:"includes"`
		Platforms   []string `json:"platforms"`
		Packages    []string `json:"packages"`
	}

	Config struct {
		Library     string       `json:"library"`
		Collections []Collection `json:"collections"`
	}
)

func Load() (config *Config, err error) {
	viper.SetConfigType("yaml")      // Set the config type to YAML
	viper.SetConfigName("collector") // Set the name of the configuration file

	viper.AddConfigPath(".") // Look for the configuration file in the current

	viper.AutomaticEnv()

	// err = viper.ReadInConfig()
	// if err != nil {
	// 	return nil, err
	// }

	err = viper.Unmarshal(&config)
	return config, err
}
