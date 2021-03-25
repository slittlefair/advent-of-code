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
type TicketCollection [][]int

var myTicket []int
var allValidTickets TicketCollection

var re = regexp.MustCompile(`\d+`)

var errorRate int

func (tf TicketFields) populateField(field []string) error {
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

func (tf TicketFields) numIsValid(num int) bool {
	for _, field := range tf {
		if field.vn[num] {
			return true
		}
	}
	return false
}

func (tf TicketFields) part1(ticket string) {
	nums := helpers.StringSliceToIntSlice(re.FindAllString(ticket, -1))
	if len(myTicket) == 0 {
		myTicket = nums
		for field := range tf {
			possibleRows := PossibleRows{}
			for i := 0; i < len(myTicket); i++ {
				possibleRows[i] = true
			}
			tf[field] = FieldParams{
				pr: possibleRows,
				vn: tf[field].vn,
			}
		}
	} else {
		validTicket := true
		for _, num := range nums {
			if isValid := tf.numIsValid(num); !isValid {
				errorRate += num
				validTicket = false
			}
		}
		if validTicket {
			allValidTickets = append(allValidTickets, nums)
		}
	}
}

func (tf TicketFields) part2() int {
	ticketColToField := make(map[int]string)
	for len(ticketColToField) < len(myTicket) {
		for _, ticket := range allValidTickets {
			for j, num := range ticket {
				for field, param := range tf {
					if !param.vn[num] {
						delete(tf[field].pr, j)
					}
					if len(tf[field].pr) == 1 {
						// There should only be one so the loop only runs once
						for key := range tf[field].pr {
							ticketColToField[key] = field
							for f := range tf {
								if f != field {
									delete(tf[f].pr, key)
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

	return runningTotal
}

func main() {
	entries := helpers.ReadFile()
	tFields := TicketFields{}
	for _, entry := range entries {
		if entry == "" || entry == "your ticket:" || entry == "nearby tickets:" {
			continue
		}
		field := strings.Split(entry, ":")
		if len(field) > 1 {
			if err := tFields.populateField(field); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			tFields.part1(entry)
		}
	}

	fmt.Println("Part 1:", errorRate)
	fmt.Println("Part 2:", tFields.part2())

}
