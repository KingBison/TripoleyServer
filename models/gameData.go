package models

type GameData struct {
	Players     []Player
	GameRunning bool
	NeedToDeal  bool
	PokerWinner Player
}
