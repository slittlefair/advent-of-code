package main

import (
	"Advent-of-Code/file"
	"fmt"
)

type Shape struct {
	ties   string
	beats  string
	loses  string
	points int
}

type Game struct {
	part1  int
	part2  int
	shapes map[string]Shape
}

func createGame() *Game {
	// rock beats C (scissors) for 1 point
	// paper beats A (rock) for 2 points
	// scissors beats B (paper) for 3 points
	rock := Shape{
		ties:   "A",
		beats:  "C",
		loses:  "B",
		points: 1,
	}
	paper := Shape{
		ties:   "B",
		beats:  "A",
		loses:  "C",
		points: 2,
	}
	scissors := Shape{
		ties:   "C",
		beats:  "B",
		loses:  "A",
		points: 3,
	}
	return &Game{
		part1: 0,
		part2: 0,
		shapes: map[string]Shape{
			"X": rock,
			"A": rock,
			"Y": paper,
			"B": paper,
			"Z": scissors,
			"C": scissors,
		},
	}
}

func (g *Game) playRound(round string) error {
	if len(round) != 3 || string(round[1]) != " " {
		return fmt.Errorf("incorrect round string provided: %s", round)
	}

	// Part 1
	me := string(round[2])
	them := string(round[0])
	g.part1 += g.shapes[me].points
	if g.shapes[me].ties == them {
		g.part1 += 3
	} else if g.shapes[me].beats == them {
		g.part1 += 6
	}

	//Part 2
	outcome := string(round[2])
	theirShape := g.shapes[them]
	switch outcome {
	case "X":
		me = theirShape.beats
	case "Y":
		me = theirShape.ties
		g.part2 += 3
	case "Z":
		me = theirShape.loses
		g.part2 += 6
	}
	g.part2 += g.shapes[me].points

	return nil
}

func (g *Game) playGames(input []string) error {
	for _, round := range input {
		err := g.playRound(round)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	input := file.Read()
	g := createGame()
	err := g.playGames(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", g.part1)
	fmt.Println("Part 2:", g.part2)
}
