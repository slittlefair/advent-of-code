package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
)

func containsAtLeast3Vowels(re *regexp.Regexp, str string) bool {
	return len(re.FindAllString(str, -1)) >= 3
}

func containsDoubleLetter(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}
	return false
}

func doesNotContainBadString(re *regexp.Regexp, str string) bool {
	return !re.MatchString(str)
}

func doesNotContainAnyBadStrings(reGroup []*regexp.Regexp, str string) bool {
	for _, re := range reGroup {
		if !doesNotContainBadString(re, str) {
			return false
		}
	}
	return true
}

func isNicePart1(vowelRe *regexp.Regexp, badStringRe []*regexp.Regexp, str string) bool {
	if !containsAtLeast3Vowels(vowelRe, str) {
		return false
	}
	if !containsDoubleLetter(str) {
		return false
	}
	if !doesNotContainAnyBadStrings(badStringRe, str) {
		return false
	}
	return true
}

func containsRepeatedPairOfLetters(str string) bool {
	letterPairs := map[string]int{}
	for i := 0; i < len(str)-1; i++ {
		letterPair := fmt.Sprintf("%b%b", str[i], str[i+1])
		if _, ok := letterPairs[letterPair]; !ok {
			letterPairs[letterPair] = 1
		} else {
			letterPairs[letterPair]++
		}
		if i < len(str)-2 && str[i] == str[i+1] && str[i] == str[i+2] {
			i++
		}
	}
	for _, val := range letterPairs {
		if val > 1 {
			return true
		}
	}
	return false
}

func repeatsLetterWithGap(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}

func isNicePart2(str string) bool {
	if !containsRepeatedPairOfLetters(str) {
		return false
	}
	if !repeatsLetterWithGap(str) {
		return false
	}
	return true
}

func getNiceStringCount(input []string) (int, int) {
	var part1Count, part2Count int
	vowelRe := regexp.MustCompile(`a|e|i|o|u`)
	badStringRe := []*regexp.Regexp{
		regexp.MustCompile(`ab`),
		regexp.MustCompile(`cd`),
		regexp.MustCompile(`pq`),
		regexp.MustCompile(`xy`),
	}
	for _, str := range input {
		if isNicePart1(vowelRe, badStringRe, str) {
			part1Count++
		}
		if isNicePart2(str) {
			part2Count++
		}
	}
	return part1Count, part2Count
}

func main() {
	input := utils.ReadFile()
	part1Count, part2Count := getNiceStringCount(input)
	fmt.Println("Part 1:", part1Count)
	fmt.Println("Part 2:", part2Count)
}
