package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"strings"
)

var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func letterIndex(s string) int {
	for i, l := range letters {
		if s == l {
			return i
		}
	}
	return -1
}

type Password []int

func incrementCharacter(n int) int {
	return (n + 1) % len(letters)
}

func (p Password) RemoveIllegalCharacters() {
	for i := 0; i < len(p); i++ {
		if p[i] == 8 || p[i] == 11 || p[i] == 14 {
			p[i] = incrementCharacter(p[i])
			for j := i + 1; j < len(p); j++ {
				p[j] = 0
			}
		}
	}
}

func (p Password) IncrementPassword() {
	p.RemoveIllegalCharacters()
	for i := len(p) - 1; i >= 0; i-- {
		p[i] = incrementCharacter(p[i])
		if p[i] != 0 {
			return
		}
	}
}

func (p Password) HasIncreasingStraightLetters() bool {
	for i := 0; i < len(p)-2; i++ {
		if p[i+1] == p[i]+1 && p[i+2] == p[i]+2 {
			return true
		}
	}
	return false
}

func (p Password) HasDifferentPairs() bool {
	doubles := map[int]bool{}
	for i := 0; i < len(p)-1; i++ {
		if p[i] == p[i+1] {
			doubles[p[i]] = true
			if i < len(p)-2 && p[i] == p[i+1] && p[i] == p[i+2] {
				i++
			}
		}
	}
	return len(doubles) >= 2
}

func (p Password) IsValid() bool {
	if !p.HasIncreasingStraightLetters() {
		return false
	}
	if !p.HasDifferentPairs() {
		return false
	}
	return true
}

func (p Password) ConvertToString() string {
	str := ""
	for _, n := range p {
		str += letters[n]
	}
	return str
}

func (p Password) GetNextValidPassword() {
	for {
		p.IncrementPassword()
		if p.IsValid() {
			return
		}
	}
}

func makePassword(passString string) *Password {
	password := make(Password, len(passString))
	for i, n := range strings.Split(passString, "") {
		password[i] = letterIndex(n)
	}
	password.RemoveIllegalCharacters()
	return &password
}

func main() {
	input := utils.ReadFile()[0]
	password := makePassword(input)
	password.GetNextValidPassword()
	fmt.Println("Part 1:", password.ConvertToString())
	password.GetNextValidPassword()
	fmt.Println("Part 2:", password.ConvertToString())
}
