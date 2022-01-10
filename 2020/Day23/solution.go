package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"strconv"
)

type Game struct {
	CurrentCup int
	Max        int
	Cups       []int
}

func createGame(input string, maxNum int) Game {
	inputNums := []int{}
	for _, num := range input {
		inputNums = append(inputNums, int(num-'0'))
	}

	g := Game{
		Cups:       make([]int, maxNum+1),
		CurrentCup: 0,
		Max:        maxNum,
	}

	start := inputNums[0]
	last := inputNums[len(inputNums)-1]
	g.CurrentCup = start

	for i := 0; i < len(inputNums)-1; i++ {
		g.Cups[inputNums[i]] = inputNums[i+1]
	}

	for i := len(inputNums) + 1; i <= maxNum; i++ {
		g.Cups[last] = i
		last = i
	}

	g.Cups[last] = start

	return g
}

func (g *Game) doMove() {
	cup1 := g.Cups[g.CurrentCup]
	cup2 := g.Cups[cup1]
	cup3 := g.Cups[cup2]
	after := g.Cups[cup3]

	destination := g.CurrentCup - 1
	if destination == 0 {
		destination = g.Max
	}
	for cup1 == destination || cup2 == destination || cup3 == destination {
		destination--
		if destination == 0 {
			destination = g.Max
		}
	}

	g.Cups[g.CurrentCup] = after

	oldDestValue := g.Cups[destination]
	g.Cups[destination] = cup1
	g.Cups[cup3] = oldDestValue
	g.CurrentCup = after
}

func (g *Game) playGame(rounds int) {
	for i := 0; i < rounds; i++ {
		g.doMove()
	}
}

func (g Game) getOrderString() string {
	str := ""
	for char := g.Cups[1]; char != 1; char = g.Cups[char] {
		str += strconv.Itoa(char)
	}

	return str
}

func (g Game) getProductOfLabels() int {
	cup1 := g.Cups[1]
	cup2 := g.Cups[cup1]
	return cup1 * cup2
}

func main() {
	input := utils.ReadFile()
	g := createGame(input[0], len(input[0]))
	g.playGame(100)
	fmt.Println("Part 1:", g.getOrderString())

	g = createGame(input[0], 1000000)
	g.playGame(10000000)
	fmt.Println("Part 2:", g.getProductOfLabels())
}
