package main

import (
	"Advent-of-Code"
	"fmt"
)

var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "j", "k", "m", "n", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func letterIndex(s string) (i int) {
	for i, l := range letters {
		if s == l {
			return i
		}
	}
	return
}

func legalLetter(s string) (isLegal bool) {
	for _, l := range letters {
		if l == s {
			isLegal = true
		}
	}
	return isLegal
}

func increment(password string) string {
	if password == "zzzzzzzz" {
		return "aaaaaaaa"
	}
	if string(password[len(password)-1]) != "z" {
		password = password[:len(password)-1] + letters[letterIndex(string(password[len(password)-1]))+1]
		return password
	}
	password = password[:len(password)-1] + "a"
	for i := len(password) - 2; i >= 1; i-- {
		if string(password[i]) != "z" {
			password = password[:i] + letters[letterIndex(string(password[i]))+1] + password[i+1:]
			return password
		}
		password = password[:i] + "a" + password[i+1:]
	}
	return password
}

func main() {
	password := helpers.ReadFile()[0]
	part1Done := false
	for {
		consecutiveMatch := false
		forbiddenLetterCheck := true
		doubles := make(map[string]int)
		for i, str := range password {
			s := string(str)
			// Check forbidden letter
			if ok := legalLetter(s); forbiddenLetterCheck && !ok {
				forbiddenLetterCheck = false
			}
			// Check consecutive match
			if i < len(password)-2 {
				idx := letterIndex(s)
				if !consecutiveMatch && idx < len(letters)-2 && string(password[i+1]) == letters[idx+1] && string(password[i+2]) == letters[idx+2] {
					consecutiveMatch = true
				}
			}
		}
		// Check doubles
		for i := 0; i < len(password); i++ {
			if i <= len(password)-2 {
				if password[i] == password[i+1] {
					if freq, ok := doubles[string(password[i])+string(password[i+1])]; !ok {
						doubles[string(password[i])+string(password[i+1])] = 1
					} else {
						doubles[string(password[i])+string(password[i+1])] = freq + 1
					}
					if i < len(password)-2 && password[i] == password[i+1] && password[i] == password[i+2] {
						i++
					}
				}
			}
		}
		numDoubles := 0
		for _, freq := range doubles {
			numDoubles += freq
		}
		if numDoubles > 1 && consecutiveMatch && forbiddenLetterCheck {
			if !part1Done {
				fmt.Println("Part 1:", password)
				part1Done = true
			} else {
				fmt.Println("Part 2:", password)
				return
			}
		}
		password = increment(password)
	}
}
