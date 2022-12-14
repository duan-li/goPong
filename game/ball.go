package game

import (
	"crypto/rand"
	"math/big"
)

type Ball struct {
	X      int
	Y      int
	Xspeed int
	Yspeed int
}

func NewBall() *Ball {
	ball := Ball{
		X:      100,
		Y:      10,
		Xspeed: RandomSpeed(),
		Yspeed: RandomSpeed(),
	}
	return &ball
}

func (b *Ball) Display() string {
	return "\u25CF"
}

func (b *Ball) Update() {
	b.X += b.Xspeed
	b.Y += b.Yspeed
}

func (b *Ball) CheckEdges(maxWidth int, maxHeight int) {
	if b.X <= 0 || b.X >= maxWidth {
		b.Xspeed *= -1
	}

	if b.Y <= 0 || b.Y >= maxHeight {
		b.Yspeed *= -1
	}
}

func (b *Ball) intersects(p Paddle) bool {
	return b.X >= p.X && b.X <= p.X+p.width && b.Y >= p.Y && b.Y <= p.Y+p.height
}

func RandomSpeed() int {
	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	num := n.Int64()
	if num > 50 {
		return 1
	} else {
		return -1
	}
}

func (b *Ball) reverseX() {
	b.Xspeed *= -1
}

func (b *Ball) reverseY() {
	b.Yspeed *= -1
}
