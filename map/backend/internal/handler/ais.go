package handler

import (
	"encoding/json"
	"net/http"
)

type (
	AIS interface {
		Encode(lat, lon, cog float64) []string
		Decode(msg []byte)
	}

	MQTT interface {
		Pub(payload []byte, topic string)
		Sub(topics []string)
	}

	AISHandler struct {
		AIS  AIS
		MQTT MQTT
	}

	AISRequest struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
		Cog float64 `json:"cog"`
	}
)

func NewAISHAndler(aisProcessor AIS, mqttClient MQTT) AISHandler {
	return AISHandler{
		AIS:  aisProcessor,
		MQTT: mqttClient,
	}
}

func (ah AISHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ar AISRequest

	err := json.NewDecoder(r.Body).Decode(&ar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	strs := ah.AIS.Encode(ar.Lat, ar.Lon, ar.Cog)

	for _, str := range strs {
		ah.MQTT.Pub([]byte(str), "test")
	}

	w.Header().Set("Access-Control-Allow-Methods", http.MethodPost)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusCreated)
}
