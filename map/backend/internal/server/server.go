package server

import (
	"fmt"
	"net/http"
	"server/internal/ais"
	"server/internal/config"
	"server/internal/mqtt"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"server/internal/handler"
)

func NewServer(c *config.Config, logger *zap.SugaredLogger) *http.Server {
	aisProcessor := ais.New(logger)
	mqttClient := mqtt.NewMQttClient(c, logger)

	aisHandler := handler.NewAISHAndler(aisProcessor, mqttClient)
	aisHandler.MQTT.Sub([]string{"test", "test2"})

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Handle(`metrics`, promhttp.Handler()).Methods(http.MethodGet)
	r.Handle(`/health`, handler.HealthHandler{}).Methods(http.MethodGet)
	r.Handle(`/api/ais`, aisHandler).Methods(http.MethodPost)

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", c.HttpHost(), c.HttpPort()),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
}
