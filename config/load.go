package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Load pulls the config data from the config file
func Load() (*Configuration, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error loading config file: %s", err)
		}
		// No config file found. Using environment variables
		viper.BindEnv("PORT")
		for _, key := range viper.AllKeys() {
			val := viper.Get(key)
			viper.Set(key, val)
		}
	}

	config := Configuration{}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unmarshalling config file: %w", err)
	}

	validator := validator.New()
	if err := validator.Struct(config); err != nil {
		return nil, fmt.Errorf("invalid config file: %w", err)
	}

	return &config, nil
}
