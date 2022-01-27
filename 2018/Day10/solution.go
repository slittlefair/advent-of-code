package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type position struct {
	X int
	Y int
}

type velocity struct {
	X int
	Y int
}

type star struct {
	position position
	velocity velocity
}

var allStars []star

var minX, maxX, minY, maxY int

var printedGrid []int

func calculateDistances(i int) {
	var tmpMinX = 10000
	var tmpMaxX = 0
	var tmpMinY = 10000
	var tmpMaxY = 0
	for _, star := range allStars {
		if star.position.X < tmpMinX {
			tmpMinX = star.position.X
		}
		if star.position.X > tmpMaxX {
			tmpMaxX = star.position.X
		}
		if star.position.Y < tmpMinY {
			tmpMinY = star.position.Y
		}
		if star.position.Y > tmpMaxY {
			tmpMaxY = star.position.Y
		}
	}
	distX := math.Abs(float64(tmpMaxX - tmpMinX))
	distY := math.Abs(float64(tmpMaxY - tmpMinY))

	if distX < 75 && distY < 75 {
		printedGrid = append(printedGrid, i)
		generateGrid(i, tmpMinX, tmpMaxX, tmpMinY, tmpMaxY)
	}
}

func generateGrid(i, mnX, mxX, mnY, mxY int) {
	fmt.Printf("After %v seconds:\n\n", i)
	for i := mnY; i <= mxY; i++ {
		for j := mnX; j <= mxX; j++ {
			printed := false
			for _, s := range allStars {
				if s.position.X == j && s.position.Y == i {
					if !printed {
						fmt.Printf("#")
						printed = true
					}
				}
			}
			if !printed {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}

func updatePositions() {
	for i, s := range allStars {
		star := s
		star.position.X += star.velocity.X
		star.position.Y += star.velocity.Y
		allStars[i] = star
	}
}

func main() {
	lines := utils.ReadFile()
	re := regexp.MustCompile(`-?\d+`)
	for _, v := range lines {
		conv := re.FindAllString(v, -1)
		var points []int
		for i := range conv {
			c, err := strconv.Atoi(conv[i])
			utils.Check(err)
			points = append(points, c)
		}
		if points[0] < minX {
			minX = points[0]
		}
		if points[0] > maxX {
			maxX = points[0]
		}
		if points[1] < minY {
			minY = points[1]
		}
		if points[1] > maxY {
			maxY = points[1]
		}
		s := star{
			position: position{X: points[0], Y: points[1]},
			velocity: velocity{X: points[2], Y: points[3]},
		}
		allStars = append(allStars, s)
	}
	var limit int
	if max := maxX > maxY; max {
		limit = maxX
	} else {
		limit = maxY
	}
	for i := 0; i <= limit; i++ {
		calculateDistances(i)
		updatePositions()
	}
}
