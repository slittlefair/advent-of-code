package card

import (
	helpers "Advent-of-Code"
	"regexp"
	"strconv"
)

type Number struct {
	Val    int
	Called bool
}

type Card struct {
	Numbers map[helpers.Co]*Number
}

func (c *Card) ParseCard(lines []string, reNum *regexp.Regexp) error {
	for i := 1; i < len(lines); i++ {
		matches := reNum.FindAllString(lines[i], -1)
		for j, m := range matches {
			// Because we convert what's been obtained by regex we should never error, but handle
			// it regardless
			val, err := strconv.Atoi(m)
			if err != nil {
				return err
			}
			c.Numbers[helpers.Co{X: j, Y: i - 1}] = &Number{Val: val}
		}
	}
	return nil
}

func (c *Card) CardIsWinner() bool {
	for x := 0; x < 5; x++ {
		winner := true
		for y := 0; y < 5; y++ {
			if !c.Numbers[helpers.Co{X: x, Y: y}].Called {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}
	for y := 0; y < 5; y++ {
		winner := true
		for x := 0; x < 5; x++ {
			if !c.Numbers[helpers.Co{X: x, Y: y}].Called {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}
	return false
}

func (c Card) CalculateScore(num int) int {
	sum := 0
	for _, n := range c.Numbers {
		if !n.Called {
			sum += n.Val
		}
	}
	return sum * num
}
