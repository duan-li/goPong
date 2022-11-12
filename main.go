package main

import (
	"github.com/duan-li/goPong/game"
	"github.com/gdamore/tcell/v2"
	"log"
)

func main() {

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

	ball := game.NewBall()

	g := game.Game{
		Screen: s,
		Ball:   ball,
	}

	go g.Run()
	game.ReactWithInput(&g)

}
