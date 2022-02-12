package main

import (
	ship "Advent-of-Code/2020/Day12/ship"
	waypoint "Advent-of-Code/2020/Day12/waypoint"
	"Advent-of-Code/file"
	"fmt"
	"strconv"
)

func parseDirection(entry string) (string, int, error) {
	dir := string(entry[0])
	val, err := strconv.Atoi(entry[1:])
	return dir, val, err
}

func part1(entries []string) (int, error) {
	s := ship.Ship{
		FacingDir: "E",
	}
	for _, entry := range entries {
		dir, val, err := parseDirection(entry)
		if err != nil {
			return 0, err
		}
		s.MoveShip(dir, val)
	}
	return s.CalculateDistance(), nil
}

func part2(entries []string) (int, error) {
	s := ship.Ship{
		FacingDir: "E",
	}
	wp := waypoint.Waypoint{
		X: 10,
		Y: 1,
	}
	for _, entry := range entries {
		dir, val, err := parseDirection(entry)
		if err != nil {
			return 0, err
		}
		wp.MoveWaypoint(&s, dir, val)
	}
	return s.CalculateDistance(), nil
}

func main() {
	entries := file.Read()
	part1Sol, err := part1(entries)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1Sol)

	part2Sol, err := part2(entries)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", part2Sol)
}
