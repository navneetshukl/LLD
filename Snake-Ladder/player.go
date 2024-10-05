package main

import "math/rand"

type player struct {
	Name     string
	Position int
	snake    *snake
	ladder   *ladder
}

func NewPlayer(name string, snake *snake, ladder *ladder) *player {
	return &player{
		Name:     name,
		Position: 0,
		snake:    snake,
		ladder:   ladder,
	}
}

func (p *player) DiceThrow() int {
	val := rand.Intn(6)
	return val+1
}

func(p *player)SetPosition(pos int){
	p.Position = pos
}