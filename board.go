package main

type Board struct {
	grid [10][10]int
}

func NewBoard() *Board {
	b := &Board{}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			b.grid[i][j] = 0
		}
	}
	return b
}

func (b *Board) CanAddShip(x int, y int, length int, down bool) bool {
	if down {
		// y increasing
		for i := y; i < y+length; i++ {
			if i > 9 {
				return false
			}
			if b.grid[x][i] == 1 {
				return false
			}
		}
		return true
	}

	// x increasing
	for i := x; i < x+length; i++ {
		if i > 9 {
			return false
		}
		if b.grid[i][y] == 1 {
			return false
		}
	}

	return true
}

func (b *Board) AddShip(x int, y int, length int, down bool) {
	if down {
		// Mark 1 for taken location.
		for i := y; i < y+length; i++ {
			b.grid[x][i] = 1
		}
	} else {
		// Mark 1 for taken location.
		for i := x; i < x+length; i++ {
			b.grid[i][y] = 1
		}
	}
}
