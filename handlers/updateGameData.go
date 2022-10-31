package handlers

import (
	"tripoley-server/helpers"
	"tripoley-server/models"
)

func updateGameData(game *models.GameData) {
	if game.NeedToDeal == true {
		helpers.Deal(game)
		game.NeedToDeal = false
	}
}
