package handlers

import (
	"net/http"
	"tripoley-server/models"
)

func HealthHandler(game *models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Tripoley Server is Running\n"))
		if game.GameRunning {
			w.Write([]byte("Game is Running"))
		} else {
			w.Write([]byte("Game is Not Running"))
		}
	}
}
