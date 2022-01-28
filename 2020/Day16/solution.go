package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/slice"
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

// var allValidTickets TicketCollection

var re = regexp.MustCompile(`\d+`)

// Populate a field with the valid numbers, so we can just look up a number on a ticket to see if it
// is valid for a field
func (tf TicketFields) populateField(field []string) error {
	rangeLimits := re.FindAllString(field[1], -1)
	rangeLimitsNums, err := slice.StringSliceToIntSlice(rangeLimits)
	if err != nil {
		return err
	}
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
func (tf TicketFields) populatePossibleValueIndices(maxIndex int) {
	for field := range tf {
		possibleValueIndices := PossibleValueIndices{}
		for i := 0; i < maxIndex; i++ {
			possibleValueIndices[i] = true
		}
		tf[field] = FieldParams{
			pvi: possibleValueIndices,
			vn:  tf[field].vn,
		}
	}
}

func (tf TicketFields) validateTicket(nums []int, allValidTickets TicketCollection) (int, TicketCollection) {
	errorRate := 0
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
	return errorRate, allValidTickets
}

func (tf TicketFields) part2(myTicket []int, allValidTickets TicketCollection) int {
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
					// indices for that field
					if !param.vn[num] {
						delete(tf[field].pvi, j)
					}
					// Once we've deleted a possible indices, check to see how many possibilities
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

func (tf TicketFields) runSolution(entries []string) ([]int, int, TicketCollection, error) {
	allValidTickets := TicketCollection{}
	myTicket := []int{}
	var errorRate, er int
	for _, entry := range entries {
		if entry == "" || entry == "your ticket:" || entry == "nearby tickets:" {
			continue
		}
		field := strings.Split(entry, ":")
		// If the split yields more than one string then it is a field, otherwise it's a ticket
		if len(field) > 1 {
			if err := tf.populateField(field); err != nil {
				return nil, 0, nil, err
			}
		} else {
			nums, err := slice.StringSliceToIntSlice(re.FindAllString(entry, -1))
			if err != nil {
				return nil, 0, nil, err
			}
			if len(myTicket) == 0 {
				myTicket = nums
				tf.populatePossibleValueIndices(len(myTicket))
			} else {
				er, allValidTickets = tf.validateTicket(nums, allValidTickets)
				errorRate += er
			}
		}
	}
	return myTicket, errorRate, allValidTickets, nil
}

func main() {
	entries := file.Read()
	tFields := TicketFields{}
	myTicket, errorRate, allValidTickets, err := tFields.runSolution(entries)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", errorRate)
	fmt.Println("Part 2:", tFields.part2(myTicket, allValidTickets))
}
