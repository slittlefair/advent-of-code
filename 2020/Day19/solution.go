package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strings"
)

type Rule struct {
	val      string
	subRules [][]string
}

type Input struct {
	Rules    map[string]Rule
	Messages []string
}

func (i *Input) parseInput(rawInput []string) {
	for _, line := range rawInput {
		if line == "" {
			continue
		}
		split := strings.Split(line, ": ")
		if len(split) == 1 {
			i.Messages = append(i.Messages, line)
		} else {
			key := split[0]
			if strings.HasPrefix(split[1], "\"") {
				i.Rules[key] = Rule{
					val: strings.Trim(split[1], "\""),
				}
				continue
			}

			newRule := Rule{
				subRules: [][]string{},
			}
			for _, str := range strings.Split(split[1], " | ") {
				newRule.subRules = append(newRule.subRules, strings.Split(str, " "))
			}
			i.Rules[key] = newRule
		}
	}
}

func main() {
	rawInput := helpers.ReadFile()
	i := &Input{
		Rules: map[string]Rule{},
	}
	i.parseInput(rawInput)
	fmt.Println(*i)
}
