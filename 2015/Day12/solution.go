package main

import (
	utils "Advent-of-Code/utils"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

func countNumbers(input string) int {
	re := regexp.MustCompile(`-?\d+`)
	nums := re.FindAllString(input, -1)
	sum := 0
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		sum += n
	}
	return sum
}

func findNonRedNumbers(input string) int {
	var result interface{}
	json.Unmarshal([]byte(input), &result)
	return recursive(result)
}

func recursive(r interface{}) int {
	count := 0
out:
	switch t := r.(type) {
	case float64:
		count += int(t)
	case []interface{}:
		for _, val := range t {
			count += recursive(val)
		}
	case map[string]interface{}:
		for _, val := range t {
			if val == "red" {
				break out
			}
		}
		for _, val := range t {
			count += recursive(val)
		}
	}
	return count
}

func main() {
	input := utils.ReadFile()[0]
	fmt.Println("Part 1:", countNumbers(input))
	fmt.Println("Part 2:", findNonRedNumbers(input))
}
