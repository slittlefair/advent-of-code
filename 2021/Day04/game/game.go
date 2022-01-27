package game

import (
	"Advent-of-Code/2021/Day04/card"
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	CardsNotWon map[*card.Card]struct{}
	Cards       []*card.Card
	Nums        []int
}

func parseNums(str string) ([]int, error) {
	nums := []int{}
	strNums := strings.Split(str, ",")
	for _, n := range strNums {
		i, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		nums = append(nums, i)
	}
	return nums, nil
}

func ParseInput(input []string) (*Game, error) {
	nums, err := parseNums(input[0])
	if err != nil {
		return nil, err
	}
	g := &Game{}
	g.Nums = nums
	reNum := regexp.MustCompile(`\d+`)
	for i := 1; i < len(input); i += 6 {
		card := &card.Card{Numbers: make(map[utils.Co]*card.Number)}
		err := card.ParseCard(input[i:i+6], reNum)
		if err != nil {
			return nil, err
		}
		g.Cards = append(g.Cards, card)
	}
	g.CardsNotWon = make(map[*card.Card]struct{})
	for _, card := range g.Cards {
		g.CardsNotWon[card] = struct{}{}
	}
	return g, nil
}

func (g *Game) PlayGame() (int, int, error) {
	part1 := -1
	part2 := -1
	for _, n := range g.Nums {
		for _, c := range g.Cards {
			for _, num := range c.Numbers {
				if num.Val == n {
					num.Called = true
				}
			}
			if c.CardIsWinner() {
				if part1 == -1 {
					part1 = c.CalculateScore(n)
				}
				delete(g.CardsNotWon, c)
				if len(g.CardsNotWon) == 0 {
					part2 = c.CalculateScore(n)
					return part1, part2, nil
				}
			}
		}
	}
	return part1, part2, fmt.Errorf("could not find last winning card after all numbers called, %d cards remaining", len(g.CardsNotWon))
}
