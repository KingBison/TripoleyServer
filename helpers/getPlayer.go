package helpers

import "tripoley-server/models"

func GetPlayer(game *models.GameData, player string) models.Player {
	for i, v := range game.Players {
		if v.Name == player {
			return game.Players[i]
		}
	}
	return models.Player{Name: ""}
}
