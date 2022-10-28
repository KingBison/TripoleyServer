package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tripoley-server/models"
)

func containsPlayer(players []models.Player, newPlayer models.Player) bool {
	for _, v := range players {
		if v.Name == newPlayer.Name {
			return true
		}
	}
	return false
}

func NewPlayer(game *models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Failed to Read Body"))
			return
		}
		var newPlayer models.Player
		err = json.Unmarshal(data, &newPlayer)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Failed to Unmarshal Body"))
			return
		}

		if containsPlayer(game.Players, newPlayer) {
			w.WriteHeader(400)
			w.Write([]byte("Player Already Exists"))
			return
		}

		if newPlayer.Name == "" {
			w.WriteHeader(400)
			w.Write([]byte("Invalid Player Name"))
			return
		}

		newPlayer.Money = 150

		game.Players = append(game.Players, newPlayer)
		w.WriteHeader(201)
	}
}
