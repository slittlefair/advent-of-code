package main

import (
	"Advent-of-Code/file"
	"fmt"
)

func generateDataStep(input string) (string, error) {
	new := input + "0"
	reverse := []byte{}
	for i := len(input) - 1; i >= 0; i-- {
		char := input[i]
		switch char {
		case '0':
			reverse = append(reverse, '1')
		case '1':
			reverse = append(reverse, '0')
		default:
			return "", fmt.Errorf("invalid character: %s", string(char))
		}
	}
	return new + string(reverse), nil
}

func generateData(input string, requiredLength int) (string, error) {
	var err error
	for len(input) < requiredLength {
		input, err = generateDataStep(input)
		if err != nil {
			return "", err
		}
	}
	return input[:requiredLength], nil
}

func calculateChecksum(data string) string {
	for {
		checksum := []byte{}
		for i := 0; i < len(data)-1; i += 2 {
			if data[i] == data[i+1] {
				checksum = append(checksum, '1')
			} else {
				checksum = append(checksum, '0')
			}
		}
		if len(checksum)%2 != 0 {
			return string(checksum)
		}
		data = string(checksum)
	}
}

func findSolution(input string, requiredDataLength int) (string, error) {
	data, err := generateData(input, requiredDataLength)
	if err != nil {
		return "", err
	}
	return calculateChecksum(data), nil
}

func findSolutions(input string, part1Length, part2Length int) (string, string, error) {
	part1, err := findSolution(input, part1Length)
	if err != nil {
		return "", "", err
	}
	part2, _ := findSolution(input, part2Length)
	return part1, part2, nil
}

func main() {
	input := file.Read()[0]
	part1, part2, err := findSolutions(input, 272, 35651584)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
