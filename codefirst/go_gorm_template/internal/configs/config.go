package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ExamplePort string
	DBDriver    string
	DNS         string
}

func LoadConfig() (*Config, error) {
	// load from env
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".") // look for .env in current dir
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println(".env file not found")
		} else {
			return nil, fmt.Errorf("failed to read .env file: %w", err)
		}
	}
	viper.AutomaticEnv()

	var config Config
	config.ExamplePort = viper.GetString("EXAMPLE_PORT")
	config.DBDriver = viper.GetString("DB_DRIVER")
	config.DNS = viper.GetString("DB_DSN")

	return &config, nil
}
