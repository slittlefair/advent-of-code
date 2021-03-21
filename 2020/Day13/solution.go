package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	entries := helpers.ReadFile()
	arrivalTime, err := strconv.Atoi(entries[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	re := regexp.MustCompile(`\d+`)
	busStrings := re.FindAllString(entries[1], -1)
	buses := []int{}
	for _, bs := range busStrings {
		bus, err := strconv.Atoi(bs)
		if err != nil {
			fmt.Println(err)
			return
		}
		buses = append(buses, bus)
	}

	smallestTimeToWait := 100000
	var busToWaitFor int
	for _, bus := range buses {
		timeToWait := bus - (arrivalTime % bus)
		if timeToWait < smallestTimeToWait {
			smallestTimeToWait = timeToWait
			busToWaitFor = bus
		}
	}
	fmt.Println("Part 1:", busToWaitFor*smallestTimeToWait)
}
