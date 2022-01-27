package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
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

func ReadFileSingleLineAsInts() ([]int, error) {
	input := ReadFile()
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
var Infinity = int(^uint(0) >> 1)

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

// CaesarCipher applies a Caesar Cipher to the given text, shifted shiftNum times
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

// Median returns the median value from an unsorted slice of ints
func Median(nums []int) float64 {
	sort.Ints(nums)
	l := len(nums)
	if l%2 != 0 {
		return float64(nums[(l-1)/2])
	}
	midWay1 := float64(nums[l/2])
	midWay2 := float64(nums[(l/2)-1])
	return (midWay1 + midWay2) / 2
}

// FindExtremities returns the max and min value from a slice of ints
func FindExtremities(nums []int) (int, int) {
	min := Infinity
	max := 0
	for _, n := range nums {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return min, max
}

// AdjacentCos returns all adjacent coordinates for the given coordinate, including diagonals
func AdjacentCos(co Co, includeDiagonals bool) []Co {
	cos := []Co{
		{X: co.X, Y: co.Y - 1},
		{X: co.X - 1, Y: co.Y},
		{X: co.X + 1, Y: co.Y},
		{X: co.X, Y: co.Y + 1},
	}
	if !includeDiagonals {
		return cos
	}
	return append(cos,
		Co{X: co.X - 1, Y: co.Y - 1},
		Co{X: co.X + 1, Y: co.Y - 1},
		Co{X: co.X - 1, Y: co.Y + 1},
		Co{X: co.X + 1, Y: co.Y + 1},
	)
}

// IsUpper returns true if all characters in a string are upper case, false otherwise
func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return false
		}
	}
	return true
}

// IsLower returns true if all characters in a string are lower case, false otherwise
func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
