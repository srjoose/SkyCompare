package handler

import (
	"encoding/json"
	"net/http"
	"skycompare-backend-main/internal/service"
)

type AirportHandler struct {
	Service *service.AirportService
}

func (h *AirportHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	airports, err := h.Service.GetAllAirports()
	if err != nil {
		http.Error(w, "Error fetching airports", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airports)
}

func (h *AirportHandler) GetWithoutOne(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	exclude := r.Form.Get("IATA")

	airports, err := h.Service.GetAirportsWithoutOne(exclude)
	if err != nil {
		http.Error(w, "Error filtering airports", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airports)
}
