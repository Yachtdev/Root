package handler

import (
	"fmt"
	"net/http"
)

type HealthHandler struct{}

func (hh HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"status":"OK"}`)
}
