package main

import (
	helpers "Advent-of-Code"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type ValidNumbers map[int]bool
type PossibleRows map[int]bool
type FieldParams struct {
	vn ValidNumbers
	pr PossibleRows
}
type TicketFields map[string]FieldParams
type AllTickets [][]int

func (tf TicketFields) populateField(re *regexp.Regexp, field []string) error {
	rangeLimits := re.FindAllString(field[1], -1)
	rangeLimitsNums := helpers.StringSliceToIntSlice(rangeLimits)
	if len(rangeLimitsNums) != 4 {
		return errors.New("range limits not as expected")
	}
	vn := ValidNumbers{}
	for i := rangeLimitsNums[0]; i <= rangeLimitsNums[3]; i++ {
		if i <= rangeLimitsNums[1] || i >= rangeLimitsNums[2] {
			vn[i] = true
		}
	}
	tf[field[0]] = FieldParams{
		vn: vn,
		pr: tf[field[0]].pr,
	}
	return nil
}

func main() {
	entries := helpers.ReadFile()
	re := regexp.MustCompile(`\d+`)
	tFields := TicketFields{}
	myTicket := []int{}
	allTickets := AllTickets{}
	errorRate := 0
	for _, entry := range entries {
		if entry == "" || entry == "your ticket:" || entry == "nearby tickets:" {
			continue
		}
		field := strings.Split(entry, ":")
		if len(field) > 1 {
			if err := tFields.populateField(re, field); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			nums := helpers.StringSliceToIntSlice(re.FindAllString(entry, -1))
			if len(myTicket) == 0 {
				myTicket = nums
				for field := range tFields {
					possibleRows := PossibleRows{}
					for i := 0; i < len(myTicket); i++ {
						possibleRows[i] = true
					}
					tFields[field] = FieldParams{
						pr: possibleRows,
						vn: tFields[field].vn,
					}
				}
			} else {
				validTicket := true
				for _, num := range nums {
					if isValid := tFields.numIsValid(num); !isValid {
						errorRate += num
						validTicket = false
					}
				}
				if validTicket {
					allTickets = append(allTickets, nums)
				}
			}
		}
	}

	fmt.Println("Part 1:", errorRate)

	ticketColToField := make(map[int]string)

	for len(ticketColToField) < len(myTicket) {
		for _, ticket := range allTickets {
			for j, num := range ticket {
				for field, param := range tFields {
					if !param.vn[num] {
						delete(tFields[field].pr, j)
					}
					if len(tFields[field].pr) == 1 {
						// There should only be one so the loop only runs once
						for key := range tFields[field].pr {
							ticketColToField[key] = field
							for f := range tFields {
								if f != field {
									delete(tFields[f].pr, key)
								}
							}
						}

					}
				}
			}
		}
	}

	runningTotal := 1
	for i, val := range myTicket {
		if strings.HasPrefix(ticketColToField[i], "departure") {
			runningTotal *= val
		}
	}

	fmt.Println("Part 2:", runningTotal)

}

func (tf TicketFields) numIsValid(num int) bool {
	for _, field := range tf {
		if field.vn[num] {
			return true
		}
	}
	return false
}
