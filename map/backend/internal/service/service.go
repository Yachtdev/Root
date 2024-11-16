package service

import (
	"go.uber.org/zap"
)

type (
	// AIS производит кодирование и декодирование строк AIS.
	AIS interface {
		Encode(lat, lon, cog float32) []string
		Decode(msg []byte)
	}

	// MQTT MQTT-client.
	MQTT interface {
		Pub(payload []byte, topic string)
		Sub(topics []string)
	}

	// Service реализует сервисный слой.
	Service struct {
		AIS    AIS
		MQTT   MQTT
		logger *zap.SugaredLogger
	}
)

// New Service-constructor.
func New(ais AIS, mqtt MQTT, l *zap.SugaredLogger) *Service {
	return &Service{
		AIS:    ais,
		MQTT:   mqtt,
		logger: l,
	}
}
