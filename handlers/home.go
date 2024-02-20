package handlers

import (
	"encoding/json"
	"net/http"
	"rest_websocket/server"
)

// client
type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// cod http 200 - >http.StatusOK
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to this curse",
			Status:  true,
		})
	}

}
