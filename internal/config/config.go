package config

import (
	"github.com/spf13/viper"
)

var AppConfig Config

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Port             string `mapstructure:"PORT"`
	Debug            bool   `mapstructure:"DEBUG"`
	BaseURL          string `mapstructure:"BASE_URL"`
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     string `mapstructure:"DB_PORT"`
	DatabaseUsername string `mapstructure:"DB_USERNAME"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
	DatabaseName     string `mapstructure:"DB_DATABASE"`
	DatabaseSSL      string `mapstructure:"DB_SSL"`
	JwtSecret        string `mapstructure:"JWT_SECRET_KEY"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (err error) {
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&AppConfig)

	return
}
