package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUrl      string
	ServerPort string
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig() // ignore if not found

	return &Config{
		DBUrl:      viper.GetString("DATABASE_URL"),
		ServerPort: viper.GetString("SERVER_PORT"),
	}, nil
}
