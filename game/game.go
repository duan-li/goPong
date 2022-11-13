package game

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"time"
)

type Game struct {
	Screen tcell.Screen
	Ball   *Ball
	Player *Player
}

func NewGame() *Game {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	ball := NewBall()

	player := NewPlayer()

	g := Game{
		Screen: s,
		Ball:   ball,
		Player: player,
	}
	return &g
}

func (g *Game) Run() {
	s := g.Screen
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	for {

		s.Clear()

		width, height := s.Size()

		// check collision between ball and paddle
		if g.Ball.intersects(g.Player.Paddle) {
			g.Ball.reverseX()
			g.Ball.reverseY()
		}

		// check collision between ball and edges
		g.Ball.CheckEdges(width, height)

		g.Ball.Update()

		// display the ball
		drawSprite(s,
			g.Ball.X,
			g.Ball.Y,
			g.Ball.X,
			g.Ball.Y,
			defStyle,
			g.Ball.Display())

		// display the paddles
		paddleStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
		drawSprite(s,
			g.Player.Paddle.X,
			g.Player.Paddle.Y,
			g.Player.Paddle.X+g.Player.Paddle.width,
			g.Player.Paddle.Y+g.Player.Paddle.height,
			paddleStyle,
			g.Player.Paddle.Display())

		time.Sleep(40 * time.Millisecond)
		s.Show()
	}

}

func drawSprite(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1

	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}
