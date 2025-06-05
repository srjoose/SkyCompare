package handler

import (
	"encoding/json"
	"net/http"
	"skycompare-backend-main/internal/service"
)

type RouteHandler struct {
	Service *service.RouteService
}

func (h *RouteHandler) GetRoutes(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dep := r.Form.Get("dep")
	arr := r.Form.Get("arr")

	airways, err := h.Service.GetRoute(dep, arr)
	if err != nil {
		http.Error(w, "Error fetching routes", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airways)
}
