package game

import (
	"errors"
	"log"
)

const SPEED = 1

type Snaker interface {
	Length() int
	Slither()
	Eat()
}

type Snake struct {
	headPosition Point
	headTurn     Turn
	length       int
	turns        []Turn
}

type slitherError struct {
	Error error
}

func (s *Snake) Length() int {
	return s.length
}

func (s *Snake) Slither() *slitherError {
	log.Print("Slithering")
	switch s.headTurn.to {
	case North:
		s.headPosition.y -= 1
		break
	case East:
		s.headPosition.x += 1
		break
	case South:
		s.headPosition.y += 1
		break
	case West:
		s.headPosition.x -= 1
		break
	}

	if len(s.turns) > 0 && s.headTurn == s.turns[0] {
		s.headTurn.from = (s.headTurn.to + 2) % 4
	}

	s.headTurn.point = s.headPosition

	if s.headPosition.x >= GAME_WIDTH ||
		s.headPosition.x < 0 ||
		s.headPosition.y >= GAME_HEIGHT ||
		s.headPosition.y < 0 {
		log.Print("Boarder reached!")
		return &slitherError{errors.New("can not slither")}
	}

	return nil
}

func (s *Snake) Eat(candy Candy) bool {
	if candy.position == s.headPosition {
		s.length++
		candy.eaten = true
	}
	return candy.eaten
}
