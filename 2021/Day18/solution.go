package main

import (
	"Advent-of-Code/utils"
	"fmt"
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

func printPair(pair *Pair) {
	fmt.Printf("[")
	if pair.leftPair != nil {
		printPair(pair.leftPair)
	} else if pair.leftVal != nil {
		fmt.Printf("%v", *pair.leftVal)
	} else {
		fmt.Print("ERRORRRRRR")
	}
	fmt.Printf(",")
	if pair.rightPair != nil {
		printPair(pair.rightPair)
	} else {
		fmt.Printf("%v", *pair.rightVal)
	}
	fmt.Printf("]")
}

func (p *Pair) explode() error {
	ep := p.findExplodingPair(&ExplodingPair{}, 0)
	if ep == nil {
		return nil
	}
	if ep.pair.leftVal == nil || ep.pair.rightVal == nil {
		return fmt.Errorf("something went wrong, expected valid ep.pair, got %+v", ep.pair)
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
		return fmt.Errorf("something went wrong, expected ep.pair to match one of parent pairs, got ep.pair %+v and parent %+v", ep.pair, ep.pair.parent)
	}
	return nil
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
	if ep.pair == nil {
		ep.left = p.leftVal
	}
	if ep.pair != nil && p.leftVal != nil && ep.right == nil {
		ep.right = p.leftVal
	}
	if p.rightPair != nil {
		p.rightPair.findExplodingPair(ep, level+1)
	}
	if ep.pair == nil {
		ep.left = p.rightVal
	}
	if ep.pair != nil && p.rightVal != nil && ep.right == nil {
		ep.right = p.rightVal
	}
	return ep
}

func debug(pair *Pair) {
	printPair(pair)
	fmt.Println()
}

type ExplodingPair struct {
	pair        *Pair
	left, right *int
}

func main() {
	input := utils.ReadFile()
	for _, line := range input {
		idx := 0
		pair, err := parseLine(line, &idx)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println()
		debug(pair)
		err = pair.explode()
		if err != nil {
			fmt.Println(err)
			return
		}
		debug(pair)
	}
}
