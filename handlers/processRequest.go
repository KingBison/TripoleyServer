package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tripoley-server/helpers"
	"tripoley-server/models"
)

func ProcessRequest(game *models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error Reading Body"))
			return
		}
		var req models.GameRequest
		err = json.Unmarshal(reqData, &req)
		player := helpers.GetPlayer(game, req.Name)
		if player.Name == "" {
			w.WriteHeader(500)
			w.Write([]byte("Player Not Found"))
			return
		}
		fmt.Println(player.Name, " requests the action: ", req.Action)

	}
}
