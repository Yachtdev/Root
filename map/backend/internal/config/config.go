package config

import "github.com/spf13/viper"

type (
	// Config struct with server confuguration
	Config struct {
		// DB Database configuration
		DB DBConfig `yaml:"db"`
		// Http API connect configuration
		Http HttpConfig `yaml:"http"`
	}

	// DBConfig struct with database configuration
	DBConfig struct {
		// DSN Database dsn
		DSN string `yaml:"dsn"`
	}

	// HttpConfig API server connect configuration
	HttpConfig struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}
)

func NewConfig(viper *viper.Viper) *Config {
	return &Config{
		DB: DBConfig{
			DSN: viper.GetString("db.dsn"),
		},
		Http: HttpConfig{
			Port: viper.GetString("http.port"),
			Host: viper.GetString("http.host"),
		},
	}
}

func (c *Config) HttpPort() string {
	return c.Http.Port
}

func (c *Config) HttpHost() string {
	return c.Http.Host
}
