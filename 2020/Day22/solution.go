package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/slice"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Deck []int

type Game struct {
	player1 Deck
	player2 Deck
}

const (
	Player1 = 1
	Player2 = 2
)

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

func (g *Game) player1Wins() {
	g.player1 = append(g.player1[1:], g.player1[0], g.player2[0])
	g.player2 = g.player2[1:]
}

func (g *Game) player2Wins() {
	g.player2 = append(g.player2[1:], g.player2[0], g.player1[0])
	g.player1 = g.player1[1:]
}

func (g *Game) playNormalRound() int {
	if g.player1[0] > g.player2[0] {
		g.player1Wins()
		return Player1
	}
	g.player2Wins()
	return Player2
}

func (g *Game) playNormalGame() Deck {
	for {
		g.playNormalRound()
		if len(g.player1) == 0 {
			return g.player2
		}
		if len(g.player2) == 0 {
			return g.player1
		}
	}
}

func (g Game) deckSeen(seen []Game) bool {
	for _, game := range seen {
		if slice.IntSlicesAreEqual(game.player1, g.player1) && slice.IntSlicesAreEqual(game.player2, g.player2) {
			return true
		}
	}
	return false
}

func (g *Game) playRecursiveRound(seen []Game) []Game {
	seen = append(seen, Game{
		player1: g.player1,
		player2: g.player2,
	})
	if g.player1[0] <= len(g.player1[1:]) && g.player2[0] <= len(g.player2[1:]) {
		g2 := Game{
			player1: append(Deck{}, g.player1[1:g.player1[0]+1]...),
			player2: append(Deck{}, g.player2[1:g.player2[0]+1]...),
		}
		winner, _ := g2.playRecursiveGame()
		if winner == Player1 {
			g.player1Wins()
		} else {
			g.player2Wins()
		}
		return seen
	}
	g.playNormalRound()
	return seen
}

func (g *Game) playRecursiveGame() (int, Deck) {
	seen := []Game{}
	roundNum := 1
	for {
		if g.deckSeen(seen) {
			return Player1, g.player1
		}
		if len(g.player1) == 0 {
			return Player2, g.player2
		}
		if len(g.player2) == 0 {
			return Player1, g.player1
		}
		seen = g.playRecursiveRound(seen)
		roundNum++
	}
}

func calculateWinningScore(deck Deck) (int, error) {
	score := 0
	if len(deck) == 0 {
		return score, errors.New("error")
	}
	for i, s := range deck {
		score += (len(deck) - i) * s
	}
	return score, nil
}

func main() {
	input := file.Read()
	game := Game{}
	err := game.parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	game2 := game
	winner := game.playNormalGame()
	score, err := calculateWinningScore(winner)
	if err != nil {
		fmt.Println("could not get score for", err)
		return
	}
	fmt.Println("Part 1:", score)
	_, winner = game2.playRecursiveGame()
	score, err = calculateWinningScore(winner)
	if err != nil {
		fmt.Println("could not get score for", err)
		return
	}
	fmt.Println("Part 2:", score)
}
