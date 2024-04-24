package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	RedisConfig  RedisConfig
	ServerConfig ServerConfig
	MongoConfig  MongoConfig
	TokenConfig  TokenConfig
}

type TokenConfig struct {
	Secret string
}

type ServerConfig struct {
	Port int
	Host string
}

type RedisConfig struct {
	RedisAddr string
	Password  string
	DB        int
}

type MongoConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
