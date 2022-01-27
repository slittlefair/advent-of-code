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

func findNonRedNumbers(input string) (int, error) {
	var result interface{}
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return -1, err
	}
	return recursive(result), nil
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
	part2, err := findNonRedNumbers(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", part2)
}
