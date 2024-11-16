package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"server/gen"
	"server/internal/config"
	"server/internal/handler"
	"server/internal/service"
)

// NewServer server-constructor.
func NewServer(c *config.Config, logger *zap.SugaredLogger, service *service.Service) *http.Server {
	swagger, err := gen.GetSwagger()
	if err != nil {
		panic(fmt.Errorf("failed to get swagger: %s", err))
	}

	swagger.Servers = nil

	handlers := handler.New(service)

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Handle(`/metrics`, promhttp.Handler()).Methods(http.MethodGet)
	r.Handle(`/health`, handler.HealthHandler{}).Methods(http.MethodGet)

	gen.HandlerFromMux(handlers, r)

	return &http.Server{
		Addr:         net.JoinHostPort(c.HttpHost(), c.HttpPort()),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
}
