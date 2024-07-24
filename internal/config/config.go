package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
    AppName string
    Port    string
    DB      DatabaseConfig
}

type DatabaseConfig struct {
    Host        string
    Port        string
    User        string
    Password    string
    Name        string
}

var AppConfig *Config

func LoadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("error reading config file, %s", err)
    }

    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("unable to decode into struct, %v", err)
    }
}
