package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var aocExampleList = func() List {
	item1 := &Item{
		value: 1,
	}
	item2 := &Item{
		value: 2,
	}
	item3 := &Item{
		value: -3,
	}
	item4 := &Item{
		value: 3,
	}
	item5 := &Item{
		value: -2,
	}
	item6 := &Item{
		value: 0,
	}
	item7 := &Item{
		value: 4,
	}

	item1.prev = item7
	item2.prev = item1
	item3.prev = item2
	item4.prev = item3
	item5.prev = item4
	item6.prev = item5
	item7.prev = item6

	item1.next = item2
	item2.next = item3
	item3.next = item4
	item4.next = item5
	item5.next = item6
	item6.next = item7
	item7.next = item1

	return List{
		item1,
		item2,
		item3,
		item4,
		item5,
		item6,
		item7,
	}
}()

func TestParseInput(t *testing.T) {
	t.Run("parses input into a linked list", func(t *testing.T) {
		input := []int{1, 2, -3, 3, -2, 0, 4}

		got, got1 := parseInput(input)
		assert.Equal(t, aocExampleList, got)
		assert.Equal(t, aocExampleList[5], got1)
	})
}

func TestMoveItemInList(t *testing.T) {
	var checkMoveItemInList = func(t *testing.T, input, want []int, moveVal, moveIdx, startVal int) {
		got, _ := parseInput(input)
		itm := got[moveIdx]
		assert.Equal(t, moveVal, itm.value)

		itm.moveItemInList(len(got))

		itm = got[0]
		assert.Equal(t, startVal, itm.value)
		nums := []int{itm.value}
		for i := 1; i < len(got); i++ {
			itm = itm.next
			nums = append(nums, itm.value)
			assert.Equal(t, want[i], itm.value)
		}
	}

	t.Run("moves item in list advent of code example 1", func(t *testing.T) {
		input := []int{4, 5, 6, 1, 7, 8, 9}
		want := []int{4, 5, 6, 7, 1, 8, 9}
		checkMoveItemInList(t, input, want, 1, 3, 4)
	})

	t.Run("moves item in list advent of code example 2", func(t *testing.T) {
		input := []int{4, -2, 5, 6, 7, 8, 9}
		want := []int{4, 5, 6, 7, 8, -2, 9}
		checkMoveItemInList(t, input, want, -2, 1, 4)
	})
}

func TestMixList(t *testing.T) {
	t.Run("mixes list, advent of code example 1", func(t *testing.T) {
		input := []int{1, 2, -3, 3, -2, 0, 4}
		want := []int{1, 2, -3, 4, 0, 3, -2}

		list, _ := parseInput(input)
		list.mixList(1, 1)
		itm := list[0]
		assert.Equal(t, 1, itm.value)
		for i := 1; i < len(list); i++ {
			itm = itm.next
			assert.Equal(t, want[i], itm.value)
		}
	})

	t.Run("mixes list, advent of code example 1", func(t *testing.T) {
		input := []int{1, 2, -3, 3, -2, 0, 4}
		want := []int{811589153, 0, -2434767459, 1623178306, 3246356612, -1623178306, 2434767459}

		list, _ := parseInput(input)
		list.mixList(811589153, 10)
		itm := list[0]
		assert.Equal(t, 811589153, itm.value)
		for i := 1; i < len(list); i++ {
			itm = itm.next
			assert.Equal(t, want[i], itm.value)
		}
	})
}

func TestGetSumOfCoordinates(t *testing.T) {
	t.Run("returns sum of coordinates from a list, advent of code example 1", func(t *testing.T) {
		input := []int{1, 2, -3, 3, -2, 0, 4}
		list, zero := parseInput(input)
		list.mixList(1, 1)
		got := list.getSumOfCoordinates(zero)
		assert.Equal(t, 3, got)
	})

	t.Run("returns sum of coordinates from a list, advent of code example 2", func(t *testing.T) {
		input := []int{1, 2, -3, 3, -2, 0, 4}
		list, zero := parseInput(input)
		list.mixList(811589153, 10)
		got := list.getSumOfCoordinates(zero)
		assert.Equal(t, 1623178306, got)
	})
}

func TestFindSolutions(t *testing.T) {
	t.Run("returns solution for given input, advent of code example", func(t *testing.T) {
		input := []int{1, 2, -3, 3, -2, 0, 4}
		got, got1 := findSolutions(input)
		assert.Equal(t, 3, got)
		assert.Equal(t, 1623178306, got1)
	})
}
