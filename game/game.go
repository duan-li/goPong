package game

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"time"
)

type Game struct {
	Screen tcell.Screen
	Ball   *Ball
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

	g := Game{
		Screen: s,
		Ball:   ball,
	}
	return &g
}

func (g *Game) Run() {
	s := g.Screen
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	for {

		s.Clear()

		width, height := s.Size()

		g.Ball.CheckEdges(width, height)

		g.Ball.Update()
		s.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, defStyle)

		time.Sleep(40 * time.Millisecond)
		s.Show()
	}

}
