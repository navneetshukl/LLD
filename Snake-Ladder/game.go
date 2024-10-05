package main

import "log"

type game struct {
	player []*player
}

func NewGame(snake *snake, ladder *ladder, names []string) *game {
	players := []*player{}
	for _, val := range names {
		pl := NewPlayer(val, snake, ladder)
		players = append(players, pl)
	}

	return &game{
		player: players,
	}
}

func (g *game) Play() {
	totalPlayer := len(g.player)
	win := 0
	idx := 0
	for {
		currentPlayer := g.player[idx%totalPlayer]
		diceResult := currentPlayer.DiceThrow()
		nextPos := currentPlayer.Position + diceResult
		if nextPos > 100 {
			continue
		}

		// check if snake
		snakePos := currentPlayer.snake.IsSnake(nextPos)
		if snakePos != -1 {
			currentPlayer.Position = snakePos
			continue
		}

		// check if ladder

		ladderPos := currentPlayer.ladder.IsLadder(nextPos)
		if ladderPos != -1 {
			currentPlayer.Position = ladderPos
		}

		if currentPlayer.Position == 100 {
			win++
			log.Printf("%s win !\n", currentPlayer.Name)
		} else{

		log.Printf("%s current position is %d !\n",currentPlayer.Name,currentPlayer.Position)
		}
		if win==totalPlayer{
			break
		}
		idx++

	}
}
