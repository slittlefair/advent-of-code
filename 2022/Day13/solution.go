package main

import (
	"Advent-of-Code/file"
	"encoding/json"
	"fmt"
)

// Determines whether two elements are in the correct order. We return a negative integer if they
// are in the correct order, a positive integer if they are in the incorrect order or zero if the
// elements are identical in this comparison.
func comparePackets(pkt1, pkt2 interface{}) int {
	left, pkt1IsArray := pkt1.([]interface{})
	right, pkt2IsArray := pkt2.([]interface{})

	if !pkt1IsArray && !pkt2IsArray {
		// If neither element is an array then both are ints, so compare their values
		return int(pkt1.(float64)) - int(pkt2.(float64))
	}

	// We either have two arrays or an array and an int, so if we have any ints we need to convert
	// it to an array containing itself
	if !pkt1IsArray {
		left = []interface{}{pkt1}
	}
	if !pkt2IsArray {
		right = []interface{}{pkt2}
	}

	// Run through each element of left and right arrays and compare them. If we have determined an
	// order, return it, otherwise move onto the next element
	for i := 0; i < len(left) && i < len(right); i++ {
		if correctOrder := comparePackets(left[i], right[i]); correctOrder != 0 {
			return correctOrder
		}
	}

	// After running through each element we can make a determination based on number of elements.
	// If the arrays are the same length then we haven't made a decision so we can keep going.
	return len(left) - len(right)
}

func findSolutions(input []string) (int, int, error) {
	sumCorrectIndices := 0
	pair := 0
	allPackets := []interface{}{}
	for i := 0; i < len(input)-1; i = i + 3 {
		pair++
		leftInput := input[i]
		rightInput := input[i+1]

		left := []interface{}{}
		right := []interface{}{}
		err := json.Unmarshal([]byte(leftInput), &left)
		if err != nil {
			return -1, -1, err
		}

		err = json.Unmarshal([]byte(rightInput), &right)
		if err != nil {
			return -1, -1, err
		}

		allPackets = append(allPackets, left, right)

		if comparePackets(left, right) <= 0 {
			sumCorrectIndices += pair
		}
	}

	// We don't have to sort any packets, we just need to determine at what position the two divider
	// packets will be. We can do this by comparing them against all other packets and keeping a
	// total of how many are "less than" them to determine their position in a total list. Once
	// we know that we need to compare them against each other, which we can do by determining
	// which has the larger index and adding 1 to it (if its larger than the other divider packet
	// then its index will be one more).
	var packetIndex2, packetIndex6 = 1, 1
	for _, pkt := range allPackets {
		if comparePackets(pkt, []interface{}{[]interface{}{float64(2)}}) <= 0 {
			packetIndex2++
		}
		if comparePackets(pkt, []interface{}{[]interface{}{float64(6)}}) <= 0 {
			packetIndex6++
		}
	}

	if packetIndex2 < packetIndex6 {
		packetIndex6++
	} else {
		packetIndex2++
	}

	return sumCorrectIndices, packetIndex2 * packetIndex6, nil
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
