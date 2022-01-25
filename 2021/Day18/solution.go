package main

import (
	"Advent-of-Code/utils"
	"fmt"
	"strconv"
	"strings"
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

func printPair(pair *Pair, b *strings.Builder) error {
	b.WriteString("[")
	if pair.leftPair != nil {
		printPair(pair.leftPair, b)
	} else if pair.leftVal != nil {
		b.WriteString(fmt.Sprintf("%v", *pair.leftVal))
	} else {
		return fmt.Errorf("got no leftPair or leftVal: %+v", pair)
	}
	b.WriteString(",")
	if pair.rightPair != nil {
		printPair(pair.rightPair, b)
	} else if pair.rightVal != nil {
		b.WriteString(fmt.Sprintf("%v", *pair.rightVal))
	} else {
		return fmt.Errorf("got no rightPair or rightVal: %+v", pair)
	}
	b.WriteString("]")
	return nil
}

func debug(pair *Pair) (*strings.Builder, error) {
	b := &strings.Builder{}
	err := printPair(pair, b)
	return b, err
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
		b, err := debug(pair)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(b)
		err = pair.explode()
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err = debug(pair)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(b)
	}
}
