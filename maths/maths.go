package maths

import "sort"

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

// Infinity is the int value of infinity, useful for looping over a range and trying to get the lowest value
var Infinity = int(^uint(0) >> 1)

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
