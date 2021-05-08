package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strconv"
)

type Cups struct {
	CurrentCupIndex int
	Nums            []int
}

func (c Cups) seenNumsBefore(seenNums [][]int) bool {
	for _, seen := range seenNums {
		if helpers.IntSlicesAreEqual(seen, c.Nums) {
			return true
		}
	}
	return false
}

func (c Cups) indexOfNum(n int) int {
	for i, val := range c.Nums {
		if val == n {
			return i
		}
	}
	return -1
}

func (c Cups) getOrderString() (string, error) {
	indexOf1 := c.indexOfNum(1)
	if indexOf1 == -1 {
		return "", fmt.Errorf("could not find 1 in nums %v", c.Nums)
	}
	numString := ""
	for i := indexOf1 + 1; i < indexOf1+len(c.Nums); i++ {
		str := strconv.Itoa(c.Nums[i%len(c.Nums)])
		numString += str
	}
	return numString, nil
}

func (c Cups) getDestinationIndex() int {
	highestCup := 0
	lowestCup := 100000000000
	for _, num := range c.Nums {
		if num > highestCup {
			highestCup = num
		}
		if num < lowestCup {
			lowestCup = num
		}
	}
	destination := c.Nums[c.CurrentCupIndex] - 1

	for {
		if destination < lowestCup {
			destination = highestCup
		}
		destinationIndex := c.indexOfNum(destination)
		if destinationIndex != -1 {
			return destinationIndex
		}
		destination -= 1
	}
}

func (c *Cups) shiftCups() {
	newNums := []int{}
	for i := 3; i < 3+len(c.Nums); i++ {
		newNums = append(newNums, c.Nums[i%len(c.Nums)])
	}
	c.Nums = newNums
}

func (c *Cups) doPickUp() []int {
	pickUp := []int{}
	if c.CurrentCupIndex+3 >= len(c.Nums) {
		c.shiftCups()
		c.CurrentCupIndex -= 3
	}
	pickUp = append(pickUp, c.Nums[c.CurrentCupIndex+1:c.CurrentCupIndex+4]...)
	c.Nums = append(c.Nums[:c.CurrentCupIndex+1], c.Nums[c.CurrentCupIndex+4:]...)
	return pickUp
}

func (c *Cups) doPutDown(pickup []int) {
	destinationIndex := c.getDestinationIndex()
	c.Nums = append(c.Nums[:destinationIndex+1], append(pickup, c.Nums[destinationIndex+1:]...)...)
	c.CurrentCupIndex = (c.CurrentCupIndex + 1) % len(c.Nums)
}

func (c *Cups) doMove() {
	c.doPutDown(c.doPickUp())
}

func (c *Cups) playGame(rounds int) {
	// seen := [][]int{}
	for i := 0; i < rounds; i++ {
		c.doMove()
		// numCopy := []int{}
		// copy(numCopy, c.Nums)
		// seen = append(seen, numCopy)
		// if c.seenNumsBefore(seen) {
		// 	fmt.Println("seen!!", seen)
		// }
	}
}

func (c *Cups) populateCupsForPart2() {
	highestCup := 0
	for _, num := range c.Nums {
		if num > highestCup {
			highestCup = num
		}
	}
	for i := highestCup + 1; i <= 1000000; i++ {
		c.Nums = append(c.Nums, i)
	}
}

func (c *Cups) productOfCupsToRightOf1() (int, error) {
	indexOf1 := c.indexOfNum(1)
	if indexOf1 == -1 {
		return -1, fmt.Errorf("could not find 1 in nums %v", c.Nums)
	}
	return c.Nums[(indexOf1+1)%len(c.Nums)] * c.Nums[(indexOf1+2)%len(c.Nums)], nil
}

func main() {
	input := helpers.ReadFile()
	cups := Cups{}
	for _, num := range input[0] {
		cups.Nums = append(cups.Nums, int(num-'0'))
	}
	cups.CurrentCupIndex = 0
	cups.playGame(100)
	str, err := cups.getOrderString()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", str)
	cups = Cups{}
	for _, num := range input[0] {
		cups.Nums = append(cups.Nums, int(num-'0'))
	}
	cups.CurrentCupIndex = 0
	cups.populateCupsForPart2()
	cups.playGame(1000000)
	sol, err := cups.productOfCupsToRightOf1()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", sol)
}
