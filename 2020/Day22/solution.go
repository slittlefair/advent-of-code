package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strconv"
	"strings"
)

type Deck []int

type Game struct {
	player1 Deck
	player2 Deck
}

func (g *Game) parseInput(input []string) error {
	var deck Deck
	for _, line := range input {
		if line == "" {
			g.player1 = deck
			deck = Deck{}
		} else if strings.HasPrefix(line, "Player") {
			continue
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				return err
			}
			deck = append(deck, num)
		}
	}
	g.player2 = deck
	return nil
}

func (g *Game) playNormalRound() {
	if g.player1[0] > g.player2[0] {
		g.player1 = append(g.player1[1:], g.player1[0], g.player2[0])
		g.player2 = g.player2[1:]
	} else {
		g.player2 = append(g.player2[1:], g.player2[0], g.player1[0])
		g.player1 = g.player1[1:]
	}
}

func (g *Game) playNormalGame() string {
	for {
		g.playNormalRound()
		if len(g.player1) == 0 {
			return "player2"
		}
		if len(g.player2) == 0 {
			return "player1"
		}
	}
}

func (g Game) calculateWinningScore(winner string) (int, error) {
	score := 0
	var winningDeck []int
	if winner == "player1" {
		winningDeck = g.player1
	}
	if winner == "player2" {
		winningDeck = g.player2
	}
	if len(winningDeck) == 0 {
		return score, fmt.Errorf(winner)
	}
	for i, s := range winningDeck {
		score += (len(winningDeck) - i) * s
	}
	return score, nil
}

func main() {
	input := helpers.ReadFile()
	game := Game{}
	err := game.parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	game2 := game
	winner := game.playNormalGame()
	score, err := game.calculateWinningScore(winner)
	if err != nil {
		fmt.Println("could not get score for", err)
		return
	}
	fmt.Println("Part 1:", score)
	fmt.Println(game)
	fmt.Println(game2)
}
