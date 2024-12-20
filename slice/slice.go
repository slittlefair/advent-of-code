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

// Contains returns a boolean based on if the given elements is in the given slice
func Contains[K comparable](s []K, v K) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}
	return false
}

// Index returns the index of the given element in the given slice. It returns -1 if
// element is not in the slice.
func Index[K comparable](sl []K, s K) int {
	for i, k := range sl {
		if k == s {
			return i
		}
	}
	return -1
}

// Permutations returns all ordered permutations of a slice
func Permutations[K comparable](arr []K) [][]K {
	var helper func([]K, int)
	res := [][]K{}

	helper = func(arr []K, n int) {
		if n == 1 {
			tmp := make([]K, len(arr))
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

// SameElements returns whether the two gives slices contain the same elements
func SameElements[K comparable](a []K, b []K) bool {
	if len(a) != len(b) {
		return false
	}
	vals := make(map[K]bool)
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

func Equal[K comparable](slice1, slice2 []K) bool {
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
func Remove[K comparable](s []K, i int) []K {
	ret := make([]K, 0)
	ret = append(ret, s[:i]...)
	return append(ret, s[i+1:]...)
}

// FindExtremities returns the max and min value from a slice of ints
func FindExtremities(nums []int) (int, int) {
	minimum := maths.Infinity
	maximum := 0
	for _, n := range nums {
		if n > maximum {
			maximum = n
		}
		if n < minimum {
			minimum = n
		}
	}
	return minimum, maximum
}
