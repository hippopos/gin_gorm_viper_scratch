package config

import "github.com/spf13/viper"

type Config struct {
	Port int
	PC   PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string

	SSL string `json:"ssl"`
}

func NewConfig() Config {

	c := new(Config)
	c.init()
	return *c
}

func (c *Config) init() {
	viper.UnmarshalKey("port", &c.Port)
	viper.UnmarshalKey("postgres", &c.PC)
}
