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

func (i Input) iterateMessages(key string, message string, index int) (bool, int) {
	// If the index is the length of the message then we've traversed each character in it
	if len(message) == index {
		return true, 0
	}

	rule := i.Rules[key]

	// If the rule val is a character then check whether it matches the character at the current
	// index of the current message
	if rule.val != "" {
		return string(message[index]) == rule.val, index + 1
	}

	for _, sr := range rule.subRules {
		subRulesMatched := true
		offset := index

		for _, k := range sr {
			match, nextIndex := i.iterateMessages(k, message, offset)
			if !match {
				subRulesMatched = false
				break
			}
			offset = nextIndex
		}

		if subRulesMatched {
			return true, offset
		}
	}

	return false, index
}

func (i Input) evaluateMessages() int {
	count := 0
	for _, message := range i.Messages {
		valid, offset := i.iterateMessages("0", message, 0)
		if valid && offset == len(message) {
			count++
		}
	}
	return count
}

func main() {
	rawInput := helpers.ReadFile()
	i := &Input{
		Rules: map[string]Rule{},
	}
	i.parseInput(rawInput)
	fmt.Println("Part 1:", i.evaluateMessages())
}
