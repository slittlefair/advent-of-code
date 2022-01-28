package slice

import (
	"Advent-of-Code/maths"
	"strconv"
)

// StringSliceToIntSlice takes a slice of strings and returns the same slice but of ints
func StringSliceToIntSlice(ss []string) ([]int, error) {
	is := []int{}
	for _, val := range ss {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		is = append(is, i)
	}
	return is, nil
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
func StringSlicesContainSameElements(a []string, b []string) (equal bool) {
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

// IntSlicesAreEqual returns a bool depending on whether the given slices are equal.
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

// Remove removes the element at index i from slice s and returns that slice, whilst keeping the original in tact
func Remove(s []int, i int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:i]...)
	return append(ret, s[i+1:]...)
}

// FindExtremities returns the max and min value from a slice of ints
func FindExtremities(nums []int) (int, int) {
	min := maths.Infinity
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
