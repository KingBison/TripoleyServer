package helpers

import (
	"tripoley-server/models"
)

func PlayPoker(game *models.GameData) []models.Player {
	players := &game.Players
	SuitC := models.GetSuitConv()
	NumC := models.GetNumberConv()

	winner := []models.Player{
		{
			PokerStats: models.PokerStats{
				Rank:    0,
				SubRank: 0,
			},
		},
	}

	for i, player := range *players {
		hand := player.PokerHand
		Phand := []models.PokerCard{}
		for _, v := range hand {
			Phand = append(Phand, models.PokerCard{
				Suit:   SuitC[v.Suit],
				Number: NumC[v.Number],
			})
		}
		game.Players[i].PokerStats.Rank = GetPokerRank(Phand)
		game.Players[i].PokerStats.SubRank = GetPokerSubRank(Phand, game.Players[i].PokerStats.Rank)
		if game.Players[i].PokerStats.Rank > winner[0].PokerStats.Rank {
			winner = []models.Player{}
			winner = append(winner, game.Players[i])
		} else if game.Players[i].PokerStats.Rank == winner[0].PokerStats.Rank && game.Players[i].PokerStats.SubRank > winner[0].PokerStats.SubRank {
			winner = []models.Player{}
			winner = append(winner, game.Players[i])
		} else if game.Players[i].PokerStats.Rank == winner[0].PokerStats.Rank && game.Players[i].PokerStats.SubRank == winner[0].PokerStats.SubRank {
			winner = append(winner, game.Players[i])
		}
	}

	return winner

}
