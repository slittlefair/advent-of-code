package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var reParenthesis = regexp.MustCompile(`\((\d*|\*|\+)*\)`)
var reAddition = regexp.MustCompile(`\d+\+\d+`)
var reChars = regexp.MustCompile(`[0-9]+|\*|\+`)

func evaluateCoreSum(sum string) string {
	runningSumTotal := 0
	currentOperator := ""
	sumParts := reChars.FindAllString(sum, -1)
	for _, char := range sumParts {
		if s := char; s == "+" {
			currentOperator = "+"
		} else if s == "*" {
			currentOperator = "*"
		} else {
			// Because the input contains only valid characters this will never retrun an error.
			// Parenthesis are already handled in the surrounded methods so characters will only
			// ever be +, * or digits
			val, _ := strconv.Atoi(s)
			if runningSumTotal == 0 {
				runningSumTotal = val
			} else if currentOperator == "+" {
				runningSumTotal += val
			} else {
				runningSumTotal *= val
			}
		}
	}
	return strconv.Itoa(runningSumTotal)
}

func evaluateCoreSumAdditionFirst(sum string) string {
	handledAllAddition := false
	for !handledAllAddition {
		handledAllAddition = true
		sum = reAddition.ReplaceAllStringFunc(sum, evaluateCoreSum)
		for _, char := range sum {
			if string(char) == "+" {
				handledAllAddition = false
				break
			}
		}
	}
	return evaluateCoreSum(sum)
}

func evaluateSum(sum string, part int) string {
	handledAllParenthesis := false
	for !handledAllParenthesis {
		handledAllParenthesis = true
		if part == 1 {
			sum = reParenthesis.ReplaceAllStringFunc(sum, evaluateCoreSum)
		} else {
			sum = reParenthesis.ReplaceAllStringFunc(sum, evaluateCoreSumAdditionFirst)
		}
		for _, char := range sum {
			if string(char) == "(" {
				handledAllParenthesis = false
				break
			}
		}
	}
	if part == 1 {
		return evaluateCoreSum(sum)
	}
	return evaluateCoreSumAdditionFirst(sum)
}

func findSolutions(sums []string) (int, int) {
	runningTotalPart1 := 0
	runningTotalPart2 := 0
	for _, sum := range sums {
		sum = strings.ReplaceAll(sum, " ", "")
		// Because the input contains only valid characters these conversions will never return
		// errors
		val, _ := strconv.Atoi(evaluateSum(sum, 1))
		runningTotalPart1 += val
		val, _ = strconv.Atoi(evaluateSum(sum, 2))
		runningTotalPart2 += val
	}
	return runningTotalPart1, runningTotalPart2
}

func main() {
	sums := file.Read()
	part1Sol, part2Sol := findSolutions(sums)
	fmt.Println("Part 1:", part1Sol)
	fmt.Println("Part 2:", part2Sol)
}
