package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"server/internal/handler"
)

type (
	Config interface {
		HttpPort() string
		HttpHost() string
	}
)

func NewServer(c Config) *http.Server {
	r := mux.NewRouter()
	r.Handle(`metrics`, promhttp.Handler())
	r.Handle(`/health`, handler.HealthHandler{})

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", c.HttpHost(), c.HttpPort()),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
}
