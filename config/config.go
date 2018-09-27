package config

import (
	"github.com/spf13/viper"
)

const configType string = "toml"
const configName string = "config"

// New creates an instance of config given path
func New(cfgFile string) (*viper.Viper, error) {

	config := viper.New()
	config.SetConfigFile(cfgFile)

	err := config.ReadInConfig()
	return config, err
}
