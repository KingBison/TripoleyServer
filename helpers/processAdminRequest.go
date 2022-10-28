package helpers

import "tripoley-server/models"

func ProcessAdminRequest(game *models.GameData, req models.GameRequest) bool {
	if req.Action == "toggle-game-running" {
		game.GameRunning = !game.GameRunning
		return true
	}
	return false
}
