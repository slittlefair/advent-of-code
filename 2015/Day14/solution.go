package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"fmt"
	"strconv"
	"strings"
)

type Reindeer struct {
	Name          string
	Speed         int
	Duration      int
	Rest          int
	CurrentMove   int
	CurrentRest   int
	TotalDistance int
	TotalPoints   int
	IsFlying      bool
}

type Racers map[string]*Reindeer

func (r *Reindeer) move() {
	r.TotalDistance += r.Speed
	r.CurrentMove++
	if r.CurrentMove == r.Duration {
		r.IsFlying = false
		r.CurrentMove = 0
	}
}

func (r *Reindeer) rest() {
	r.CurrentRest++
	if r.CurrentRest == r.Rest {
		r.IsFlying = true
		r.CurrentRest = 0
	}
}

func (rc Racers) givePoints() {
	leadingDistance := 0
	for _, r := range rc {
		if r.TotalDistance > leadingDistance {
			leadingDistance = r.TotalDistance
		}
	}
	for _, r := range rc {
		if r.TotalDistance == leadingDistance {
			r.TotalPoints++
		}
	}
}

func (rc Racers) runRace(length int) (int, int) {
	for i := 0; i < length; i++ {
		for _, r := range rc {
			if r.IsFlying {
				r.move()
			} else {
				r.rest()
			}
		}
		rc.givePoints()
	}

	winningDist := 0
	winningPoints := 0
	for _, r := range rc {
		if r.TotalDistance > winningDist {
			winningDist = r.TotalDistance
		}
		if r.TotalPoints > winningPoints {
			winningPoints = r.TotalPoints
		}
	}
	return winningDist, winningPoints
}

func parseInput(input []string) Racers {
	racers := Racers{}
	for _, reindeer := range input {
		split := strings.Split(reindeer, " ")
		name := split[0]
		nums := regex.MatchNums.FindAllString(reindeer, -1)
		// We can ignore the errors as we know they'll convert due to regex match
		speed, _ := strconv.Atoi(nums[0])
		duration, _ := strconv.Atoi(nums[1])
		rest, _ := strconv.Atoi(nums[2])
		racers[name] = &Reindeer{
			Name:     name,
			Speed:    speed,
			Duration: duration,
			Rest:     rest,
			IsFlying: true,
		}
	}
	return racers
}

func main() {
	input := file.Read()
	racers := parseInput(input)
	winningDist, winningPoints := racers.runRace(2503)
	fmt.Println("Part 1:", winningDist)
	fmt.Println("Part 2:", winningPoints)
}
