package snake

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/lombold/snake/pkg/snake/game"
	"image/color"
)

type GameViewer interface {
	Update(screen *ebiten.Image, cont *game.Board)
}

type GameView struct {
	cont *GameController
}

func (gv *GameView) Update(screen *ebiten.Image, board *game.Board) {
	screen.Fill(color.RGBA{0xff, 0, 0, 0x99})
	drawGitter(screen, board)
	drawMeta(screen, board)
}

func drawGitter(screen *ebiten.Image, board *game.Board) {
	//r := rand.New(rand.NewSource(99))
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))


	for x := 0; x < len(board.Board()); x++ {
		for y := 0; y < len(board.Board()[x]); y++ {
			switch board.Board()[x][y] {
			case 1:
				drawField(screen, x, y, color.Black)
				break
			case 2:
				drawField(screen, x, y, color.RGBA{
					R: 125,
					G: 50,
					B: 200,
					A: 255,
				})
				break
			case 3:
				drawField(screen, x, y, color.RGBA{
					R: 125,
					G: 50,
					B: 0,
					A: 255,
				})
				break
			default:
				drawField(screen, x, y, color.White)
				//drawField(screen, x, y, color.RGBA{
				//	R: uint8(r.Intn(255)),
				//	G: uint8(r.Intn(255)),
				//	B: uint8(r.Intn(255)),
				//	A: uint8(r.Intn(255)),
				//})
			}
		}
	}
}

func drawMeta(screen *ebiten.Image, board *game.Board){
	ebitenutil.DebugPrint(screen, fmt.Sprint("Snake Length: %i", board.Snake().Length()))
}

func drawField(screen *ebiten.Image, posX int, posY int, color color.Color) {
	//var x, y, width, height float64, clr color.Color
	var screenWidth, screenHeight = screen.Size()

	// log.Println("ScreenWidth: ", screenWidth, "; ScreenHeight: ", screenHeight)

	var x = float64(screenWidth * posX / game.GAME_WIDTH)
	var y = float64(screenHeight * posY / game.GAME_HEIGHT)

	var width = float64(screenWidth / game.GAME_WIDTH)
	var height = float64(screenHeight / game.GAME_HEIGHT)

	ebitenutil.DrawRect(screen, x, y, width, height, color)
}
