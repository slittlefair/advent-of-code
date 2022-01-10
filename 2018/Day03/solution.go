package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
	"strconv"
)

// claim allows us to work out the coordinates in a claim
type claim struct {
	id          int
	topLeft     co
	bottomRight co
}

type co struct {
	x int
	y int
}

var coords = make(map[co][]int)

func main() {
	lines := utils.ReadFile()
	var matches int
	var ids []int
	// cycle through cuts (or claims) and pick out the relevant information
	for _, cut := range lines {
		re := regexp.MustCompile("[^# @ ,: x]+")
		s := re.FindAllString(cut, -1)
		var t []int
		// convert strings from regex match to int
		for _, v := range s {
			v, err := strconv.Atoi(v)
			utils.Check(err)
			t = append(t, v)
		}
		// make a claim with top left and bottom right coordinates
		c := claim{
			id:          t[0],
			topLeft:     co{t[1], t[2]},
			bottomRight: co{t[1] + t[3] - 1, t[2] + t[4] - 1},
		}
		ids = append(ids, c.id)
		// cycle through each claim and make coordinates for every inch in the claim
		for i := c.topLeft.x; i <= c.bottomRight.x; i++ {
			for j := c.topLeft.y; j <= c.bottomRight.y; j++ {
				// add the id of the claim to the coordinate in our map
				if val, ok := coords[co{i, j}]; ok {
					coords[co{i, j}] = append(val, c.id)
				} else {
					coords[co{i, j}] = []int{c.id}
				}
			}
		}
	}
	// lastly cycle through coordinates and count all instances where there are more than one coord
	for _, val := range coords {
		if len(val) > 1 {
			matches++
		}
	}
	fmt.Println("Part 1:", matches)

	for _, id := range ids {
		intact := true
		for _, val := range coords {
			if len(val) > 1 {
				for _, i := range val {
					if i == id {
						intact = false
					}
				}
			}
		}
		if intact {
			fmt.Println("Part 2:", id)
			return
		}
	}
}
