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
			desc = fmt.Sprintf("%s%s%d", desc, toLetter(s.startX), (s.startY + i + 1))
		} else {
			desc = fmt.Sprintf("%s%s%d", desc, toLetter(s.startX+i), (s.startY + 1))
		}
		if i < s.length-1 {
			desc += ","
		}
	}
	return desc
}

func toLetter(x int) string {
	letters := "ABCDEFGHIJ"
	return string(letters[x])
}

func (s *Ship) IncomingMissle(location Location) bool {
	if s.hasBeenHit(location) {
		return false
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
