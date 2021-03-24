package main

import (
	helpers "Advent-of-Code"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type ValidNumbers map[int]bool
type TicketFields map[string]ValidNumbers
type AllTickets [][]int

func main() {
	entries := helpers.ReadFile()
	re := regexp.MustCompile(`\d+`)
	tFields := TicketFields{}
	myTicket := []int{}
	allTickets := AllTickets{}
	for _, entry := range entries {
		if entry == "" || entry == "your ticket:" || entry == "nearby tickets:" {
			continue
		}
		field := strings.Split(entry, ":")
		if len(field) > 1 {
			rangeLimits := re.FindAllString(field[1], -1)
			rangeLimitsNums := helpers.StringSliceToIntSlice(rangeLimits)
			if len(rangeLimitsNums) != 4 {
				fmt.Println(errors.New("range limits not as expected"), rangeLimitsNums)
			}
			vn := ValidNumbers{}
			for i := rangeLimitsNums[0]; i <= rangeLimitsNums[3]; i++ {
				if i <= rangeLimitsNums[1] || i >= rangeLimitsNums[2] {
					vn[i] = true
				}
			}
			tFields[field[0]] = vn
		} else {
			nums := helpers.StringSliceToIntSlice(re.FindAllString(entry, -1))
			if len(myTicket) == 0 {
				myTicket = nums
			} else {
				allTickets = append(allTickets, nums)
			}
		}
	}

	errorRate := 0

	for _, ticket := range allTickets {
		for _, num := range ticket {
			if isValid := tFields.numIsValid(num); !isValid {
				errorRate += num
			}
		}
	}

	fmt.Println("Part 1:", errorRate)
}

func (tf TicketFields) numIsValid(num int) bool {
	for _, field := range tf {
		if field[num] {
			return true
		}
	}
	return false
}
