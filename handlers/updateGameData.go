package handlers

import (
	"tripoley-server/helpers"
	"tripoley-server/models"
)

func updateGameData(game *models.GameData) {
	if !game.GameRunning {
		return
	}
	if game.GameState == "begin" {
		for _, v := range game.Players {
			if v.Ready == false {
				return
			}
		}
		game.GameState = "deal"
		return
	}
	if game.GameState == "deal" {
		helpers.Deal(game)
		game.GameState = "bidding"
		for i := range game.Players {
			game.Players[i].Ready = false
			game.Players[i].PokerBid = 0
		}
	}

	if game.GameState == "bidding" {
		if helpers.BiddingManager(game) {
			game.GameState = "poker-select"
			return
		}
		return
	}

	if game.GameState == "poker-select" {
		for _, v := range game.Players {
			if v.Ready == false && len(v.PokerHand) != 5 {
				return
			}
		}
		helpers.PlayPoker(game)

		for i := range game.Players {
			game.Players[i].Ready = false
		}
		return
	}
	if game.GameState == "poker-resolve" {
		for _, v := range game.Players {
			if v.Ready == false {
				return
			}
		}
		game.GameState = "begin"
		for i := range game.Players {
			game.Players[i].Ready = false
		}
		return
	}

}
