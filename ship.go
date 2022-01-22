package main

import "fmt"

type FireResult int

const (
	Miss FireResult = iota + 1
	Hit
	Sunk
)

type Ship struct {
	length int
	down   bool
	startX int
	startY int
	sunk   bool
	hits   []Location
}

func NewShip(x int, y int, length int, down bool) *Ship {
	s := &Ship{}
	s.length = length
	s.down = down
	s.startX = x
	s.startY = y
	s.sunk = false
	s.hits = []Location{}

	return s
}

func (s *Ship) GetLocationDesc() string {
	desc := ""
	for i := 0; i < s.length; i++ {

		if s.down {
			location := Location{x: s.startX, y: s.startY + i}
			desc = fmt.Sprintf("%s%s", desc, location.Description())
		} else {
			location := Location{x: s.startX + i, y: s.startY}
			desc = fmt.Sprintf("%s%s", desc, location.Description())
		}
		if i < s.length-1 {
			desc += ","
		}
	}
	return desc
}

func (s *Ship) IsHit(location Location) bool {
	if s.hasBeenHit(location) {
		return true
	}
	hit := false
	if s.down {
		hitX := location.x == s.startX
		if hitX {
			hit = location.y >= s.startY && location.y < s.startY+s.length
		}
	} else {
		hitY := location.y == s.startY
		if hitY {
			hit = location.x >= s.startX && location.x < s.startX+s.length
		}
	}

	if hit {
		s.hits = append(s.hits, location)
	}

	return hit
}

func (s *Ship) hasBeenHit(location Location) bool {
	for _, hit := range s.hits {
		if hit == location {
			return true
		}
	}
	return false
}

func (s *Ship) IsSunk() bool {
	return len(s.hits) == s.length
}
