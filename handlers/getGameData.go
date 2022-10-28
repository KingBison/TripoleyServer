package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tripoley-server/models"
)

func GetGameData(game *models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		G, err := json.Marshal(game)
		if err != nil {
			w.Write([]byte("Failed to Get Game Data"))
			return
		}
		updateGameData(game)
		fmt.Println(game)
		w.Write(G)
	}
}
