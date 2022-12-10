package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/slice"
	"fmt"
)

type Knot struct {
	co     graph.Co
	parent *Knot
}

type Rope struct {
	visited map[graph.Co]bool
	knots   map[int]*Knot
}

// PrintGrid prints a grid similar to AoC's, except "H" is "0" here instead. It is unused but was
// useful for debugging so has been retained here for reference.
//
// For part 1 example use r.PrintGrid(0, 5, -4, 0)
// For large part 2 example use r.PrintGrid(-11, 14, -15, 5)
func (r Rope) PrintGrid(minX, maxX, minY, maxY int) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			for i := 0; i < len(r.knots); i++ {
				if r.knots[i].co.X == x && r.knots[i].co.Y == y {
					fmt.Print(i)
					goto out
				}
			}
			if x == 0 && y == 0 {
				fmt.Print("s")
			} else {
				fmt.Print(".")
			}

		out:
		}
		fmt.Println()
	}
	fmt.Println()
}

func (r *Rope) moveRope(inst string) error {
	var dir string
	var dist int
	_, err := fmt.Sscanf(inst, "%s %d", &dir, &dist)
	if err != nil {
		return err
	}

	for i := 0; i < dist; i++ {
		// Move the head of the rope in the direction given by the instruction
		switch dir {
		case "U":
			r.knots[0].co.Y--
		case "D":
			r.knots[0].co.Y++
		case "L":
			r.knots[0].co.X--
		case "R":
			r.knots[0].co.X++
		default:
			return fmt.Errorf("invalid instruction %s", inst)
		}

		// Move the rest of the knots based on the position of the previous knot (their parent)
		for i := 1; i < len(r.knots); i++ {
			curKnot := r.knots[i]
			curCo := curKnot.co
			parCo := curKnot.parent.co

			// Get a slice of all adjacent possible knot positions, including the current position
			// since overlapping is valid.
			validCos := append([]graph.Co{curCo}, graph.AdjacentCos(curCo, true)...)

			// If the previous knot is not adjacent then it also needs to move.
			if !slice.Contains(validCos, parCo) {
				// If the current knot's x axis value is different to the previous knot's, move it
				if parCo.X > curCo.X {
					curKnot.co.X++
				} else if parCo.X < curCo.X {
					curKnot.co.X--
				}

				// If the current knot's y axis value is different to the previous knot's, move it
				if parCo.Y > curCo.Y {
					curKnot.co.Y++
				} else if parCo.Y < curCo.Y {
					curKnot.co.Y--
				}
			}
		}
		// Keep a map of all coordinates the tail has been in
		r.visited[r.knots[len(r.knots)-1].co] = true
	}
	return nil
}

func (r *Rope) followInstructions(input []string) error {
	for _, line := range input {
		err := r.moveRope(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func makeRope(n int) Rope {
	knots := map[int]*Knot{
		0: {},
	}
	for i := 1; i < n; i++ {
		knots[i] = &Knot{parent: knots[i-1]}
	}
	return Rope{
		visited: make(map[graph.Co]bool),
		knots:   knots,
	}
}

func main() {
	input := file.Read()
	r1 := makeRope(2)
	err := r1.followInstructions(input)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", len(r1.visited))

	r2 := makeRope(10)
	// If the instructions didn't produce an error for r1 then they won't for r2, so we can safely
	// ignore the returned error here, knowing it's nil
	_ = r2.followInstructions(input)
	fmt.Println("Part 2:", len(r2.visited))
}
