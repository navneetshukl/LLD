package main

import "math/rand"

type player struct {
	Name     string
	Position int
	IsWin    bool
}

func NewPlayer(name string) *player {
	return &player{
		Name:     name,
		Position: 0,
		IsWin: false,
	}
}

func (p *player) DiceThrow() int {
	val := rand.Intn(6)
	return val + 1
}

func (p *player) SetPosition(pos int) {
	p.Position = pos
}
