package main

import "fmt"

func main() {
	snakesCount, laddersCount, playersCount := 0, 0, 0
	fmt.Println("Enter the total number of snakes.")
	fmt.Scan(&snakesCount)
	snakesmap := make(map[int]int, snakesCount)
	fmt.Println("Enter the starting and ending position for snakes (e.g., 14 7).")
	for i := 0; i < snakesCount; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		snakesmap[a] = b
	}

	fmt.Println("Enter the total number of ladders.")
	fmt.Scan(&laddersCount)
	laddersmap := make(map[int]int, laddersCount)
	fmt.Println("Enter the starting and ending position for ladders (e.g., 3 22).")
	for i := 0; i < laddersCount; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		laddersmap[a] = b
	}

	fmt.Println("Enter the total number of players.")
	fmt.Scan(&playersCount)
	names := make([]string, playersCount)
	fmt.Println("Enter the names of the players.")
	for i := 0; i < playersCount; i++ {
		fmt.Scan(&names[i])
	}

	snakeSrv := NewSnake(snakesCount, snakesmap)
	ladderSrv := NewLadder(laddersCount, laddersmap)
	gameSrv := NewGame(snakeSrv, ladderSrv, names)
	gameSrv.Play()
}
