package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	POSTGRES_DB_URL   string `mapstructure:"POSTGRES_DB_CONNECTION"`
	SERVER_PORT       string `mapstructure:"SERVER_PORT"`
	REDIS_PASSWORD    string `mapstructure:"REDIS_PASSWORD"`
	REDIS_SERVER      string `mapstructure:"REDIS_SERVER"`
	REDIS_SERVER_PORT string `mapstructure:"REDIS_SERVER_PORT"`
	JWT_SECRET_KEY    string `mapstructure:"JWT_SECRET_KEY"`
	PASSWORD_KEY      string `mapstructure:"PASSWORD_KEY"`
	ACCESS_REGION     string `mapstructure:"ACCESS_REGION"`
	ACCESS_KEYS       string `mapstructure:"ACCESS_KEYS"`
	SECRET_ACESS_KEYS string `mapstructure:"SECRET_ACESS_KEYS"`
}

func LoadInitializer(path string) (cfg Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("keys")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	err = viper.Unmarshal(&cfg)

	return
}
