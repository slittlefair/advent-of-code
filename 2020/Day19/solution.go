package main

import (
	utils "Advent-of-Code/utils"
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

func (i Input) iterateMessages(key string, remainingRules []string, message string, index int, seen map[string]bool) map[string]bool {
	rule := i.Rules[key]
	if rule.val != "" {
		if index >= len(message) {
			return seen
		}

		if rule.val == string(message[index]) {
			if len(remainingRules) > 0 {
				seen = i.iterateMessages(remainingRules[0], remainingRules[1:], message, index+1, seen)
			} else {
				if index == len(message)-1 {
					seen[message] = true
				}
			}
		}

		return seen
	}

	rule1 := rule.subRules[0]
	seen = i.iterateMessages(rule1[0], append(rule1[1:], remainingRules...), message, index, seen)

	if len(rule.subRules) > 1 {
		rule2 := rule.subRules[1]
		seen = i.iterateMessages(rule2[0], append(rule2[1:], remainingRules...), message, index, seen)
	}

	return seen
}

func (i Input) changeRulesForPart2() {
	for key := range i.Rules {
		if key == "8" {
			i.Rules["8"] = Rule{
				subRules: [][]string{{"42"}, {"42", "8"}},
			}
		} else if key == "11" {
			i.Rules["11"] = Rule{
				subRules: [][]string{{"42", "31"}, {"42", "11", "31"}},
			}
		}
	}
}

func (i Input) evaluateMessages() int {
	count := 0
	seen := map[string]bool{}

	for _, message := range i.Messages {
		seen = i.iterateMessages(i.Rules["0"].subRules[0][0], i.Rules["0"].subRules[0][1:], message, 0, seen)
		_, ok := seen[message]
		if ok {
			count++
		}
	}

	return count
}

func main() {
	rawInput := utils.ReadFile()
	i := &Input{
		Rules: map[string]Rule{},
	}
	i.parseInput(rawInput)
	fmt.Println("Part 1:", i.evaluateMessages())
	i.changeRulesForPart2()
	fmt.Println("Part 2:", i.evaluateMessages())
}
