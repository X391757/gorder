package config

import "github.com/spf13/viper"

func NewViperConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("../common/config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}
