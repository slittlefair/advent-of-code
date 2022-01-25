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
			return currentPair, nil
		}
		if char := string(line[*i]); char == "[" {
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
		} else if char == "]" {
			return currentPair, nil
		} else if char == "," {
			// skip
		} else {
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

func (p *Pair) explode() (bool, error) {
	ep := p.findExplodingPair(&ExplodingPair{}, 0)
	if ep.pair == nil {
		return false, nil
	}
	if ep.pair.leftVal == nil || ep.pair.rightVal == nil {
		return false, fmt.Errorf("something went wrong, expected valid ep.pair, got %+v", ep.pair)
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
	} else {
		return false, fmt.Errorf("something went wrong, expected ep.pair to match one of parent pairs, got ep.pair %+v and parent %+v", ep.pair, ep.pair.parent)
	}
	return true, nil
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

type ExplodingPair struct {
	pair        *Pair
	left, right *int
}

func (sp *SplittingPair) foundSplittingPair() bool {
	return sp.pair != nil
}

type SplittingPair struct {
	pair *Pair
}

func (p *Pair) findSplittingPair(sp *SplittingPair) *SplittingPair {
	if sp.foundSplittingPair() {
		return sp
	}
	if p.leftVal != nil && *p.leftVal > 9 {
		sp.pair = p
		return sp
	}
	if p.leftPair != nil {
		p.leftPair.findSplittingPair(sp)
	}
	if sp.foundSplittingPair() {
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

func (p *Pair) split() (bool, error) {
	sp := p.findSplittingPair(&SplittingPair{}).pair
	if sp == nil {
		return false, nil
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
		return true, nil
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
		return true, nil
	}
	return false, fmt.Errorf("not got correct splittingPair, leftVal: %v, rightVal: %v", sp.leftVal, sp.rightVal)
}

func (p *Pair) addPair(pair *Pair) *Pair {
	parent := &Pair{}
	p.parent = parent
	pair.parent = parent
	parent.leftPair = p
	parent.rightPair = pair
	return parent
}

func (p *Pair) doSum(newPair *Pair) (*Pair, error) {
	p = p.addPair(newPair)
	for {
		for {
			didExplode, err := p.explode()
			if err != nil {
				return nil, err
			}
			if !didExplode {
				break
			}
		}
		didSplit, err := p.split()
		if err != nil {
			return nil, err
		}
		if !didSplit {
			return p, nil
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
	var err error
	for i := 1; i < len(numbers); i++ {
		pair, err = pair.doSum(numbers[i])
		if err != nil {
			return -1, err
		}
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
				sum, err := pair1.doSum(pair2)
				if err != nil {
					return -1, err
				}
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
	part2, err := part2(input)
	if err != nil {
		return -1, -1, err
	}
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
