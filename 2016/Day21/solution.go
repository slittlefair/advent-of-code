package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/slice"
	"fmt"
	"strconv"
	"strings"
)

func swapPosition(password []string, i, j int) []string {
	password[i], password[j] = password[j], password[i]
	return password
}

func swapLetter(password []string, x, y string) []string {
	i := slice.Index(password, x)
	j := slice.Index(password, y)
	return swapPosition(password, i, j)
}

func rotateLeft(password []string, x int) []string {
	r := x % len(password)
	return append(password[r:], password[:r]...)
}

func rotateRight(password []string, x int) []string {
	r := len(password) - x%len(password)
	return append(password[r:], password[:r]...)
}

func rotateByPosition(password []string, x string, rotateFunc func([]string, int) []string) []string {
	i := slice.Index(password, x)
	password = rotateFunc(password, i+1)
	if i > 3 {
		password = rotateFunc(password, 1)
	}
	return password
}

func reverse(password []string, x, y int) []string {
	for i, j := x, y; i < j; i, j = i+1, j-1 {
		password = swapPosition(password, i, j)
	}
	return password
}

func move(password []string, x, y int) []string {
	elem := password[x]
	password = append(password[:x], password[x+1:]...)
	return append(password[:y], append([]string{elem}, password[y:]...)...)
}

func followSwap(password, split []string) ([]string, error) {
	x, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(split[5])
	if err != nil {
		return nil, err
	}
	return swapPosition(password, x, y), nil
}

func followRotate(password, split []string, rotateFunc func([]string, int) []string) ([]string, error) {
	i, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, err
	}
	return rotateFunc(password, i), nil
}

func followReverse(password, split []string) ([]string, error) {
	x, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(split[4])
	if err != nil {
		return nil, err
	}
	return reverse(password, x, y), nil
}

func followMove(password, split []string) ([]string, error) {
	x, err := strconv.Atoi(split[2])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(split[5])
	if err != nil {
		return nil, err
	}
	return move(password, x, y), nil
}

func followInstruction(password []string, line string, reverse bool) ([]string, error) {
	split := strings.Split(line, " ")
	switch fmt.Sprintf("%s %s", split[0], split[1]) {
	case "swap position":
		return followSwap(password, split)
	case "swap letter":
		return swapLetter(password, split[2], split[5]), nil
	case "rotate left":
		if reverse {
			return followRotate(password, split, rotateRight)
		}
		return followRotate(password, split, rotateLeft)
	case "rotate right":
		if reverse {
			return followRotate(password, split, rotateLeft)
		}
		return followRotate(password, split, rotateRight)
	case "rotate based":
		if reverse {
			return rotateByPosition(password, split[6], rotateLeft), nil
		}
		return rotateByPosition(password, split[6], rotateRight), nil
	case "reverse positions":
		return followReverse(password, split)
	case "move position":
		return followMove(password, split)
	default:
		return nil, fmt.Errorf("invalid instruction: %s", line)
	}
}

func scramblePassword(password string, instructions []string) (string, error) {
	p := strings.Split(password, "")
	var err error
	for _, line := range instructions {
		p, err = followInstruction(p, line, false)
		if err != nil {
			return "", err
		}
	}
	return strings.Join(p, ""), nil
}

func unscramblePassword(password string, instructions []string) (string, error) {
	p := strings.Split(password, "")
	var err error
	fmt.Println(p)
	for i := len(instructions) - 1; i >= 0; i-- {
		p, err = followInstruction(p, instructions[i], false)
		if err != nil {
			return "", err
		}
		fmt.Println(p)
	}
	return strings.Join(p, ""), nil
}

func findSolutions(password1, password2 string, input []string) (string, string, error) {
	part1, err := scramblePassword(password1, input)
	if err != nil {
		return "", "", err
	}
	part2, err := unscramblePassword(password2, input)
	return part1, part2, err
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions("abcdefgh", "decab", input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
