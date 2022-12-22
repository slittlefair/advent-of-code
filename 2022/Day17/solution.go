package main

import (
	chmb "Advent-of-Code/2022/Day17/chamber"
	"Advent-of-Code/2022/Day17/rock"
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/timer"
	"fmt"
	"time"
)

type State struct {
	colHeights [7]int
	rockNum    int
	height     int
}

// Cache keeps a record of the state of the chamber at each jet/piece combination. We use this to
// see if we've seen the current state before and if so, we can "jump" forward a number of pieces as
// we know we'll see this formation again.
type Cache map[[2]int]State

func playTetris(instructions string, n1, n2 int) (int, int, error) {
	chamber := chmb.CreateChamber()
	jet := 0
	cache := Cache{}
	var part1 int
	for rockNum := 0; rockNum < n2; rockNum++ {
		// Keep a record of the height of the rocks when we reach part 1 to return later.
		if rockNum == n1 {
			part1 = chamber.HighestPoint()
		}

		// Find the highest point and extend the side walls so the next rock can move accurately.
		highestPoint := chamber.HighestPoint()
		chamber.ExtendWalls(highestPoint)

		// Drop the next piece into the chamber
		pieceIdx := rockNum % len(rock.Pieces)
		piece := rock.Pieces[pieceIdx](highestPoint)
		for _, co := range piece {
			chamber[co] = true
		}

		// Start moving the piece and keep doing so until it comes to rest
		moved := true
		for moved {
			inst := string(instructions[jet])
			// Move left or right if it can. We don't care if it's moved or not so ignore that return
			switch inst {
			case "<":
				piece, _ = chamber.Move(piece, -1, 0)
			case ">":
				piece, _ = chamber.Move(piece, 1, 0)
			default:
				return -1, -1, fmt.Errorf("invalid jet instruction given: %s", inst)
			}

			// Move the piece down if it can
			piece, moved = chamber.Move(piece, 0, -1)

			// If it's not moved down then it's come to rest
			if !moved {
				key := [2]int{pieceIdx, jet}
				ch := chamber.GetColumnHeightsAndMinimiseChamber()

				// Check the cache to see the state the last time we saw this rockIdx/jet
				// combination. As the length of rock pieces and jet instructions are prime then if
				// the highest points are the same for that key as the last time then we've found a
				// cycle that will repeat. In this case we jump that many cycles by increasing the
				// rockNum as well as all coordinates of rocks in the chamber to simulate the
				// heights as if we'd continued. Since we're removing unnecessary coordinates as we
				// go this is quick. We then resume dropping rocks and following instructions until
				// we reach our goal.
				if comp, ok := cache[key]; ok && ch == comp.colHeights && rockNum > n1 {
					heightDiff := chamber.HighestPoint() - comp.height
					cycleLength := rockNum - comp.rockNum
					numCycles := (n2 - rockNum) / cycleLength

					newChamber := chmb.Chamber{}
					for co := range chamber {
						newChamber[graph.Co{X: co.X, Y: co.Y + (heightDiff * numCycles)}] = true
					}
					chamber = newChamber

					rockNum += cycleLength * numCycles
				} else {
					cache[key] = State{
						colHeights: ch,
						rockNum:    rockNum,
						height:     chamber.HighestPoint(),
					}
				}
			}

			jet = (jet + 1) % len(instructions)
		}
	}
	return part1, chamber.HighestPoint(), nil
}

func main() {
	t := time.Now()
	input := file.Read()
	part1, part2, err := playTetris(input[0], 2022, 1000000000000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
	timer.Track(t)
}
