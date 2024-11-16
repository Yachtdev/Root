package handler

import (
	"net/http"
)

type (
	// HealthHandler реализует ответчик health-check.
	HealthHandler struct{}
)

func (hh HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"OK"}`))
}
