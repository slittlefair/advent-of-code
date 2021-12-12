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

// Min returns the minimum of two ints
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max returns the maximum of two ints
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
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

// IntSlicesAreEqual returns a bool depending on whether the given slices are equal,
func IntSlicesAreEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, num := range slice1 {
		if slice2[i] != num {
			return false
		}
	}
	return true
}

// Infinity is the int value of infinity, useful for looping over a range and trying to get the lowest value
var Infinty = int(^uint(0) >> 1)

// Remove removes the element at index i from slice s and returns that slice, whilst keeping the original in tact
func Remove(s []int, i int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:i]...)
	return append(ret, s[i+1:]...)
}

// CalculateManhattanDistance calculates the manhattan distance between the origin
func CalculateManhattanDistance(co1, co2 Co) int {
	x := co1.X - co2.X
	y := co1.Y - co2.Y
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

// CaeserCipher applies a Caeser Cipher to the given text, shifted shiftNum times
func CaesarCipher(text string, shiftNum int) string {
	shift, offset := rune(shiftNum%26), rune(26)

	runes := []rune(text)

	for index, char := range runes {
		if char >= 'a' && char <= 'z'-shift ||
			char >= 'A' && char <= 'Z'-shift {
			char = char + shift
		} else if char > 'z'-shift && char <= 'z' ||
			char > 'Z'-shift && char <= 'Z' {
			char = char + shift - offset
		}

		// Above handles both upper and lower case ASCII
		// characters; anything else is returned as is (includes
		// numbers, punctuation and space).
		runes[index] = char
	}

	return string(runes)
}
