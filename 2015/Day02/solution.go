package main

import (
	"Advent-of-Code"
	"fmt"
	"regexp"
)

func main() {
	presents := helpers.ReadFile()
	re := regexp.MustCompile("\\d+")
	totalPaperFootage := 0
	totalRibbonFootage := 0
	for _, present := range presents {
		nums := helpers.StringSliceToIntSlice(re.FindAllString(present, -1))
		h, l, w := nums[0], nums[1], nums[2]

		// Paper
		face1, face2, face3 := h*l, h*w, l*w
		var m int
		for i, n := range []int{face1, face2, face3} {
			if i == 0 || n < m {
				m = n
			}
		}
		totalPaperFootage += 2*face1 + 2*face2 + 2*face3 + m

		// Ribbon
		perim1, perim2, perim3 := h+h+l+l, h+h+w+w, l+l+w+w
		for i, n := range []int{w, l, h} {
			if i == 0 || n > m {
				m = n
			}
		}
		var ribbon int
		switch m {
		case w:
			ribbon = perim1
		case l:
			ribbon = perim2
		case h:
			ribbon = perim3
		}
		totalRibbonFootage += ribbon + w*h*l
	}
	fmt.Println("Part 1:", totalPaperFootage)
	fmt.Println("Part 2:", totalRibbonFootage)
}
