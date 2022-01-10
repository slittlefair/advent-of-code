package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
	"sort"
)

var reString = regexp.MustCompile("[A-Z]")

type step struct {
	start  string
	finish string
}

var solution string
var allSteps = make(map[step]bool)

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
			doing := possibleCandidates[0]
			for k := range allSteps {
				if k.start == doing {
					allSteps[k] = true
				}
			}
			solution += doing
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
					fmt.Println(solution)
					return
				}
			}
		}
	}
}
