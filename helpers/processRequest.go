package helpers

import (
	"net/url"
	"tripoley-server/models"
)

func ProcessRequest(game *models.GameData, player *models.Player, req models.GameRequest, params url.Values) bool {
	if req.Action == "toggle-ready" {
		player.Ready = !player.Ready
		return true
	}
	if req.Action == "add-poker-card" {
		for _, v := range player.Hand {
			if v.Number == params["number"][0] && v.Suit == params["suit"][0] {
				for _, val := range player.PokerHand {
					if val.Number == params["number"][0] && val.Suit == params["suit"][0] {
						return false
					}
				}
				player.PokerHand = append(player.PokerHand, models.Card{Number: v.Number, Suit: v.Suit})
				return true
			}
		}
		return false
	}
	if req.Action == "remove-poker-card" {
		for i, v := range player.PokerHand {
			if v.Number == params["number"][0] && v.Suit == params["suit"][0] {
				player.PokerHand = append(player.PokerHand[:i], player.PokerHand[i+1:]...)
				return true
			}
		}
		return false
	}
	return false
}
