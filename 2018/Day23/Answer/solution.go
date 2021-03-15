package day23

import (
	"Advent-of-Code"
	"fmt"
	"math"
	"regexp"
)

type coordinate struct {
	X int
	Y int
	Z int
}

type nano struct {
	r  int
	co coordinate
}

var allNanos []nano

func makeNano(line string) nano {
	re := regexp.MustCompile("-?\\d+")
	nums := helpers.StringSliceToIntSlice(re.FindAllString(line, -1))
	return nano{
		r:  nums[3],
		co: coordinate{X: nums[0], Y: nums[1], Z: nums[2]},
	}
}

func absoluteDist(startCo coordinate, endCo coordinate) int {
	xDist := math.Abs(float64(startCo.X - endCo.X))
	yDist := math.Abs(float64(startCo.Y - endCo.Y))
	zDist := math.Abs(float64(startCo.Z - endCo.Z))
	return int(xDist + yDist + zDist)
}

func main() {
	lines := helpers.ReadFile()
	for _, l := range lines {
		allNanos = append(allNanos, makeNano(l))
	}

	maxRange := 0
	maxRangeCo := coordinate{}
	for _, nano := range allNanos {
		if nano.r > maxRange {
			maxRange = nano.r
			maxRangeCo = coordinate{X: nano.co.X, Y: nano.co.Y, Z: nano.co.Z}
		}
	}

	nanosInProximity := 0
	for _, nano := range allNanos {
		if absoluteDist(nano.co, maxRangeCo) <= maxRange {
			nanosInProximity++
		}
	}

	fmt.Println("Part 1:", nanosInProximity)

	coordsInProximity := make(map[coordinate]int)
	for _, nano := range allNanos {
		for x := nano.co.X - nano.r; x <= nano.co.X+nano.r; x++ {
			fmt.Println(x)
			for y := nano.co.Y - (nano.r - int(math.Abs(float64(nano.co.X-x)))); y <= nano.co.Y+(nano.r-int(math.Abs(float64(nano.co.X-x)))); y++ {
				// fmt.Println(y)
				for z := nano.co.Z - (nano.r - int(math.Abs(float64(nano.co.X-x)))) - (nano.r - int(math.Abs(float64(nano.co.Y-y)))); z <= nano.co.Z+(nano.r-int(math.Abs(float64(nano.co.X-x))))+(nano.r-int(math.Abs(float64(nano.co.Y-y)))); z++ {
					// fmt.Println(z)
					// fmt.Println(nano.co, nano.r, x, y, z)
					if absoluteDist(coordinate{x, y, z}, nano.co) <= nano.r {
						// fmt.Println(coordinate{x, y, z})
						if freq, ok := coordsInProximity[coordinate{x, y, z}]; !ok {
							coordsInProximity[coordinate{x, y, z}] = 1
						} else {
							coordsInProximity[coordinate{x, y, z}] = freq + 1
						}
					}
				}
			}
		}
	}

	highestFreq := 0
	tiebreakers := make(map[coordinate]int)
	for _, freq := range coordsInProximity {
		if freq > highestFreq {
			highestFreq = freq
		}
	}
	fmt.Println("hightestfreq", highestFreq)

	for co, freq := range coordsInProximity {
		if freq == highestFreq {
			tiebreakers[co] = absoluteDist(co, coordinate{0, 0, 0})
		}
	}
	fmt.Println("tiebreakers", tiebreakers)

	nearestCoordinate := coordinate{}
	smallestDist := 10000000000
	for co, dist := range tiebreakers {
		if dist < smallestDist {
			smallestDist = dist
			nearestCoordinate = co
		}
	}

	fmt.Println("Part 2:", smallestDist, nearestCoordinate)
}
