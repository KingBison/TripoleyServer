package helpers

import (
	"tripoley-server/models"
)

func ProcessAdminRequest(game *models.GameData, req models.GameRequest) bool {
	if req.Action == "deal" {
		game.GameRunning = !game.GameRunning
		game.NeedToDeal = true
		return true
	}
	if req.Action == "clear-players" {
		game.Players = []models.Player{}
		return true
	}
	return false
}
