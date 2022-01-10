package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
)

type coordinate struct {
	X int
	Y int
}

var maxCol = 0
var maxRow = 0
var tick = 0

type cart struct {
	coord     coordinate
	tick      int
	direction string
	turns     int
}

var track [][]string

var trackPieces = make(map[coordinate]string)

var validCarts = make(map[string]bool)

var allCarts = make(map[coordinate]cart)

func populateTrack(input []string) (track [][]string) {
	track = [][]string{}
	for j, col := range input {
		track = append(track, []string{})
		if j > maxCol {
			maxCol = j
		}
		for i, row := range col {
			piece := string(row)
			track[j] = append(track[j], piece)
			co := coordinate{X: i, Y: j}
			trackPieces[co] = piece
			if _, ok := validCarts[piece]; ok {
				populateCarts(co, piece)
				convertCartToTrack(co, piece)
			}
			if i > maxRow {
				maxRow = i
			}
		}
	}
	return
}

func convertCartToTrack(co coordinate, cart string) {
	if cart == "v" || cart == "^" {
		trackPieces[co] = "|"
	} else {
		trackPieces[co] = "-"
	}
}

func printTrack(track [][]string) {
	for j, row := range track {
		for i := range row {
			if val, ok := allCarts[coordinate{i, j}]; ok {
				fmt.Print(val.direction)
			} else {
				fmt.Print(trackPieces[coordinate{i, j}])
			}
		}
		fmt.Println()
	}
}

func checkForCrash(co coordinate) (crashed bool) {
	for _, c := range allCarts {
		if c.coord.X == co.X && c.coord.Y == co.Y {
			return true
		}
	}
	return
}

func populateCarts(co coordinate, piece string) {
	cart := cart{
		coord:     co,
		tick:      0,
		direction: piece,
		turns:     0,
	}
	allCarts[co] = cart
}

var handleTurns = map[string][]string{
	"^": {"<", "^", ">", ">", "<"},
	">": {"^", ">", "v", "^", "v"},
	"v": {">", "v", "<", "<", ">"},
	"<": {"v", "<", "^", "v", "^"},
}

func moveCart(cart cart) {
	oldCoord := cart.coord
	newCoord := coordinate{}
	switch cart.direction {
	case "^":
		newCoord = coordinate{X: cart.coord.X, Y: cart.coord.Y - 1}
	case ">":
		newCoord = coordinate{X: cart.coord.X + 1, Y: cart.coord.Y}
	case "v":
		newCoord = coordinate{X: cart.coord.X, Y: cart.coord.Y + 1}
	case "<":
		newCoord = coordinate{X: cart.coord.X - 1, Y: cart.coord.Y}
	}
	if trackPieces[newCoord] == "/" {
		cart.direction = handleTurns[cart.direction][3]
	} else if trackPieces[newCoord] == "\\" {
		cart.direction = handleTurns[cart.direction][4]
	} else if trackPieces[newCoord] == "+" {
		cart.direction = handleTurns[cart.direction][cart.turns%3]
		cart.turns++
	}
	track[oldCoord.Y][oldCoord.X] = trackPieces[oldCoord]
	track[newCoord.Y][newCoord.X] = cart.direction
	cart.coord = newCoord
	cart.tick++
	delete(allCarts, oldCoord)
	if crashed := checkForCrash(newCoord); crashed {
		delete(allCarts, newCoord)
		return
	}
	allCarts[newCoord] = cart
}

func main() {
	input := utils.ReadFile()
	validCarts["^"] = true
	validCarts[">"] = true
	validCarts["v"] = true
	validCarts["<"] = true
	track = populateTrack(input)
	printTrack(track)

	for {
		for j := 0; j <= maxCol; j++ {
			for i := 0; i <= maxRow; i++ {
				if val, ok := allCarts[coordinate{i, j}]; ok && val.tick == tick {
					moveCart(val)
				}
			}
		}
		// printTrack(track)
		// fmt.Println()
		if len(allCarts) == 1 {
			for k := range allCarts {
				fmt.Println(k)
			}
			return
		}
		tick++
	}
}
