package config

import (
	"github.com/spf13/viper"
)

const PageSize = 30
const configPath = "./config.yaml"

type Config struct {
	Logger   LoggerConfig   `yaml:"logger"`
	HTTP     HTTPConfig     `yaml:"http"`
	Database PostgresConfig `yaml:"database"`
	Jwt      Jwt            `yaml:"jwt"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type HTTPConfig struct {
	Port int `yaml:"port"`
}

type PostgresConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Jwt struct {
	Key string `yaml:"key"`
}

func ReadConfig() (*Config, error) {
	var config Config
	viper.SetConfigFile("../../config.local.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
