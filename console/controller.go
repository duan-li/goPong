package console

import (
	_g "github.com/duan-li/goPong/game"
	"github.com/gdamore/tcell/v2"
	"os"
)

func Control(game *_g.Game) {
	for {
		width, height := game.Screen.Size()
		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				game.Player.Paddle.MoveUp()
			} else if event.Key() == tcell.KeyDown {
				game.Player.Paddle.MoveDown(height)
			} else if event.Key() == tcell.KeyRight {
				game.Player.Paddle.MoveRight(width)
			} else if event.Key() == tcell.KeyLeft {
				game.Player.Paddle.MoveLeft()
			} else if event.Rune() == 'w' {
				game.Player.Paddle.MoveUp()
			} else if event.Rune() == 's' {
				game.Player.Paddle.MoveDown(height)
			}
		}
	}
}
