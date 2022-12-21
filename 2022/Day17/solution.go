package main

import (
	"Advent-of-Code/2022/Day17/rock"
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"Advent-of-Code/slice"
	"fmt"
)

type Chamber map[graph.Co]bool

// Create the initial chamber, we will extend the side walls as and when
func createChamber() Chamber {
	chamber := Chamber{}
	for x := 0; x <= 8; x++ {
		chamber[graph.Co{X: x, Y: 0}] = true
	}
	for y := 0; y <= 10; y++ {
		chamber[graph.Co{X: 0, Y: y}] = true
		chamber[graph.Co{X: 8, Y: y}] = true
	}
	return chamber
}

func (c Chamber) extendWalls(n int) {
	for y := n; y <= n+10; y++ {
		c[graph.Co{X: 0, Y: y}] = true
		c[graph.Co{X: 8, Y: y}] = true
	}
}

func (c Chamber) highestRocks() int {
	max := 0
	for co := range c {
		if co.X > 0 && co.X < 8 {
			max = maths.Max(max, co.Y)
		}
	}
	return max
}

func (c Chamber) move(r rock.Rock, x int, y int) (rock.Rock, bool) {
	potential := rock.Rock{}
	for _, co := range r {
		potential = append(potential, graph.Co{X: co.X + x, Y: co.Y + y})
	}
	for _, pCo := range potential {
		for co := range c {
			if !slice.Contains(r, pCo) && co == pCo {
				return r, false
			}
		}
	}
	for _, co := range r {
		delete(c, co)
	}
	for _, co := range potential {
		c[co] = true
	}
	return potential, true
}

func (c Chamber) highestPoints() [7]int {
	hp := [7]int{}
	for co := range c {
		if co.X > 0 && co.X < 8 {
			hp[co.X-1] = maths.Max(hp[co.X-1], co.Y)
		}
	}
	lowestHP := maths.Infinity
	for _, v := range hp {
		lowestHP = maths.Min(lowestHP, v)
	}
	for i, v := range hp {
		hp[i] = v - lowestHP
	}
	for co := range c {
		if co.Y < lowestHP {
			delete(c, co)
		}
	}
	return hp
}

type state struct {
	points [7]int
	rock   int
	height int
}

func playTetris(instructions string, n1, n2 int) (int, int, error) {
	chamber := createChamber()
	j := 0
	xTaken := map[int]int{}
	base := 0
	cache := map[[2]int]state{}
	var part1 int
	for i := 0; i < n2; i++ {
		if i == n1 {
			part1 = chamber.highestRocks()
		}
		moved := true
		highestPoint := chamber.highestRocks()

		chamber.extendWalls(highestPoint)
		piece := rock.Pieces[i%len(rock.Pieces)](highestPoint)
		for _, co := range piece {
			chamber[co] = true
		}
		for moved {
			inst := string(instructions[j%len(instructions)])
			switch inst {
			case "<":
				piece, _ = chamber.move(piece, -1, 0)
			case ">":
				piece, _ = chamber.move(piece, 1, 0)
			default:
				return -1, -1, fmt.Errorf("invalid jet instruction given: %s", inst)
			}
			piece, moved = chamber.move(piece, 0, -1)
			if !moved {
				newBase := maths.Infinity
				for _, co := range piece {
					chamber[co] = true
					xTaken[co.X] = co.Y
					if base != 0 {
						newBase = maths.Min(co.Y, newBase)
					}
				}
				base = newBase

				hp := chamber.highestPoints()
				key := [2]int{i % len(rock.Pieces), j % len(instructions)}

				if comp, ok := cache[key]; ok && hp == comp.points && i > n1 {
					heightDiff := chamber.highestRocks() - comp.height
					cycle := i - comp.rock
					multiplier := (n2 - i) / cycle
					newChamber := Chamber{}
					for co := range chamber {
						if co.X > 0 && co.X < 8 && co.Y > 0 {
							newChamber[graph.Co{X: co.X, Y: co.Y + (heightDiff * multiplier)}] = true
						}
					}
					chamber = newChamber
					i += cycle * multiplier
					cache = make(map[[2]int]state)
				} else {
					cache[key] = state{
						points: hp,
						rock:   i,
						height: chamber.highestRocks(),
					}
				}
			}
			j++
		}
	}
	return part1, chamber.highestRocks(), nil
}

func main() {
	input := file.Read()
	part1, part2, err := playTetris(input[0], 2022, 1000000000000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
