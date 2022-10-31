package helpers

import "tripoley-server/models"

func GetPokerSubRank(hand []models.PokerCard, rank int) int {
	if rank == 10 || rank == 9 || rank == 6 || rank == 5 || rank == 1 {
		maxScore := 0

		for _, v := range hand {
			newScore := 100*v.Number + v.Suit
			if newScore > maxScore {
				maxScore = newScore
			}
		}

		return maxScore
	}

	if rank == 8 || rank == 7 || rank == 4 {
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
			if v.Count == 4 || v.Count == 3 {
				return v.Num
			}
		}
	}

	if rank == 2 || rank == 3 {
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

		score := 0

		for _, v := range numCounts {
			if v.Count == 2 {
				bigger := 0
				smaller := 1000
				for _, v := range numCounts {
					if v.Count == 2 {
						if v.Num > bigger {
							bigger = v.Num
						}
						if v.Num < smaller {
							smaller = v.Num
						}
					}
				}
				score += 10000 * bigger
				score += 100 * smaller
			} else {
				score += v.Num
			}

		}

		return score

	}

	return -1

}
