package game

import (
	"github.com/gdamore/tcell/v2"
	"os"
)

func ReactWithInput(g *Game) {
	for {

		switch event := g.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			g.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				g.Screen.Fini()
				os.Exit(0)
			}
		}
	}
}
