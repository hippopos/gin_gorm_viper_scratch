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

	c.Port = viper.GetInt("port")
	c.PC.Host = viper.GetString("postgres.host")
	c.PC.Port = viper.GetInt("postgres.port")
	c.PC.User = viper.GetString("postgres.user")
	c.PC.Password = viper.GetString("postgres.password")
	c.PC.Database = viper.GetString("postgres.database")
	c.PC.SSL = viper.GetString("postgres.ssl")
}
