package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port uint
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

func InitConfiguration() *Config {
	var C *Config = &Config{}

	loadDefault()
	loadFile()
	
	viper.Unmarshal(C)
	return C
}

func loadDefault() {
	viper.SetDefault("server.port", 50051)
}

func loadFile() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}