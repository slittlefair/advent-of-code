package rock

import "Advent-of-Code/graph"

type Rock []graph.Co

func Dash(y int) Rock {
	return Rock{
		{X: 3, Y: y + 4},
		{X: 4, Y: y + 4},
		{X: 5, Y: y + 4},
		{X: 6, Y: y + 4},
	}
}

func Cross(y int) Rock {
	return Rock{
		{X: 3, Y: y + 5},
		{X: 4, Y: y + 4},
		{X: 4, Y: y + 5},
		{X: 4, Y: y + 6},
		{X: 5, Y: y + 5},
	}
}

func L(y int) Rock {
	return Rock{
		{X: 3, Y: y + 4},
		{X: 4, Y: y + 4},
		{X: 5, Y: y + 4},
		{X: 5, Y: y + 5},
		{X: 5, Y: y + 6},
	}
}

func I(y int) Rock {
	return Rock{
		{X: 3, Y: y + 4},
		{X: 3, Y: y + 5},
		{X: 3, Y: y + 6},
		{X: 3, Y: y + 7},
	}
}

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
