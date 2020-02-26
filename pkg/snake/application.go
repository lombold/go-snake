package snake

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lombold/snake/pkg/snake/game"
	"log"
)

var config = Frame{
	Width:  1000,
	Height: 1000,
	Scale:  1,
	Title:  "Snakeerito",
}

func Init() {
	gameView := new(GameView)
	board := game.InitBoard()
	gameController := InitGameController(gameView, &board)
	if err := ebiten.Run(gameController.Update, config.Width, config.Height, config.Scale, config.Title); err != nil {
		log.Fatal(err)
	}
}
