package game

type Player struct {
	Paddle Paddle
}

func NewPlayer() *Player {
	paddle := NewPaddle()

	player := Player{
		Paddle: *paddle,
	}
	return &player
}
