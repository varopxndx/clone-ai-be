package config

// Configuration struct
type Configuration struct {
	Port        string `mapstructure:"port"`
	OpenAIToken string `mapstructure:"open_api_key"`
}
