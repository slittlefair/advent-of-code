package main

import (
	"Advent-of-Code/file"
	"fmt"
	"strconv"
)

type Disk struct {
	maxIndex  int
	blockSize int
	blocks    map[int]int
}

func parseInput(input []string) (*Disk, error) {
	disk := &Disk{
		blocks: make(map[int]int),
	}
	id := 0
	index := 0
	for i, c := range input[0] {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			return nil, err
		}
		disk.maxIndex += n
		if i%2 == 0 {
			for j := index; j < index+n; j++ {
				disk.blocks[j] = id
			}
			id++
			disk.blockSize += n
		}
		index += n
	}
	return disk, nil
}

func (d Disk) Print() {
	for i := 0; i < d.maxIndex; i++ {
		if v, ok := d.blocks[i]; ok {
			fmt.Print(v)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func main() {
	input := file.Read()
	disk, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := disk.maxIndex - 1; i >= disk.blockSize; i-- {
		if _, ok := disk.blocks[i]; !ok {
			continue
		}
		for j := 0; j < disk.maxIndex; j++ {
			if _, ok := disk.blocks[j]; !ok {
				disk.blocks[j] = disk.blocks[i]
				delete(disk.blocks, i)
				break
			}
		}
	}

	checksum := 0
	for i := 0; i <= disk.blockSize; i++ {
		checksum += i * disk.blocks[i]
	}
	fmt.Printf("Part1: %v\n", checksum)
}
