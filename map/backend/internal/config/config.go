package config

import "github.com/spf13/viper"

type (
	// Config struct with server confuguration
	Config struct {
		// DB Database configuration
		DB DBConfig `yaml:"db"`
		// Http API connect configuration
		Http HttpConfig `yaml:"http"`
		Mqtt MqttConfig `yaml:"mqtt"`
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

	MqttConfig struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		ClientID string `yaml:"client_id"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
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
		Mqtt: MqttConfig{
			Host:     viper.GetString("mqtt.host"),
			Port:     viper.GetString("mqtt.port"),
			ClientID: viper.GetString("mqtt.client_id"),
			Username: viper.GetString("mqtt.username"),
			Password: viper.GetString("mqtt.password"),
		},
	}
}

func (c *Config) HttpPort() string {
	return c.Http.Port
}

func (c *Config) HttpHost() string {
	return c.Http.Host
}

func (c *Config) MqttHost() string {
	return c.Mqtt.Host
}

func (c *Config) MqttPort() string {
	return c.Mqtt.Port
}

func (c *Config) MqttClientID() string {
	return c.Mqtt.ClientID
}

func (c *Config) MqttUsername() string {
	return c.Mqtt.Username
}

func (c *Config) MqttPassword() string {
	return c.Mqtt.Password
}
