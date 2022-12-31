package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
)

type Item struct {
	value int
	prev  *Item
	next  *Item
}

type List []*Item

func parseInput(input []int) List {
	list := List{}
	for i := 0; i < len(input); i++ {
		item := &Item{
			value: input[i],
		}
		list = append(list, item)
	}
	for i, itm := range list {
		prev := list[maths.Modulo(i-1, len(list))]
		itm.prev = prev
		next := list[maths.Modulo(i+1, len(list))]
		itm.next = next
	}
	return list
}

func (itm *Item) moveItemInList(listLength int) {
	// Return early if the item's value is 0 as we don't want to move it
	if itm.value == 0 {
		return
	}

	fmt.Println("mod", maths.Modulo(itm.value, listLength-1))

	oldNext := itm.next
	oldPrev := itm.prev

	oldNext.prev = oldPrev
	oldPrev.next = oldNext

	n := itm
	steps := maths.Modulo(itm.value, listLength-1)
	for steps > 0 {
		n = n.next
		steps--
	}

	newPrev := n
	newNext := n.next

	newPrev.next = itm
	itm.prev = newPrev

	newNext.prev = itm
	itm.next = newNext
}

func (l List) mixList(decryptionKey, times int) {
	for _, itm := range l {
		itm.value *= decryptionKey
	}
	for i := 0; i < times; i++ {
		fmt.Println("RUN", i+1)
		for _, itm := range l {
			itm.moveItemInList(len(l))
		}
	}
}

func (l List) getSumOfCoordinates() int {
	sum := 0
	var itm *Item
	for _, i := range l {
		if i.value == 0 {
			itm = i
			break
		}
	}
	for i := 1; i <= 3000; i++ {
		itm = itm.next
		if i%1000 == 0 {
			sum += itm.value
		}
	}
	return sum
}

func findSolutions(input []int) (int, int) {
	list := parseInput(input)
	// list.mixList(1, 1)
	// part1 := list.getSumOfCoordinates()

	// list = parseInput(input)
	list.mixList(811589153, 10)
	part2 := list.getSumOfCoordinates()

	return 0, part2
}

func main() {
	input := file.ReadAsInts()
	part1, part2 := findSolutions(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
