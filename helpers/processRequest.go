package helpers

import "tripoley-server/models"

func ProcessRequest(game *models.GameData, player *models.Player, req models.GameRequest) bool {
	if req.Action == "toggle-ready" {
		player.Ready = !player.Ready
		return true
	}
	return false
}
