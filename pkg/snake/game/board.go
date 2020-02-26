package game

import (
	"log"
	"math/rand"
	"time"
)

const GAME_WIDTH = 12
const GAME_HEIGHT = 12
const UpdateNanos = 0.5 * 1000 * 1000 * 1000 // => 0.5 Sec


type Boarder interface {
	Board() [GAME_WIDTH][GAME_HEIGHT]int8
	Snake() *Snake
	Tick()
	Reset()
	Turn(direction Direction)

}

type Board struct {
	candy      Candy
	snake      Snake
	board      [GAME_WIDTH][GAME_HEIGHT]int8
	lastUpdate int64
}

func InitBoard() Board {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var gb = Board{
		candy: Candy{
			position: Point{x: r.Intn(GAME_WIDTH), y: r.Intn(GAME_HEIGHT)},
		},
		snake: Snake{
			headPosition: Point{x: GAME_WIDTH/2, y: GAME_HEIGHT/2},
			headTurn:     Turn{from: South, to: North},
			length:       5,
			turns:        []Turn{}, //{from: East, point: Point{x: 12, y: 13}}
		},
		lastUpdate: time.Now().UnixNano(),
	}

	return gb
}

func (gb *Board) Board() [GAME_WIDTH][GAME_HEIGHT]int8 {
	return gb.board
}

func (gb *Board) Snake() *Snake {
	return &gb.snake
}

func (gb *Board) Tick() {

	var now = time.Now().UnixNano()
	var diff =  now-gb.lastUpdate
	if diff >= UpdateNanos {
		gb.update()
		gb.lastUpdate = now
	}
}

func (gb *Board) Reset() {
	tempBoard := InitBoard()
	gb.snake = tempBoard.snake
	gb.candy = tempBoard.candy
	gb.actualiseMatrix()
}

func (gb *Board) update()  {
	if e := gb.snake.Slither(); e != nil {
		log.Println(e.Error)
		gb.gameOver()
	}

	if gb.snake.Eat(gb.candy) {
		gb.candy = gb.newCandy()
	}

	gb.actualiseMatrix()
}

func (gb *Board) gameOver() {
	gb.Reset()
}

//------- GAME LOGIC ----------

func (gb *Board) newCandy() Candy {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c := Candy{position: Point{x: r.Intn(GAME_WIDTH), y: r.Intn(GAME_HEIGHT)}}

	// Try again if there is already an object at the candy's position
	if gb.board[c.position.x][c.position.y] > 0 {
		return gb.newCandy()
	}

	return c
}

func (gb *Board) Turn(direction Direction) {
	if direction == gb.snake.headTurn.from || direction == gb.snake.headTurn.to {
		return
	}
	gb.snake.headTurn = Turn{point: gb.snake.headPosition, from: gb.snake.headTurn.from, to: direction}
	gb.snake.turns = append([]Turn{gb.snake.headTurn}, gb.snake.turns...)
	log.Println("Head Turn from: ", gb.snake.headTurn.from, " to: ", gb.snake.headTurn.to, " at ", gb.snake.headTurn.point.x, gb.snake.headTurn.point.x)
}

//------- GAME MATRIX ---------

func (gb *Board) actualiseMatrix() {
	gb.clear()
	gb.placeSnake()
	gb.placeCandy()
}

func (gb *Board) clear() {
	for x := 0; x < len(gb.board); x++ {
		for y := 0; y < len(gb.board[x]); y++ {
			gb.board[x][y] = 0
		}
	}
	log.Print("Filled board empty")
}

func (gb *Board) placeSnake() {
	currentPoint := gb.snake.headPosition
	currentTurn := gb.snake.headTurn
	turnIndex := 0
	log.Print("PlaceSnake")
	if currentPoint.x < 0 || currentPoint.y < 0 {
		log.Print("Panic")
	}

	for elementIndex := 0; elementIndex < gb.snake.length; elementIndex++ {

		log.Println("Head Position ", gb.snake.headPosition)
		log.Println("Head Turn ", gb.snake.headTurn)
		log.Println("Turns ", gb.snake.turns)

		if gb.board[currentPoint.x][currentPoint.y] == 3 || gb.board[currentPoint.x][currentPoint.y] == 2{
			gb.gameOver()
		}

		if elementIndex == 0 {
			gb.board[currentPoint.x][currentPoint.y] = 3
		} else {
			gb.board[currentPoint.x][currentPoint.y] = 2
		}

		switch currentTurn.from {
		case North:
			currentPoint = Point{x: currentPoint.x, y: currentPoint.y - 1}
			break
		case East:
			currentPoint = Point{x: currentPoint.x + 1, y: currentPoint.y}
			break
		case South:
			currentPoint = Point{x: currentPoint.x, y: currentPoint.y + 1}
			break
		case West:
			currentPoint = Point{x: currentPoint.x - 1, y: currentPoint.y}
			break
		}

		if len(gb.snake.turns) > turnIndex && gb.snake.turns[turnIndex].point == currentPoint {
			currentTurn = gb.snake.turns[turnIndex]
			turnIndex++
		}
	}
}

func (gb *Board) placeCandy() {
	if !gb.candy.eaten {
		gb.board[gb.candy.position.x][gb.candy.position.y] = 1
	}
}
