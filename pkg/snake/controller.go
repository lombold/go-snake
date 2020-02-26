package snake

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/lombold/snake/pkg/snake/game"
	"log"
)

var direction game.Direction = -1

type Gamer interface {
	Update() bool
}

type GameController struct {
	view *GameView
	model *game.Board
}

func InitGameController(view *GameView, board *game.Board) *GameController {
	return &GameController{
		view: view,
		model:  board,
	}
}

func (g *GameController) Update(screen *ebiten.Image) error {

	// Tick GameController Logic
	checkInteraction()
	if &direction != nil && direction >= 0 {
		g.model.Turn(direction)
		direction = -1
	}
	g.model.Tick()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	g.view.Update(screen, g.model)

	return nil
}

func checkInteraction()  {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		direction = game.North
		log.Println("Changed Direction to", direction)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA){
		direction = game.West
		log.Println("Changed Direction to", direction)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS){
		direction = game.South
		log.Println("Changed Direction to", direction)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD){
		direction = game.East
		log.Println("Changed Direction to", direction)
	}
}
