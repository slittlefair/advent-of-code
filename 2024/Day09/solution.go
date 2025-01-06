package main

import (
	"Advent-of-Code/file"
	"fmt"
	"strconv"
)

type Disk []int

// Take a line of input and return a Disk, a slice of ints where each value is the id of the block
// at that index, or a value of -1 in the case of a space.
func parseInput(input []string) (Disk, error) {
	intConv := []int{}
	// First convert the string to a slice of ints. Keep track of the blockSum, which is the sum
	// of all the ints in the input, so we know what the length of the desk needs to be
	blockSum := 0
	for _, c := range input[0] {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			return nil, err
		}
		intConv = append(intConv, n)
		blockSum += n
	}

	// Setting up the slice with a capacity and length allows us to add by assignment, rather than
	// appending, which is slightly more expensive
	disk := make([]int, blockSum, blockSum)
	id := 0
	index := 0
	for i, n := range intConv {
		// Every other number is a block, in which case assign the id in the disk
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				disk[index+j] = id
			}
			id++
		} else {
			for j := 0; j < n; j++ {
				disk[index+j] = -1
			}
		}
		index += n
	}
	return disk, nil
}

// copy returns a copy of a disk so we can manipulate the disk in part 1 and not have to parseInput
// again for part 2
func (d Disk) copy() Disk {
	disk := make(Disk, len(d), len(d))
	for i, v := range d {
		disk[i] = v
	}
	return disk
}

// Print is a utility function which prints out the disk as a single line, printing id for blocks
// and "." for spaces
func (d Disk) Print() {
	for _, v := range d {
		if v == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Println()
}

// Compact a disk for part 1
func (d Disk) compact() {
	// Go through the disk from end to beginning, moving each block we find to the first (nearest
	// the beginning) space we find.
	for i := len(d) - 1; i >= 0; i-- {
		// Continue if it's a space
		if d[i] == -1 {
			continue
		}
		// Go through the disk until we find a space. Set the value at the space's index to be the
		// id of the block and set the value of the block to be a space (-1). Then move to the next
		// block.
		// We only need to look for spaces up to the index of the current block as we can only move
		// it in one direction.
		for j := 0; j < i; j++ {
			if d[j] == -1 {
				d[j] = d[i]
				d[i] = -1
				break
			}
		}
	}
}

// Compact a disk for part 2
func (d Disk) compact2() {
	// Cycle through a disk from end to beginning. When we find a block keep a record of its indexes
	// and then try and find a space (starting at the beginning) where it could fit.
	for i := len(d) - 1; i >= 0; i-- {
		id := d[i]
		// If we find space, continue
		if id == -1 {
			continue
		}
		// Keep a slice of indexes where the current block is situated, ending when we reach a value
		// that's not the current block (different id)
		blocks := []int{}
		for k := i; k >= 0; k-- {
			if d[k] == id {
				blocks = append(blocks, k)
			} else {
				break
			}
		}
		// Look for spaces to move the block to, starting at the beginning of the disk
		for j := 0; j < i; j++ {
			// Skip non-spaces
			if d[j] != -1 {
				continue
			}
			// Keep a slice of indexes of continuous spaces
			spaces := []int{}
			for k := j; k < i; k++ {
				if d[k] == -1 {
					spaces = append(spaces, k)
				} else {
					break
				}
			}
			// If the length of block indexes isn't greater than the length of adjacent spaces then
			// there's enough room to move the block.
			if len(blocks) <= len(spaces) {
				// Change the disk value at the space indexes to be the id of the block, only as
				// many indexes as the block takes up
				for i, k := range spaces {
					if i >= len(blocks) {
						break
					}
					d[k] = id
				}
				// Change the disk values of the where the block was to be a space
				for _, k := range blocks {
					d[k] = -1
				}
				break
			}
		}
		// Once we've evaluated a block we can descrease i by the length of the block-1. This
		// ensures we move onto a different block rather than trying to evaluate the same block
		// block again minus teh element we just started at.
		i -= len(blocks) - 1
	}
}

func (d Disk) calculateChecksum() int {
	checksum := 0
	for i := 0; i < len(d); i++ {
		n := d[i]
		if n == -1 {
			continue
		}
		checksum += i * n
	}
	return checksum
}

func findSolutions(input []string) (int, int, error) {
	part1 := 0
	part2 := 0
	disk, err := parseInput(input)
	if err != nil {
		return part1, part2, err
	}
	// Copy the disk before manipulation for part 2
	diskCopy := disk.copy()

	disk.compact()
	part1 = disk.calculateChecksum()

	diskCopy.compact2()
	part2 = diskCopy.calculateChecksum()

	return part1, part2, nil
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
