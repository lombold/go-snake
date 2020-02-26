package game

type Direction int8

const (
	North Direction = iota
	East
	South
	West
)

type Turn struct {
	from  Direction
	to    Direction
	point Point
}
