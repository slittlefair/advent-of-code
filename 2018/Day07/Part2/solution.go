package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var reString = regexp.MustCompile("[A-Z]")

type step struct {
	start  string
	finish string
}

var solution string
var allSteps = make(map[step]bool)
var inProgress = make(map[string]int)
var totalTime = 0

func main() {
	instructions := utils.ReadFile()
	for _, line := range instructions {
		s := reString.FindAllString(line, -1)
		allSteps[step{s[1], s[2]}] = false
	}
	for {
		var possibleCandidates []string
		for key, val := range allSteps {
			if !val {
				candidate := key.start
				for k, v := range allSteps {
					if k.finish == candidate {
						if !v {
							candidate = ""
							break
						}
					}
				}
				if candidate != "" {
					possibleCandidates = append(possibleCandidates, candidate)
				}
			}
		}
		if len(possibleCandidates) > 0 {
			sort.Strings(possibleCandidates)
			for _, val := range possibleCandidates {
				if _, ok := inProgress[val]; len(inProgress) < 5 && !strings.Contains(solution, val) && !ok {
					inProgress[val] = int([]rune(val)[0]) - 4
				}
			}
			for key, val := range inProgress {
				inProgress[key]--
				if val == 1 {
					delete(inProgress, key)
					solution += key
					for k := range allSteps {
						if k.start == key {
							allSteps[k] = true
						}
					}
				}
			}
			totalTime++
		}
		finished := true
		for _, val := range allSteps {
			if !val {
				finished = false
			}
		}
		if finished {
			for key := range allSteps {
				var lastChar = key.finish
				found := true
				for k := range allSteps {
					if lastChar == k.start {
						found = false
					}
				}
				if found {
					solution += key.finish
					totalTime += int([]rune(key.finish)[0]) - 4
					fmt.Println(totalTime)
					return
				}
			}
		}
	}
}
