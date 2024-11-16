package mqtt

import (
	"fmt"
	"go.uber.org/zap"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type (
	// Config конфигурация.
	Config interface {
		// MqttHost узел на котором развёрнут сервер.
		MqttHost() string
		// MqttPort порт на котором отвечает сервер.
		MqttPort() string
		// MqttClientID идентификатор клиента.
		MqttClientID() string
		// MqttUsername логин пользователя.
		MqttUsername() string
		// MqttPassword пароль.
		MqttPassword() string
	}

	// MQTT реализует логику работы с mqtt-сервером.
	MQTT struct {
		Client             mqtt.Client
		logger             *zap.SugaredLogger
		messagePubHandler  mqtt.MessageHandler
		connectHandler     mqtt.OnConnectHandler
		connectLostHandler mqtt.ConnectionLostHandler
	}
)

// NewMQttClient constructor.
func NewMQttClient(cfg Config, logger *zap.SugaredLogger) *MQTT {
	Mqtt := &MQTT{
		messagePubHandler: func(client mqtt.Client, msg mqtt.Message) {
			logger.Infof("[yachtdev-map-server] Received message: %s from topic: %s", msg.Payload(), msg.Topic())
		},
		connectHandler: func(client mqtt.Client) {
			logger.Infof("[yachtdev-map-server] Connected to mqtt-broker")
		},
		connectLostHandler: func(client mqtt.Client, err error) {
			logger.Infof("[yachtdev-map-server] Connect lost to mqtt-broker: %v", err)
		},
		logger: logger,
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", cfg.MqttHost(), cfg.MqttPort()))
	opts.SetClientID(cfg.MqttClientID())
	opts.SetUsername(cfg.MqttUsername())
	opts.SetPassword(cfg.MqttPassword())
	opts.SetDefaultPublishHandler(Mqtt.messagePubHandler)
	opts.OnConnect = Mqtt.connectHandler
	opts.OnConnectionLost = Mqtt.connectLostHandler
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	Mqtt.Client = client

	return Mqtt
}

// Pub отправляет сообщения в топик mqtt.
func (m *MQTT) Pub(payload []byte, topic string) {
	token := m.Client.Publish(topic, 0, false, payload)
	token.Wait()

	m.logger.Infof("[yachtdev-map-server] Send to topic %s: %s", topic, payload)
}

// Sub Подписывает на сообщения из топиков mqtt.
func (m *MQTT) Sub(topics []string) {
	for _, topic := range topics {
		token := m.Client.Subscribe(topic, 1, nil)
		token.Wait()

		m.logger.Infof("[yachtdev-map-server] Subscribed to topic: %s", topic)
	}
}
