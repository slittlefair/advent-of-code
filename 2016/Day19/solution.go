package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/timer"
	"fmt"
	"time"
)

type game struct {
	elves    map[int]struct{}
	numElves int
}

func createGame(num int) game {
	e := map[int]struct{}{}
	for i := 0; i < num; i++ {
		e[i] = struct{}{}
	}
	return game{
		elves:    e,
		numElves: num,
	}
}

func (g game) getPlayers() []int {
	players := []int{}
	for i := 0; i < g.numElves; i++ {
		if _, ok := g.elves[i]; ok {
			players = append(players, i)
		}
	}
	return players
}

func (g game) findWinner() int {
	i := 0
	for len(g.elves) > 1 {
		if _, ok := g.elves[i]; ok {
			j := (i + 1) % g.numElves
			for {
				if _, ok := g.elves[j]; ok {
					delete(g.elves, (j)%g.numElves)
					break
				}
				j++
				j = j % g.numElves
			}
		}
		i++
		i = i % g.numElves
	}
	var winner int
	for k := range g.elves {
		winner = k + 1
	}
	return winner
}

func main() {
	t := time.Now()
	input, err := file.ReadSingleLineAsInts()
	if err != nil {
		fmt.Println(err)
		return
	}
	g := createGame(input[0])
	winner := g.findWinner()
	fmt.Println("Part 1:", winner)
	timer.Track(t)
}
