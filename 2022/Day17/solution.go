package main

import (
	"Advent-of-Code/2022/Day17/rock"
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"Advent-of-Code/slice"
	"Advent-of-Code/timer"
	"fmt"
	"time"
)

type Chamber map[graph.Co]string

func (c Chamber) PrintChamber(animate, overwrite bool) {
	highest := c.highestRocks()
	for y := highest + 10; y > 0; y-- {
		fmt.Print("|")
		for x := 1; x <= 7; x++ {
			co := graph.Co{X: x, Y: y}
			if v, ok := c[co]; ok {
				fmt.Print(v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|", y)
		fmt.Println()
	}
	fmt.Println("+-------+")
}

// Create the initial chamber, we will extend the side walls as and when
func createChamber() Chamber {
	chamber := Chamber{}
	for x := 0; x <= 8; x++ {
		chamber[graph.Co{X: x, Y: 0}] = "#"
	}
	for y := 0; y <= 10; y++ {
		chamber[graph.Co{X: 0, Y: y}] = "#"
		chamber[graph.Co{X: 8, Y: y}] = "#"
	}
	return chamber
}

func (c Chamber) extendWalls(n int) {
	for y := n; y <= n+10; y++ {
		c[graph.Co{X: 0, Y: y}] = "#"
		c[graph.Co{X: 8, Y: y}] = "#"
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

const (
	Left  = -1
	Right = 1
)

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
		c[co] = "@"
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
			fmt.Println("part1!!")
			part1 = chamber.highestRocks()
		}
		if i%100 == 0 {
			fmt.Printf("%d down, %d to go...\n", i, n2-i)
		}
		moved := true
		highestPoint := chamber.highestRocks()
		// for y := 1; y <= highestPoint; y++ {
		// 	count := 0
		// 	for x := 1; x <= 7; x++ {
		// 		if _, ok := chamber[graph.Co{X: x, Y: y}]; ok {
		// 			count++
		// 		}
		// 	}
		// 	if count == 7 {
		// 		fmt.Println("TETRIS!", i, y, len(chamber))
		// 		chamber.minimiseChamber(y)
		// 		fmt.Println(len(chamber))
		// 	}
		// }

		chamber.extendWalls(highestPoint)
		piece := rock.Pieces[i%len(rock.Pieces)](highestPoint)
		for _, co := range piece {
			chamber[co] = "@"
		}
		for moved {
			inst := string(instructions[j%len(instructions)])
			// fmt.Println("AAA", j, inst)
			// chamber.PrintChamber(false, false)
			switch inst {
			case "<":
				piece, _ = chamber.move(piece, -1, 0)
			case ">":
				piece, _ = chamber.move(piece, 1, 0)
			default:
				return -1, -1, fmt.Errorf("invalid jet instruction given: %s", inst)
			}
			// fmt.Println("CCC", j, inst)
			// chamber.PrintChamber(false, false)
			piece, moved = chamber.move(piece, 0, -1)
			// fmt.Println("DDD", j, inst)
			// chamber.PrintChamber(false, false)
			if !moved {
				newBase := maths.Infinity
				for _, co := range piece {
					chamber[co] = "#"
					xTaken[co.X] = co.Y
					if base != 0 {
						newBase = maths.Min(co.Y, newBase)
					}
				}
				base = newBase
				// if len(xTaken) == 7 {
				// 	// fmt.Println("TETRIS!", i, len(chamber), xTaken)
				// 	// chamber.PrintChamber(false, false)
				// 	chamber.minimiseChamber(xTaken)
				// 	// fmt.Println(len(chamber))
				// 	// chamber.PrintChamber(false, false)
				// 	// for co := range chamber {
				// 	// 	fmt.Println(co)
				// 	// }
				// 	base = 0
				// 	xTaken = map[int]int{}
				// }
				hp := chamber.highestPoints()
				// chamber.minimiseChamber(hp)
				// fmt.Println("XXX", hp, i%len(rock.Pieces), j%len(instructions))
				// fmt.Println("EEE", j, inst)
				// chamber.PrintChamber(false, false)
				key := [2]int{i % len(rock.Pieces), j % len(instructions)}
				if comp, ok := cache[key]; ok && hp == comp.points && i > n1 {
					fmt.Println("GOT A MATCH!!!!!")
					heightDiff := chamber.highestRocks() - comp.height
					cycle := i - comp.rock
					fmt.Println("hd, cycle", heightDiff, cycle)
					multiplier := (n2 - i) / cycle
					fmt.Println("multiplier", multiplier)
					fmt.Println("xxx", chamber.highestRocks())
					newChamber := Chamber{}
					// chamber.PrintChamber(false, false)
					for co := range chamber {
						if co.X > 0 && co.X < 8 && co.Y > 0 {
							newChamber[graph.Co{X: co.X, Y: co.Y + (heightDiff * multiplier)}] = "#"
						}
					}
					chamber = newChamber
					fmt.Println("xxx", chamber.highestRocks())
					fmt.Println("i was", i)
					i += cycle * multiplier
					fmt.Println("i is", i)
					fmt.Println("increased!!", i, n2-i)
					cache = make(map[[2]int]state)
					// fmt.Println("EEE", j, inst)
					// chamber.PrintChamber(false, false)
					// return chamber.highestRocks(), nil
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
	t := time.Now()
	part1, part2, err := playTetris(input[0], 2022, 1000000000000)
	timer.Track(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
