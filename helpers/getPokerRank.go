package helpers

import (
	"math"
	"tripoley-server/models"
)

func GetPokerRank(hand []models.PokerCard) int {
	/// Royal Flush
	if checkRoyalFlush(hand) {
		return 10
	}
	if checkStraightFlush(hand) {
		return 9
	}
	if check4OfAKind(hand) {
		return 8
	}
	if checkFullHouse(hand) {
		return 7
	}
	if checkFlush(hand) {
		return 6
	}
	if checkStraight(hand) {
		return 5
	}
	if check3OfAKind(hand) {
		return 4
	}
	if check2Pair(hand) {
		return 3
	}
	if checkPair(hand) {
		return 2
	}
	return 1
}

func checkRoyalFlush(hand []models.PokerCard) bool {
	suit := hand[0].Suit
	for _, v := range hand {
		if v.Suit != suit {
			return false
		}
		if v.Number < 8 {
			return false
		}
	}
	return true
}

func checkStraightFlush(hand []models.PokerCard) bool {
	suit := hand[0].Suit
	for _, v := range hand {
		if v.Suit != suit {
			return false
		}
		for _, val := range hand {
			if !(int(math.Abs(float64(v.Number-val.Number))) < 5 || int(math.Abs(float64(val.Number-v.Number))) < 5) {
				return false
			}
		}

	}
	return true
}

func check4OfAKind(hand []models.PokerCard) bool {
	/// iter 1
	numCounts := []models.CC{}
	for _, v := range hand {
		new := true
		for i, c := range numCounts {
			if v.Number == c.Num {
				new = false
				numCounts[i].Count++
				break
			}
		}
		if new {
			numCounts = append(numCounts, models.CC{
				Num:   v.Number,
				Count: 1,
			})
		}
	}
	for _, v := range numCounts {
		if v.Count == 4 {
			return true
		}
	}

	return false
}

func checkFullHouse(hand []models.PokerCard) bool {
	/// iter 1
	numCounts := []models.CC{}
	for _, v := range hand {
		new := true
		for i, c := range numCounts {
			if v.Number == c.Num {
				new = false
				numCounts[i].Count++
				break
			}
		}
		if new {
			numCounts = append(numCounts, models.CC{
				Num:   v.Number,
				Count: 1,
			})
		}
	}

	check2 := false
	check3 := false
	for _, v := range numCounts {
		if v.Count == 3 {
			check3 = true
		}
		if v.Count == 2 {
			check2 = true
		}
	}
	if check2 && check3 {
		return true
	}
	return false

}

func checkFlush(hand []models.PokerCard) bool {
	suit := hand[0].Suit
	for _, v := range hand {
		if v.Suit != suit {
			return false
		}
	}
	return true
}

func checkStraight(hand []models.PokerCard) bool {
	repeats := []int{}
	for _, v := range hand {
		for _, val := range hand {
			if !(int(math.Abs(float64(v.Number-val.Number))) < 5 || int(math.Abs(float64(val.Number-v.Number))) < 5) {
				return false
			}

		}
		for _, r := range repeats {
			if r == v.Number {
				return false
			}
		}
		repeats = append(repeats, v.Number)

	}
	return true
}

func check3OfAKind(hand []models.PokerCard) bool {
	/// iter 1
	numCounts := []models.CC{}
	for _, v := range hand {
		new := true
		for i, c := range numCounts {
			if v.Number == c.Num {
				new = false
				numCounts[i].Count++
				break
			}
		}
		if new {
			numCounts = append(numCounts, models.CC{
				Num:   v.Number,
				Count: 1,
			})
		}
	}
	for _, v := range numCounts {
		if v.Count == 3 {
			return true
		}
	}

	return false
}

func check2Pair(hand []models.PokerCard) bool {
	/// iter 1
	numCounts := []models.CC{}
	for _, v := range hand {
		new := true
		for i, c := range numCounts {
			if v.Number == c.Num {
				new = false
				numCounts[i].Count++
				break
			}
		}
		if new {
			numCounts = append(numCounts, models.CC{
				Num:   v.Number,
				Count: 1,
			})
		}
	}
	pairCounts := 0
	for _, v := range numCounts {
		if v.Count == 2 {
			pairCounts++
		}
	}

	if pairCounts == 2 {
		return true
	}

	return false
}

func checkPair(hand []models.PokerCard) bool {
	/// iter 1
	numCounts := []models.CC{}
	for _, v := range hand {
		new := true
		for i, c := range numCounts {
			if v.Number == c.Num {
				new = false
				numCounts[i].Count++
				break
			}
		}
		if new {
			numCounts = append(numCounts, models.CC{
				Num:   v.Number,
				Count: 1,
			})
		}
	}

	for _, v := range numCounts {
		if v.Count == 2 {
			return true
		}
	}

	return false
}
