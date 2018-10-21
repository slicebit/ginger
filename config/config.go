package config

import (
	"github.com/spf13/viper"
)

// New creates an instance of config given path
func New(cfgFile string, config interface{}) error {
	viper.SetConfigFile(cfgFile)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(config)
	return err
}
