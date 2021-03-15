package main

import (
	"fmt"
	"math"
)

var grid [300][300]int

func calculatePower(x, y, serialNumber int) (power int) {
	rackID := x + 10
	power = rackID * y
	power += serialNumber
	power *= rackID
	if power < 100 {
		power = 0
	} else {
		power = power % int(math.Pow(10, float64(3)))
		power = power / int(math.Pow(10, float64(2)))
	}
	power -= 5
	return
}

func main() {
	var serialNumber = 7857
	for i, v := range grid {
		for j := range v {
			grid[i][j] = calculatePower(j+1, i+1, serialNumber)
		}
	}
	var maxPower int
	type maxPowerCoord struct {
		X int
		Y int
	}
	var maxCoord maxPowerCoord
	var maxSquareSize int
	for squareSize := 1; squareSize <= 300; squareSize++ {
		for i, yval := range grid {
			for j := range yval {
				coordPower := 0
				if i+squareSize <= 300 && j+squareSize <= 300 {
					for y := i; y < i+squareSize; y++ {
						for x := j; x < j+squareSize; x++ {
							coordPower += grid[y][x]
						}
					}
				}
				if coordPower > maxPower {
					maxPower = coordPower
					maxCoord = maxPowerCoord{j + 1, i + 1}
					maxSquareSize = squareSize
				}
			}
		}
	}
	fmt.Printf("%v,%v,%v", maxCoord.X, maxCoord.Y, maxSquareSize)
}
