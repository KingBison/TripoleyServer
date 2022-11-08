package models

type Player struct {
	Name           string
	Money          int
	Ready          bool
	Hand           []Card
	PokerHand      []Card
	PokerStats     PokerStats
	PokerBid       int
	PokerBidStatus string
}
