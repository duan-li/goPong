package main

import (
	"github.com/duan-li/goPong/console"
	"github.com/duan-li/goPong/game"
)

func main() {
	g := game.NewGame()

	go g.Run()

	console.Control(g)
}
