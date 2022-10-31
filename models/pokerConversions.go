package models

func GetSuitConv() map[string]int {
	Conv := map[string]int{
		"C": 0,
		"D": 1,
		"H": 2,
		"S": 3,
	}

	return Conv
}

func GetNumberConv() map[string]int {
	Conv := map[string]int{
		"2":  0,
		"3":  1,
		"4":  2,
		"5":  3,
		"6":  4,
		"7":  5,
		"8":  6,
		"9":  7,
		"10": 8,
		"J":  9,
		"Q":  10,
		"K":  11,
		"A":  12,
	}

	return Conv
}

type CC struct {
	Num   int
	Count int
}
