package main

import (
	helpers "Advent-of-Code"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Valid numbers for a field which the numbers on a ticket could correspond to
type ValidNumbers map[int]bool

// Possible values of a ticket
type PossibleValueIndices map[int]bool
type FieldParams struct {
	vn  ValidNumbers
	pvi PossibleValueIndices
}
type TicketFields map[string]FieldParams
type TicketCollection [][]int

var myTicket []int
var allValidTickets TicketCollection

var re = regexp.MustCompile(`\d+`)

var errorRate int

// Populate a field with the valid numbers, so we can just look up a number on a ticket to see if it
// is valid for a field
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
	// TODO figure out how to use pointers to make this better
	tf[field[0]] = FieldParams{
		vn:  vn,
		pvi: tf[field[0]].pvi,
	}
	return nil
}

// Check a number on a ticket against all valid numbers for each field and return whether it appears
// for at least one field
func (tf TicketFields) numIsValid(num int) bool {
	for _, field := range tf {
		if field.vn[num] {
			return true
		}
	}
	return false
}

// For each field, set possible value indices, which initially is every index in a ticket
func (tf TicketFields) populatePossibleValueIndices() {
	for field := range tf {
		possibleValueIndices := PossibleValueIndices{}
		for i := 0; i < len(myTicket); i++ {
			possibleValueIndices[i] = true
		}
		tf[field] = FieldParams{
			pvi: possibleValueIndices,
			vn:  tf[field].vn,
		}
	}
}

func (tf TicketFields) part1(ticket string, nums []int) {
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

func (tf TicketFields) part2() int {
	ticketColToField := make(map[int]string)
	// Keep looping whilst we still have fields to find value for
	for len(ticketColToField) < len(myTicket) {
		// Loop over all valid tickets
		for _, ticket := range allValidTickets {
			// For each ticket, loop over the numbers in the ticket
			for j, num := range ticket {
				// For each number, loop over every field and parameter
				for field, param := range tf {
					// If the number from the ticket is not valid for that field, then the index of
					// the number cannot relate to that field, so delete it from the possible
					// indeces for that field
					if !param.vn[num] {
						delete(tf[field].pvi, j)
					}
					// Once we've deleted a possible indeces, check to see how many possibilities
					// are left. If there's only one, by process of elimination we have found the
					// index for which the field relates. In which case we add the index and field
					// pair to ticketColToField and delete the index from all other fields' possible
					// indices
					if len(tf[field].pvi) == 1 {
						// There is only one so the loop only runs once
						for key := range tf[field].pvi {
							ticketColToField[key] = field
							for f := range tf {
								if f != field {
									delete(tf[f].pvi, key)
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
		// If the split yields more than one string then it is a field, otherwise it's a ticket
		if len(field) > 1 {
			if err := tFields.populateField(field); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			nums := helpers.StringSliceToIntSlice(re.FindAllString(entry, -1))
			if len(myTicket) == 0 {
				myTicket = nums
				tFields.populatePossibleValueIndices()
			} else {
				tFields.part1(entry, nums)
			}
		}
	}

	fmt.Println("Part 1:", errorRate)
	fmt.Println("Part 2:", tFields.part2())

}
