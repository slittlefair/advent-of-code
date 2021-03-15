package main

import (
	"Advent-of-Code"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
)

func findNumbers(input interface{}) (numbers []int) {
	switch input := input.(type) {
	case []interface{}:
		for _, value := range input {
			numbers = append(numbers, findNumbers(value)...)
		}
	case map[string]interface{}:
		noRed := true
		for _, value := range input {
			if str, ok := value.(string); ok && str == "red" {
				noRed = false
				break
			}
		}
		if noRed {
			for _, value := range input {
				numbers = append(numbers, findNumbers(value)...)
			}
		}
	case float64:
		numbers = append(numbers, int(input))
	}

	return
}

func main() {
	re := regexp.MustCompile("-?\\d+")
	line := helpers.ReadFile()
	numbers := re.FindAllString(line[0], -1)
	total := 0
	for _, num := range numbers {
		total += helpers.StringToInt(num)
	}
	fmt.Println("Part 1:", total)

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{}, 0)
	json.Unmarshal(input, &data)

	sum := 0
	for _, num := range findNumbers(data) {
		sum += num
	}

	fmt.Println("Part 2:", sum)

}
