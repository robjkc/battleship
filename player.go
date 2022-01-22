package main

import "fmt"

type Player struct {
	ships []*Ship
}

func NewPlayer(ships []*Ship) *Player {
	p := &Player{}
	p.ships = ships

	return p
}

func (p *Player) DisplayShips() {
	for _, ship := range p.ships {
		fmt.Println(ship.GetLocationDesc())
	}
}

func (p *Player) NumShips() int {
	return len(p.ships)
}

func (p *Player) IsAlive() bool {
	for _, ship := range p.ships {
		if !ship.IsSunk() {
			return true
		}
	}
	return false
}

func (p *Player) CheckIncomingMissle(location Location) FireResult {
	for _, ship := range p.ships {
		hit := ship.IsHit(location)
		if hit {
			if ship.IsSunk() {
				return Sunk
			} else {
				return Hit
			}
		}
	}
	return Miss
}
