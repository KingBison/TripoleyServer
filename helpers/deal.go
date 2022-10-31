package helpers

import (
	"math/rand"
	"time"
	"tripoley-server/models"
)

func Deal(game *models.GameData) {
	suits := []string{"C", "D", "H", "S"}
	numbers := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	deck := []models.Card{}
	for _, suit := range suits {
		for _, number := range numbers {
			deck = append(deck, models.Card{
				Number: number,
				Suit:   suit,
			})
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	ppointer := 0
	for len(deck) > 0 {
		pop := deck[0]
		deck = deck[1:]
		game.Players[ppointer].Hand = append(game.Players[ppointer].Hand, pop)
		ppointer++
		if ppointer > (len(game.Players) - 1) {
			ppointer = 0
		}

	}

}
