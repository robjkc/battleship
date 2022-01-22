package main

import (
	"fmt"
	"strconv"
)

func main() {

	numPlayers := getNumPlayers()

	game := NewGame(numPlayers)
	game.LayoutShips()
	game.Play()
}

func getNumPlayers() int {
	for {
		fmt.Println("Enter number of players (1 or 2): ")
		var numPlayers string
		fmt.Scanln(&numPlayers)
		if _, err := strconv.Atoi(numPlayers); err == nil {
			players, _ := strconv.Atoi(numPlayers)
			if players < 1 || players > 2 {
				continue
			}
			return players
		}
	}
}
