package rock

import "Advent-of-Code/graph"

type Rock []graph.Co

// Dash creates a dash shaped rock relative to the given y coordinate
func Dash(y int) Rock {
	return Rock{
		{X: 3, Y: y + 4},
		{X: 4, Y: y + 4},
		{X: 5, Y: y + 4},
		{X: 6, Y: y + 4},
	}
}

// Dash creates a cross shaped rock relative to the given y coordinate
func Cross(y int) Rock {
	return Rock{
		{X: 3, Y: y + 5},
		{X: 4, Y: y + 4},
		{X: 4, Y: y + 5},
		{X: 4, Y: y + 6},
		{X: 5, Y: y + 5},
	}
}

// Dash creates an L shaped rock relative to the given y coordinate
func L(y int) Rock {
	return Rock{
		{X: 3, Y: y + 4},
		{X: 4, Y: y + 4},
		{X: 5, Y: y + 4},
		{X: 5, Y: y + 5},
		{X: 5, Y: y + 6},
	}
}

// Dash creates an I shaped rock relative to the given y coordinate
func I(y int) Rock {
	return Rock{
		{X: 3, Y: y + 4},
		{X: 3, Y: y + 5},
		{X: 3, Y: y + 6},
		{X: 3, Y: y + 7},
	}
}

// Dash creates a square shaped rock relative to the given y coordinate
func Square(y int) Rock {
	return Rock{
		{X: 3, Y: y + 4},
		{X: 4, Y: y + 4},
		{X: 3, Y: y + 5},
		{X: 4, Y: y + 5},
	}
}

var Pieces = []func(int) Rock{
	Dash, Cross, L, I, Square,
}
