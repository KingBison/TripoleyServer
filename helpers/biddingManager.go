package helpers

import "tripoley-server/models"

func BiddingManager(game *models.GameData) bool {
	for _, v := range game.Players {
		if v.PokerBidStatus == "done" {
			return true
		}
	}

	return false
}
