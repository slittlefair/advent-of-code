package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Co is a simple struct for a graph coordinate with points x, y
type Co struct {
	X int
	Y int
}

// Check handles any errors we have
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// StringToInt takes a string (hopefully of numbers), converts it to int and handles any errors
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

// StringSliceToIntSlice takes a slice of strings and returns the same slice but of ints
func StringSliceToIntSlice(ss []string) (is []int) {
	for _, val := range ss {
		is = append(is, StringToInt(val))
	}
	return is
}

// ReadFile takes lines in a txt files and put them in an array
func ReadFile() []string {
	file, err := os.Open("./input.txt")
	Check(err)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// ReadFileAsInts takes lines in a text file and puts them into an array as integers
func ReadFileAsInts() []int {
	file, err := os.Open("./input.txt")
	Check(err)
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, StringToInt(scanner.Text()))
	}
	return lines
}

// Abs returns the absolute value of the int provided
func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// StringInSlice returns a boolean based on if the given string is in the given slice
func StringInSlice(s string, sl []string) bool {
	for _, k := range sl {
		if k == s {
			return true
		}
	}
	return false
}

// Permutations returns all ordered permutations of a slice of strings
func Permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

// StringSlicesEqual takes two slices and returns whether they contain the same elements
func StringSlicesEqual(a []string, b []string) (equal bool) {
	if len(a) != len(b) {
		return false
	}
	vals := make(map[string]bool)
	for _, v := range a {
		vals[v] = true
	}
	for _, v := range b {
		if _, ok := vals[v]; !ok {
			return false
		}
	}
	return true
}

// TimeTrack prints how long a function took to run
func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("took %s\n", elapsed)
}
