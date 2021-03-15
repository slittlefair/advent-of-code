package main

import (
	"Advent-of-Code"
	"fmt"
	"regexp"
)

var totalTime = 2503

type reindeer struct {
	name        string
	speed       int
	duration    int
	rest        int
	currentMove int
	currentRest int
	flying      bool
	totalDist   int
	points      int
}

var herd = make(map[string]reindeer)

func reindeerMove(r reindeer) reindeer {
	if !r.flying {
		r.currentRest++
		if r.currentRest == r.rest {
			r.flying = true
			r.currentRest = 0
		}
	} else {
		r.totalDist += r.speed
		r.currentMove++
		if r.currentMove == r.duration {
			r.flying = false
			r.currentMove = 0
		}
	}
	return r
}

func main() {
	lines := helpers.ReadFile()
	reNums := regexp.MustCompile("\\d+")
	reName := regexp.MustCompile("[A-Z][a-z]+")
	for _, l := range lines {
		nums := helpers.StringSliceToIntSlice(reNums.FindAllString(l, -1))
		name := reName.FindAllString(l, -1)
		herd[name[0]] = reindeer{
			name:     name[0],
			speed:    nums[0],
			duration: nums[1],
			rest:     nums[2],
			flying:   true,
		}
	}

	for i := 0; i < totalTime; i++ {
		for name, deer := range herd {
			herd[name] = reindeerMove(deer)
		}
		leadingDist := 0
		for _, deer := range herd {
			if deer.totalDist > leadingDist {
				leadingDist = deer.totalDist
			}
		}
		for name, deer := range herd {
			if deer.totalDist == leadingDist {
				deer.points++
				herd[name] = deer
			}
		}
	}

	var winningDist = 0
	var winningPoints = 0
	for _, deer := range herd {
		if deer.totalDist > winningDist {
			winningDist = deer.totalDist
		}
		if deer.points > winningPoints {
			winningPoints = deer.points
		}
	}
	fmt.Println("Part 1:", winningDist)
	fmt.Println("Part 2:", winningPoints)
}
