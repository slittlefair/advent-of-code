package main

import (
	"Advent-of-Code/utils"
	"fmt"
	"math"
	"strconv"
)

type Pair struct {
	parent    *Pair
	leftPair  *Pair
	rightPair *Pair
	leftVal   *int
	rightVal  *int
}

func (p *Pair) leftPopulated() bool {
	return p.leftPair != nil || p.leftVal != nil
}

func parseLine(line string, i *int) (*Pair, error) {
	*i++
	currentPair := &Pair{}
	for {
		if *i > len(line)-1 {
			return nil, fmt.Errorf("expected line to end with closing bracket: %s", line)
		}
		char := string(line[*i])
		switch char {
		case "[":
			newPair, err := parseLine(line, i)
			if err != nil {
				return nil, err
			}
			if !currentPair.leftPopulated() {
				currentPair.leftPair = newPair
			} else {
				currentPair.rightPair = newPair
			}
			newPair.parent = currentPair
		case "]":
			return currentPair, nil
		case ",":
			// do nothing
		default:
			val, err := strconv.Atoi(char)
			if err != nil {
				return nil, fmt.Errorf("expected int, got %s", char)
			}
			if !currentPair.leftPopulated() {
				currentPair.leftVal = &val
			} else {
				currentPair.rightVal = &val
			}
		}
		*i++
	}
}

type ExplodingPair struct {
	pair        *Pair
	left, right *int
}

func (p *Pair) findExplodingPair(ep *ExplodingPair, level int) *ExplodingPair {
	if ep.right != nil {
		return ep
	}
	if level == 4 && ep.pair == nil {
		ep.pair = p
		return ep
	}
	if p.leftPair != nil {
		p.leftPair.findExplodingPair(ep, level+1)
	}
	if ep.pair == nil && p.leftVal != nil {
		ep.left = p.leftVal
	}
	if ep.pair != nil && p.leftVal != nil && ep.right == nil {
		ep.right = p.leftVal
	}
	if p.rightPair != nil {
		p.rightPair.findExplodingPair(ep, level+1)
	}
	if ep.pair == nil && p.rightVal != nil {
		ep.left = p.rightVal
	}
	if ep.pair != nil && p.rightVal != nil && ep.right == nil {
		ep.right = p.rightVal
	}
	return ep
}

func (p *Pair) explode() bool {
	ep := p.findExplodingPair(&ExplodingPair{}, 0)
	if ep.pair == nil {
		return false
	}
	if ep.left != nil {
		*ep.left += *ep.pair.leftVal
	}
	if ep.right != nil {
		*ep.right += *ep.pair.rightVal
	}
	newVal := 0
	if ep.pair == ep.pair.parent.leftPair {
		ep.pair.parent.leftPair = nil
		ep.pair.parent.leftVal = &newVal
	} else if ep.pair == ep.pair.parent.rightPair {
		ep.pair.parent.rightPair = nil
		ep.pair.parent.rightVal = &newVal
	}
	return true
}

type SplittingPair struct {
	pair *Pair
}

func (p *Pair) split() bool {
	sp := p.findSplittingPair(&SplittingPair{}).pair
	if sp == nil {
		return false
	}
	if sp.leftVal != nil && *sp.leftVal > 9 {
		lVal := int(math.Floor(float64(*sp.leftVal) / 2))
		rVal := int(math.Ceil(float64(*sp.leftVal) / 2))
		sp.leftPair = &Pair{
			leftVal:  &lVal,
			rightVal: &rVal,
			parent:   sp,
		}
		sp.leftVal = nil
		return true
	}
	if sp.rightVal != nil && *sp.rightVal > 9 {
		lVal := int(math.Floor(float64(*sp.rightVal) / 2))
		rVal := int(math.Ceil(float64(*sp.rightVal) / 2))
		sp.rightPair = &Pair{
			leftVal:  &lVal,
			rightVal: &rVal,
			parent:   sp,
		}
		sp.rightVal = nil
		return true
	}
	return false
}

func (p *Pair) findSplittingPair(sp *SplittingPair) *SplittingPair {
	if p.leftVal != nil && *p.leftVal > 9 {
		sp.pair = p
		return sp
	}
	if p.leftPair != nil {
		p.leftPair.findSplittingPair(sp)
	}
	if sp.pair != nil {
		return sp
	}
	if p.rightVal != nil && *p.rightVal > 9 {
		sp.pair = p
		return sp
	}
	if p.rightPair != nil {
		p.rightPair.findSplittingPair(sp)
	}
	return sp
}

func (p *Pair) addPair(pair *Pair) *Pair {
	parent := &Pair{}
	p.parent = parent
	pair.parent = parent
	parent.leftPair = p
	parent.rightPair = pair
	return parent
}

func (p *Pair) doSum(newPair *Pair) *Pair {
	p = p.addPair(newPair)
	for {
		for {
			didExplode := p.explode()
			if !didExplode {
				break
			}
		}
		didSplit := p.split()
		if !didSplit {
			return p
		}
	}
}

func (p *Pair) findMagnitude() int {
	mag := 0
	if p.leftVal != nil {
		mag += 3 * *p.leftVal
	} else {
		mag += 3 * p.leftPair.findMagnitude()
	}
	if p.rightVal != nil {
		mag += 2 * *p.rightVal
	} else {
		mag += 2 * p.rightPair.findMagnitude()
	}
	return mag
}

func part1(input []string) (int, error) {
	numbers := []*Pair{}
	for _, line := range input {
		idx := 0
		pair, err := parseLine(line, &idx)
		if err != nil {
			return -1, err
		}
		numbers = append(numbers, pair)
	}
	pair := numbers[0]
	for i := 1; i < len(numbers); i++ {
		pair = pair.doSum(numbers[i])
	}
	return pair.findMagnitude(), nil
}

func part2(input []string) (int, error) {
	greatestMag := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i != j {
				idx := 0
				pair1, err := parseLine(input[i], &idx)
				if err != nil {
					return -1, err
				}
				idx = 0
				pair2, err := parseLine(input[j], &idx)
				if err != nil {
					return -1, err
				}
				sum := pair1.doSum(pair2)
				mag := sum.findMagnitude()
				if mag > greatestMag {
					greatestMag = mag
				}
			}
		}
	}
	return greatestMag, nil
}

func findSolutions(input []string) (int, int, error) {
	part1, err := part1(input)
	if err != nil {
		return -1, -1, err
	}
	// Any error obtained by part2 is from parseLine, which would have already returned an error
	// from part1. So if part1 does not return an error, then part2 won't either, so it can be
	// safely ignored.
	part2, _ := part2(input)
	return part1, part2, nil
}

func main() {
	input := utils.ReadFile()
	part1, part2, err := findSolutions(input)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

// Debugging
// func buildPair(pair *Pair, b *strings.Builder) {
// 	b.WriteString("[")
// 	if pair.leftPair != nil {
// 		buildPair(pair.leftPair, b)
// 	} else if pair.leftVal != nil {
// 		b.WriteString(fmt.Sprintf("%v", *pair.leftVal))
// 	}
// 	b.WriteString(",")
// 	if pair.rightPair != nil {
// 		buildPair(pair.rightPair, b)
// 	} else if pair.rightVal != nil {
// 		b.WriteString(fmt.Sprintf("%v", *pair.rightVal))
// 	}
// 	b.WriteString("]")
// }

// func printPair(pair *Pair) *strings.Builder {
// 	b := &strings.Builder{}
// 	if pair != nil {
// 		buildPair(pair, b)
// 	}
// 	return b
// }
