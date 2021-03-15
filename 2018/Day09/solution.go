package main

import "fmt"

const players = 455
const lastMarble = 7122300

var playerScores [players]int

var marbles = []int{0}

func main() {
	currentIndex := 0
	currentPlayer := 0
	for i := 1; i <= lastMarble; i++ {
		if i%23 == 0 {
			playerScores[currentPlayer] += i
			toRemove := (currentIndex - 7) % len(marbles)
			if toRemove < 0 {
				toRemove += len(marbles)
			}
			playerScores[currentPlayer] += marbles[toRemove]
			marbles = append(marbles[:toRemove], marbles[toRemove+1:]...)
			currentIndex = (toRemove) % len(marbles)
		} else {
			place := (currentIndex + 2) % len(marbles)
			marbles = append(marbles, 0)
			copy(marbles[place+1:], marbles[place:])
			marbles[place] = i
			currentIndex = place
		}
		if currentPlayer++; currentPlayer == players {
			currentPlayer = 0
		}
	}
	maxScore := 0
	for _, v := range playerScores {
		if v > maxScore {
			maxScore = v
		}
	}
	fmt.Println(maxScore)
}
