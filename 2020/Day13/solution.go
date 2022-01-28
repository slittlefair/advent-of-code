package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
	"regexp"
	"strconv"
)

type Bus struct {
	id     int
	offset int
}

type Buses []Bus

func parseInput(entries []string) (int, Buses, error) {
	re := regexp.MustCompile(`\w+`)
	arrivalTime, err := strconv.Atoi(entries[0])
	if err != nil {
		return 0, nil, err
	}
	busStrings := re.FindAllString(entries[1], -1)
	buses := []Bus{}
	for offset, id := range busStrings {
		bus, err := strconv.Atoi(id)
		if err != nil {
			// We don't want to raise the error here, in fact this allows us to skip "x" buses
			continue
		}
		buses = append(buses, Bus{
			id:     bus,
			offset: offset,
		})
	}

	return arrivalTime, buses, nil
}

func (b *Buses) part1(arrivalTime int) int {
	smallestTimeToWait := maths.Infinity
	var busToWaitFor int
	for _, bus := range *b {
		timeToWait := bus.id - (arrivalTime % bus.id)
		if timeToWait < smallestTimeToWait {
			smallestTimeToWait = timeToWait
			busToWaitFor = bus.id
		}
	}
	return busToWaitFor * smallestTimeToWait
}

func (b *Buses) part2() int {
	t := 0
	runningT := 1
	for _, bus := range *b {
		for (t+bus.offset)%bus.id != 0 {
			t += runningT
		}
		runningT *= bus.id
	}
	return t
}

func main() {
	entries := file.Read()
	arrivalTime, buses, err := parseInput(entries)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", buses.part1(arrivalTime))
	fmt.Println("Part 2:", buses.part2())
}
