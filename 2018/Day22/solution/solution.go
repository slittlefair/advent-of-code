package main

import (
	"fmt"
)

// var inputDepth = 11394
// var inputX = 7
// var inputY = 701

// var mouth = helpers.Co{X: minX, Y: minY}
// var target = helpers.Co{X: inputX, Y: inputY}
// var depth = inputDepth

// var g = dijkstra.Graph{}

// const minX = 0

// var maxX = inputX + 150

// const minY = 0

// var maxY = inputY + 150

// type region struct {
// 	co              helpers.Co
// 	rType           int
// 	erosionLevel    int
// 	geologicalIndex int
// }

// type meType struct {
// 	co       helpers.Co
// 	torch    bool
// 	climbing bool
// }

// var me = meType{
// 	co:       helpers.Co{X: 0, Y: 0},
// 	torch:    true,
// 	climbing: false,
// }

// type node struct {
// 	key  string
// 	cost int
// }

// var typeToString = map[int]string{
// 	0: ".",
// 	1: "=",
// 	2: "|",
// }

// // CaveMapType is a map of coordinates to regions which we can use to form the map
// type CaveMapType map[helpers.Co]region

// // CaveMap is an instance of CaveMapType
// var CaveMap = CaveMapType{}

// func (cm CaveMapType) printCave() {
// 	for y := minY; y <= maxY; y++ {
// 		for x := minX; x <= maxX; x++ {
// 			if x == mouth.X && y == mouth.Y {
// 				fmt.Print("M")
// 			} else if x == target.X && y == target.Y {
// 				fmt.Print("T")
// 			} else {
// 				fmt.Print(typeToString[cm[helpers.Co{X: x, Y: y}].rType])
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func populateCaveMap() {
// 	for y := minY; y <= maxY; y++ {
// 		for x := minX; x <= maxX; x++ {
// 			CaveMap[helpers.Co{X: x, Y: y}] = region{co: helpers.Co{X: x, Y: y}}
// 		}
// 	}
// }

// func (r region) calculateGeoIndex() int {
// 	if r.co.Y == 0 {
// 		return r.co.X * 16807
// 	} else if r.co.X == 0 {
// 		return r.co.Y * 48271
// 	} else {
// 		adjXCoord := helpers.Co{X: r.co.X - 1, Y: r.co.Y}
// 		adjYCoord := helpers.Co{X: r.co.X, Y: r.co.Y - 1}
// 		return CaveMap[adjXCoord].erosionLevel * CaveMap[adjYCoord].erosionLevel
// 	}
// }

// func (r region) calculateErosionLevel() int {
// 	return (r.geologicalIndex + depth) % 20183
// }

// func (r region) calculateType() int {
// 	return r.erosionLevel % 3
// }

// func (cm CaveMapType) calculateRisk() (total int) {
// 	for co, reg := range cm {
// 		if co.X <= inputX && co.Y <= inputY {
// 			total += reg.rType
// 		}
// 	}
// 	return total
// }

func main() {
	// populateCaveMap()
	// // CaveMap.printCave()
	// for y := minY; y <= maxY; y++ {
	// 	for x := minX; x <= maxX; x++ {
	// 		co := helpers.Co{X: x, Y: y}
	// 		reg := CaveMap[co]
	// 		if (x == minX && y == minY) || (x == inputX && y == inputY) {
	// 			reg.geologicalIndex = 0
	// 		} else {
	// 			reg.geologicalIndex = reg.calculateGeoIndex()
	// 		}
	// 		reg.erosionLevel = reg.calculateErosionLevel()
	// 		reg.rType = reg.calculateType()
	// 		CaveMap[co] = reg
	// 	}
	// }
	// // CaveMap.printCave()
	// fmt.Println("Part 1:", CaveMap.calculateRisk())

	// for co := range CaveMap {
	// 	coType := CaveMap[co].rType
	// 	for _, equip := range dijkstra.TypeToEquipment[coType] {
	// 		coString := fmt.Sprintf("%v,%v,%v", strconv.Itoa(co.X), strconv.Itoa(co.Y), equip)
	// 		g[coString] = make(map[string]int)
	// 		adjacentCo := make(map[helpers.Co]int)
	// 		xOffset := []int{-1, 1, 0, 0}
	// 		yOffset := []int{0, 0, -1, 1}
	// 		for i := range xOffset {
	// 			coord := helpers.Co{X: co.X + xOffset[i], Y: co.Y + yOffset[i]}
	// 			if _, ok := CaveMap[coord]; ok {
	// 				adjacentCo[coord] = CaveMap[coord].rType
	// 			}
	// 		}
	// 		g.CalculateTimes(coString, coType, adjacentCo, equip)
	// 	}
	// }

	// mouthString := fmt.Sprintf("%v,%v,torch", strconv.Itoa(mouth.X), strconv.Itoa(mouth.Y))
	// targetString := fmt.Sprintf("%v,%v,torch", strconv.Itoa(target.X), strconv.Itoa(target.Y))

	// path, time, _ := g.Path(mouthString, targetString)
	// fmt.Println("Part 2:", time)

	// total := 0
	// for i, co := range path {
	// 	if i < len(path)-1 {
	// 		coSplit := strings.Split(co, ",")
	// 		nextCoSplit := strings.Split(path[i+1], ",")
	// 		if coSplit[2] == nextCoSplit[2] {
	// 			total++
	// 		} else {
	// 			total += 8
	// 		}
	// 		fmt.Printf("Moving from (%v,%v) to (%v,%v). Cost is %d after %d steps.  %v is now equipped.\n", coSplit[0], coSplit[1], nextCoSplit[0], nextCoSplit[1], total, i+1, nextCoSplit[2])
	// 	}
	// }
	// fmt.Println("Total cost is", total, "after", len(path), "steps.")

	fmt.Println("TODO")
}
