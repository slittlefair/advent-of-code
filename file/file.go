package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read takes lines in a txt files and put them in an array
func Read() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// ReadFileAsInts takes lines in a text file and puts them into an array as integers
func ReadAsInts() []int {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		lines = append(lines, t)
	}
	return lines
}

// ReadSingleLineAsInts reads the input file using Read and returns a slice of ints in that file.
// It expects the file to be only one line long and errors if it's not.
func ReadSingleLineAsInts() ([]int, error) {
	input := Read()
	if l := len(input); l != 1 {
		return nil, fmt.Errorf("error getting input, expected 1 line, got %d: %v", l, input)
	}
	ints := strings.Split(input[0], ",")
	nums := []int{}
	for _, i := range ints {
		n, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}
