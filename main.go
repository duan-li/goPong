package main

import (
	"github.com/duan-li/goPong/game"
)

func main() {
	g := game.NewGame()

	go g.Run()

	game.ReactWithInput(g)
}
