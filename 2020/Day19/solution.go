package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
	"strings"
)

type Rules struct {
	NewRules    []string
	Evaluated   map[string][][]string
	Unevaluated map[string][][]string
}

type Input struct {
	Rules    Rules
	Messages []string
}

var re = regexp.MustCompile(`\d+`)
var reAMatch = regexp.MustCompile("a")
var reBMatch = regexp.MustCompile("b")

func (i *Input) parseInput(rawInput []string) {
	for _, line := range rawInput {
		if line == "" {
			continue
		}
		split := strings.Split(line, ":")
		if len(split) == 1 {
			i.Messages = append(i.Messages, line)
		} else {
			if reAMatch.MatchString(split[1]) {
				i.Rules.Evaluated[split[0]] = [][]string{{"a"}}
				i.Rules.NewRules = append(i.Rules.NewRules, split[0])
			} else if reBMatch.MatchString(split[1]) {
				i.Rules.Evaluated[split[0]] = [][]string{{"b"}}
				i.Rules.NewRules = append(i.Rules.NewRules, split[0])
			} else {
				i.Rules.Unevaluated[split[0]] = [][]string{}
				components := strings.Split(split[1], "|")
				for _, c := range components {
					i.Rules.Unevaluated[split[0]] = append(i.Rules.Unevaluated[split[0]], re.FindAllString(c, -1))
				}
			}
		}
	}
}

func allNumsEvaluated(rules [][]string) bool {
	for _, arr := range rules {
		for _, str := range arr {
			if !reAMatch.MatchString(str) && !reBMatch.MatchString(str) {
				return false
			}
		}
	}
	return true
}

func combineStrings(rules [][]string) [][]string {
	strings := [][]string{}
	for _, arr := range rules {
		s := ""
		for _, str := range arr {
			s += str
		}
		strings = append(strings, []string{s})
	}
	return strings
}

func (input *Input) evaluateRules() {
	newRules := []string{}
	for _, newRule := range input.Rules.NewRules {
		for ruleNum, rules := range input.Rules.Unevaluated {
			for _, ruleSet := range rules {
				for i, r := range ruleSet {
					if r == newRule {
						for _, evaluated := range input.Rules.Evaluated[newRule] {
							for _, ev := range evaluated {
								//
								for j, rAgain := range ruleSet {
									if r != rAgain && (reAMatch.MatchString(rAgain) && reBMatch.MatchString(rAgain)) {
										newRuleSet := ruleSet
										newRuleSet[j] = newRuleSet[j] + rAgain
										ruleSet = append(ruleSet, newRuleSet...)
									}
								}
								ruleSet[i] = ev
							}
						}
						if allNumsEvaluated(rules) {
							newRules = append(newRules, ruleNum)
							input.Rules.Evaluated[ruleNum] = combineStrings(rules)
						}
					}
				}
			}
		}
	}
	input.Rules.NewRules = newRules
}

func main() {
	rawInput := helpers.ReadFile()
	i := &Input{
		Rules: Rules{
			Evaluated:   make(map[string][][]string),
			Unevaluated: make(map[string][][]string),
		},
		Messages: []string{},
	}
	i.parseInput(rawInput)
	i.evaluateRules()
	fmt.Println(*i)
	i.evaluateRules()
	fmt.Println(*i)
}
