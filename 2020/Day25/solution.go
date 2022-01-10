package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"strconv"
)

func transformNumber(value int, subjectNumber int) int {
	value *= subjectNumber
	return value % 20201227
}

func getLoopSize(key int) int {
	loopSize := 0
	value := 1
	for {
		value = transformNumber(value, 7)
		loopSize++
		if value == key {
			return loopSize
		}
	}
}

func getEncryptionKey(subjectNumber int, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value = transformNumber(value, subjectNumber)
	}
	return value
}

func main() {
	input := utils.ReadFile()
	cardKey, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	doorKey, err := strconv.Atoi(input[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	cardLoopSize := getLoopSize(cardKey)
	doorLoopSize := getLoopSize(doorKey)

	encryptionKeyFromCardKey := getEncryptionKey(cardKey, doorLoopSize)
	encryptionKeyFromDoorKey := getEncryptionKey(doorKey, cardLoopSize)

	if encryptionKeyFromCardKey != encryptionKeyFromDoorKey {
		fmt.Println("Error: encryptionkeys don't match", encryptionKeyFromCardKey, encryptionKeyFromDoorKey)
		return
	}

	fmt.Println("Part 1:", encryptionKeyFromCardKey)
}
