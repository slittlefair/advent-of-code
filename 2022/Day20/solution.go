package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
)

// Constantly editing an array, especially with an input as big as it is, is extremely slow as we're
// going to keep moving all items along and changing their indexes. For this problem we don't
// actually care what each item's index is, we only care about which value comes before and after
// it. This means when we move an item we only have to alter the links (prev and next items) of that
// item, its old prev and next items links and its new prev and next items links, we can leave all
// items between its new and old positions alone as their links don't change.
type Item struct {
	value int
	prev  *Item
	next  *Item
}

type List []*Item

func parseInput(input []int) (List, *Item) {
	list := List{}
	var zero *Item
	// Run through the input and create items with those values. We can't populate prev and next
	// values as they don't all exist yet, so we'll do that later.
	for i := 0; i < len(input); i++ {
		val := input[i]
		item := &Item{
			value: val,
		}
		// If the value is 0 then return this item to use when we find the coordinates relative to
		// it.
		if val == 0 {
			zero = item
		}
		list = append(list, item)
	}
	// Run through all the items againa and populate prev and next values
	for i, itm := range list {
		prev := list[maths.Modulo(i-1, len(list))]
		itm.prev = prev
		next := list[maths.Modulo(i+1, len(list))]
		itm.next = next
	}
	return list, zero
}

func (itm *Item) moveItemInList(listLength int) {
	// Return early if the item's value is 0 as we don't want to move it
	if itm.value == 0 {
		return
	}

	// Remove the item from its old position by "connecting" its previous and next values.
	oldNext := itm.next
	oldPrev := itm.prev
	oldNext.prev = oldPrev
	oldPrev.next = oldNext

	// Find the new position of the item. Since we loop around from end to beginning we need the
	// modulo value of the item compared to the length of the list minus 1, since the item has
	// already been removed. The value of n is going to be the item before our moved item, or its
	// prev value.
	n := itm
	for i := 0; i < maths.Modulo(itm.value, listLength-1); i++ {
		n = n.next
	}

	// Since n is going to be the moved item's prev value, the moved item's next value is going to
	// be the item after n, that is its next value.
	newPrev := n
	newNext := n.next

	// Update the prev and next values of itm and its prev and next items.
	newPrev.next = itm
	itm.prev = newPrev
	newNext.prev = itm
	itm.next = newNext
}

func (l List) mixList(decryptionKey, times int) {
	// Before we start mixing the list we need to multiply all list values by the decryptionKey
	for _, itm := range l {
		itm.value *= decryptionKey
	}

	// Move all items in the list the given number of times. Even though the links of the items
	// change their order never does so we can safely keep ranging over it.
	for i := 0; i < times; i++ {
		for _, itm := range l {
			itm.moveItemInList(len(l))
		}
	}
}

func (l List) getSumOfCoordinates(zero *Item) int {
	sum := 0
	itm := zero
	// Keep finding the next item relative to the current one, starting at the item with value 0. We
	// want to sum the values of items 1000, 2000 and 3000 after zero.
	for i := 1; i <= 3000; i++ {
		itm = itm.next
		if maths.Modulo(i, 1000) == 0 {
			sum += itm.value
		}
	}
	return sum
}

func findSolutions(input []int) (int, int) {
	list, zero := parseInput(input)
	list.mixList(1, 1)
	part1 := list.getSumOfCoordinates(zero)

	list, zero = parseInput(input)
	list.mixList(811589153, 10)
	part2 := list.getSumOfCoordinates(zero)

	return part1, part2
}

func main() {
	input := file.ReadAsInts()
	part1, part2 := findSolutions(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
