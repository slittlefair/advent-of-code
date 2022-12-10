package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"Advent-of-Code/strings"
	"fmt"
)

type Grid map[graph.Co]string

type cpu struct {
	cycle  int
	x      int
	co     graph.Co
	signal int
	pixels map[graph.Co]string
}

func (cpu *cpu) checkCycle() {
	cpu.cycle++

	// Part 1
	if (cpu.cycle+20)%40 == 0 {
		cpu.signal += cpu.cycle * cpu.x
	}

	// Part 2
	if maths.Abs(cpu.co.X-cpu.x) <= 1 {
		cpu.pixels[cpu.co] = strings.BLOCK
	} else {
		cpu.pixels[cpu.co] = " "
	}

	// If we're at the end of a row, move down to the start of the next
	cpu.co.X++
	if cpu.co.X == 40 {
		cpu.co.X = 0
		cpu.co.Y++
	}
}

func (cpu *cpu) handleInstruction(inst string) error {
	var v int
	if inst == "noop" {
		cpu.checkCycle()
	} else {
		_, err := fmt.Sscanf(inst, "addx %d", &v)
		if err != nil {
			return err
		}
		for i := 0; i < 2; i++ {
			cpu.checkCycle()
		}
	}
	cpu.x += v
	return nil
}

func (cpu *cpu) completeCycles(input []string) error {
	for _, inst := range input {
		err := cpu.handleInstruction(inst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cpu cpu) printGrid() {
	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			fmt.Print(cpu.pixels[graph.Co{X: x, Y: y}])
		}
		fmt.Println()
	}
}

func main() {
	input := file.Read()
	cpu := cpu{x: 1, pixels: make(map[graph.Co]string)}
	err := cpu.completeCycles(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", cpu.signal)
	fmt.Println("Part 2:")
	cpu.printGrid()
}
