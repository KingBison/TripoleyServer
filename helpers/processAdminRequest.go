package helpers

import (
	"tripoley-server/models"
)

func ProcessAdminRequest(game *models.GameData, req models.GameRequest) bool {
	if req.Action == "clear-players" {
		game.Players = []models.Player{}
		return true
	}
	if req.Action == "redeal" {
		Deal(game)
		return true
	}
	return false
}
