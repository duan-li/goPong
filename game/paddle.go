package game

import "strings"

type Paddle struct {
	X      int
	Y      int
	Xspeed int
	Yspeed int
	width  int
	height int
}

func NewPaddle() *Paddle {
	paddle := Paddle{
		X:      3,
		Y:      1,
		Xspeed: 0,
		Yspeed: 3,
		width:  20,
		height: 1,
	}
	return &paddle
}

func (p *Paddle) Display() string {
	return strings.Repeat(" ", p.width)
}

func (p *Paddle) MoveUp() {

	if p.Y > 0 {
		p.Y -= p.Yspeed
	}
}

func (p *Paddle) MoveDown(windowHeight int) {
	if p.Y < windowHeight-p.height {
		p.Y += p.Yspeed
	}
}

func (p *Paddle) MoveLeft() {
	if p.Y > 0 {
		p.Y -= p.Xspeed
	}
}

func (p *Paddle) MoveRight(windowWidth int) {
	if p.Y < windowWidth-p.width {
		p.Y += p.Xspeed
	}
}
