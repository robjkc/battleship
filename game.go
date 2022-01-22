package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Location struct {
	x int
	y int
}

func (l Location) Description() string {
	letters := "ABCDEFGHIJ"
	return fmt.Sprintf("%s%d", string(letters[l.x]), l.y+1)
}

type Game struct {
	numPlayers      int
	players         [2]*Player
	currentPlayer   int
	board           *Board
	computerChoices map[Location]bool
}

func NewGame(numPlayers int) *Game {
	g := &Game{}
	g.numPlayers = numPlayers
	g.currentPlayer = 0
	g.board = NewBoard()
	g.computerChoices = make(map[Location]bool)

	return g
}

func (g *Game) LayoutShips() {
	rand.Seed(time.Now().UnixNano())

	// Create the ships and players.
	for i := 0; i < 2; i++ {
		ships := g.getShips()
		g.players[i] = NewPlayer(ships)
	}

	if g.numPlayers == 2 {
		fmt.Println("Player 1 ships:")
		g.players[0].DisplayShips()
		fmt.Println("Player 2 ships:")
		g.players[1].DisplayShips()
	} else {
		fmt.Println("My ships:")
		g.players[0].DisplayShips()
		fmt.Println("Computer ships:")
		g.players[1].DisplayShips()
	}
}

func (g *Game) Play() {
	if g.numPlayers == 2 {
		g.playTwoPlayer()
	} else {
		g.playComputer()
	}
}

func (g *Game) playTwoPlayer() {
	for {
		location := g.getLocation()
		targetPlayer := g.GetTargetPlayer()

		result := g.FireMissle(targetPlayer, location)
		fmt.Println(result)

		if !targetPlayer.IsAlive() {
			fmt.Println("[ Player", g.currentPlayer+1, "] is the winner!")
			break
		}

		// Toggle the player.
		g.currentPlayer = g.GetTargetPlayerNumber()
	}
}

func (g *Game) playComputer() {
	for {
		location := g.getLocation()

		targetPlayer := g.GetTargetPlayer()

		result := g.FireMissle(targetPlayer, location)
		if g.currentPlayer == 1 {
			fmt.Println("[Computer guessed:", location.Description()+"]:", result)
		} else {
			fmt.Println(result)
		}

		if !targetPlayer.IsAlive() {
			if g.currentPlayer == 0 {
				fmt.Println("You are the winner!")
			} else {
				fmt.Println("Computer won!")
			}
			break
		}

		// Toggle the player.
		g.currentPlayer = g.GetTargetPlayerNumber()
	}
}

func (g *Game) GetTargetPlayerNumber() int {
	if g.currentPlayer == 0 {
		return 1
	}
	return 0
}

func (g *Game) GetTargetPlayer() *Player {
	return g.players[g.GetTargetPlayerNumber()]
}

func (g *Game) FireMissle(targetPlayer *Player, location Location) string {

	result := targetPlayer.CheckIncomingMissle(location)
	switch result {
	case Hit:
		return "Hit!"

	case Sunk:
		return "You sunk my battleship!"
	}
	return "Miss!"
}

func (g *Game) getShips() []*Ship {
	ships := []*Ship{}
	for length := 2; length < 6; length++ {
		for {
			x := rand.Intn(10)
			y := rand.Intn(10)
			down := rand.Intn(2) == 1
			if g.board.CanAddShip(x, y, length, down) {
				g.board.AddShip(x, y, length, down)
				ship := NewShip(x, y, length, down)
				ships = append(ships, ship)
				//fmt.Println("Added ship x=", x, " y=", y, "length=", length, " down=", down)
				break
			}
		}
		if len(ships) == 4 {
			break
		}
	}
	return ships
}

func (g *Game) getComputerLocation() Location {
	for {
		// Generate a random location to fire for the computer.
		x := rand.Intn(10)
		y := rand.Intn(10)
		l := Location{x: x, y: y}
		_, exists := g.computerChoices[l]
		if exists {
			// Computer already made this choice.
			continue
		}
		g.computerChoices[l] = true
		return l
	}
}

func (g *Game) getLocation() Location {
	if g.numPlayers == 1 && g.currentPlayer == 1 {
		return g.getComputerLocation()
	}
	for {
		if g.numPlayers == 2 {
			fmt.Println("[ Player", g.currentPlayer+1, "] Specify location to fire (a1 - j10): ")
		} else {
			fmt.Println("Specify location to fire (a1 - j10): ")
		}
		var location string
		fmt.Scanln(&location)
		if len(location) > 3 {
			continue
		}
		x := toX(location)
		if x < 0 || x > 9 {
			continue
		}
		y := toY(location)

		if y < 0 || y > 9 {
			continue
		}
		l := Location{x: x, y: y}
		return l
	}
}

func toX(location string) int {
	letters := "ABCDEFGHIJ"

	firstLetter := strings.ToUpper(location[0:1])
	for i := 0; i < len(letters); i++ {
		if firstLetter == string(letters[i]) {
			return i
		}
	}
	return -1
}

func toY(location string) int {
	lastLetter := strings.ToUpper(location[1:])
	if _, err := strconv.Atoi(lastLetter); err == nil {
		y, _ := strconv.Atoi(lastLetter)
		return y - 1
	}
	return -1
}
