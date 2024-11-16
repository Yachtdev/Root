package config

import (
	"github.com/spf13/viper"
)

type (
	// Config struct with server confuguration.
	Config struct {
		// DB Database configuration.
		DB DBConfig `yaml:"db"`
		// Http API connect configuration.
		Http HttpConfig `yaml:"http"`
		Mqtt MqttConfig `yaml:"mqtt"`
	}

	// DBConfig struct with database configuration.
	DBConfig struct {
		// DSN Database dsn
		DSN string `yaml:"dsn"`
	}

	// HttpConfig API server connect configuration.
	HttpConfig struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}

	// MqttConfig конфигурация mqtt-сервера.
	MqttConfig struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		ClientID string `yaml:"client_id"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
)

// NewConfig constructor.
func NewConfig(viper *viper.Viper) *Config {
	return &Config{
		DB: DBConfig{
			DSN: viper.GetString("db.dsn"),
		},
		Http: HttpConfig{
			Port: viper.GetString("http.port"),
			Host: viper.GetString("http.host"),
		},
		Mqtt: MqttConfig{
			Host:     viper.GetString("mqtt.host"),
			Port:     viper.GetString("mqtt.port"),
			ClientID: viper.GetString("mqtt.client_id"),
			Username: viper.GetString("mqtt.username"),
			Password: viper.GetString("mqtt.password"),
		},
	}
}

// HttpPort возвращает порт, на котором отвечает сервер.
func (c *Config) HttpPort() string {
	return c.Http.Port
}

// HttpHost возвращает адрес, на котором будет развёрнут сервер.
func (c *Config) HttpHost() string {
	return c.Http.Host
}

// MqttHost возвращает адрес, на котором развёрнут mqtt-сервер.
func (c *Config) MqttHost() string {
	return c.Mqtt.Host
}

// MqttPort возвращает порт, на котором отвечает mqtt-сервер.
func (c *Config) MqttPort() string {
	return c.Mqtt.Port
}

// MqttClientID возвращает идентификатор клиента.
func (c *Config) MqttClientID() string {
	return c.Mqtt.ClientID
}

// MqttUsername возвращает логин пользователя.
func (c *Config) MqttUsername() string {
	return c.Mqtt.Username
}

// MqttPassword возвращает пароль пользователя.
func (c *Config) MqttPassword() string {
	return c.Mqtt.Password
}
