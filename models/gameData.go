package models

type GameData struct {
	Players        []Player
	GameRunning    bool
	GameState      string
	BiddingPointer Player
	PokerWinner    Player
	PokerStart     Player
}
