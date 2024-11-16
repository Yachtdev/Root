package handler

import (
	"encoding/json"
	"net/http"
	"server/gen"
)

// AddAisPoint добавляет AIS-точку.
func (h *Handler) AddAisPoint(w http.ResponseWriter, r *http.Request) {
	var aisPoint gen.AisPoint
	if err := json.NewDecoder(r.Body).Decode(&aisPoint); err != nil {
		h.sendError(w, http.StatusBadRequest, "Invalid format for aisPoint")
		return
	}

	h.service.SendAisPointToMqtt(aisPoint.Lat, aisPoint.Lon, *aisPoint.Course)

	w.Header().Set("Access-Control-Allow-Methods", http.MethodPost)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(aisPoint)
}
