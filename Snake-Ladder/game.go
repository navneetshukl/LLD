package main

import "log"

type game struct {
	player []*player
	snake  *snake
	ladder *ladder
}

func NewGame(snake *snake, ladder *ladder, names []string) *game {
	players := []*player{}
	for _, val := range names {
		pl := NewPlayer(val)
		players = append(players, pl)
	}

	return &game{
		player: players,
		snake:  snake,
		ladder: ladder,
	}
}

// Play start the game, each player take turn to throw a dice,
// and move the position accordingly. If the player land on a snake,
// the player position will be changed to the snake end position.
// If the player land on a ladder, the player position will be changed
// to the ladder end position. If the player position exceed 100, the
// player will stay in the same position. The game will end when all
// player reach the position 100.

func (g *game) Play() {
	totalPlayer := len(g.player)
	win := 0
	idx := 0
	for {
		currentPlayer := g.player[idx%totalPlayer]
		diceResult := currentPlayer.DiceThrow()
		nextPos := currentPlayer.Position + diceResult
		if nextPos > 100 || currentPlayer.IsWin {
			idx++
			continue
		}
		pos:=nextPos

		// check if snake
		snakePos := g.snake.IsSnake(nextPos)
		if snakePos != -1 {
			pos=snakePos
		}

		// check if ladder

		ladderPos := g.ladder.IsLadder(nextPos)
		if ladderPos != -1 {
			pos=ladderPos
		}
		log.Printf("%s rolled a %d and moved to position %d\n",currentPlayer.Name,diceResult,pos)
		currentPlayer.SetPosition(pos)


		if currentPlayer.Position == 100 {
			win++
			currentPlayer.IsWin = true
			log.Printf("%s win !\n", currentPlayer.Name)
		}
		if win == totalPlayer {
			break
			
		}
		idx++

	}
}
