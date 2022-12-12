package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type Monkey struct {
	id           int
	items        []int
	operation    func(int) int
	test         int
	ifTrue       int
	ifFalse      int
	inspectCount int
}

type Troop []*Monkey

// Take a set of monkey instructions and create a monkey and its list of items. The items
// are returned separately so that after part 1 each monkey's items can be reset, allowing us to run
// part 2 without having to create all the monkeys again.
func createMonkey(input []string) (*Monkey, []int, error) {
	if len(input) != 6 {
		return nil, nil, fmt.Errorf("incorrect lines received, expected 6, got %d, %v", len(input), input)
	}
	monkey := &Monkey{}

	// Get ID
	_, err := fmt.Sscanf(input[0], "Monkey %d", &monkey.id)
	if err != nil {
		return nil, nil, err
	}

	// Get starting items
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(input[1], -1)
	items := []int{}
	for _, m := range matches {
		// We know all matched can be converted to ints due to regex matching, so we can safely
		// ignore the error
		itm, _ := strconv.Atoi(m)
		items = append(items, itm)
	}

	// Get operation
	if input[2] == "  Operation: new = old * old" {
		monkey.operation = func(old int) int { return old * old }
	} else {
		var op string
		var j int
		_, err := fmt.Sscanf(input[2], "  Operation: new = old %s %d", &op, &j)
		if err != nil {
			return nil, nil, err
		}
		switch op {
		case "+":
			monkey.operation = func(old int) int { return old + j }
		case "*":
			monkey.operation = func(old int) int { return old * j }
		default:
			return nil, nil, fmt.Errorf("invalid operation: %s", input[2])
		}
	}

	// Get test
	_, err = fmt.Sscanf(input[3], "  Test: divisible by %d", &monkey.test)
	if err != nil {
		return nil, nil, err
	}

	// Get ifTrue & ifFalse
	_, err = fmt.Sscanf(input[4], "    If true: throw to monkey %d", &monkey.ifTrue)
	if err != nil {
		return nil, nil, err
	}
	_, err = fmt.Sscanf(input[5], "    If false: throw to monkey %d", &monkey.ifFalse)
	if err != nil {
		return nil, nil, err
	}

	return monkey, items, nil
}

// Create a troop of monkeys and all items
func parseInput(input []string) (Troop, [][]int, error) {
	t := Troop{}
	var allItems [][]int
	for i := 0; i < len(input); i = i + 7 {
		monkey, items, err := createMonkey(input[i : i+6])
		if err != nil {
			return nil, nil, err
		}
		allItems = append(allItems, items)
		t = append(t, monkey)
	}
	return t, allItems, nil
}

func (t Troop) throwItems(n int, worryFunc func(worryLevel int) int) {
	for i := 0; i < n; i++ {
		for _, monkey := range t {
			for _, itm := range monkey.items {
				worryLevel := monkey.operation(itm)
				worryLevel = worryFunc(worryLevel)
				var toMonkey *Monkey
				if worryLevel%monkey.test == 0 {
					toMonkey = t[monkey.ifTrue]
				} else {
					toMonkey = t[monkey.ifFalse]
				}
				toMonkey.items = append(toMonkey.items, worryLevel)
				monkey.items = []int{}
				monkey.inspectCount++
			}
		}
	}
}

// Set all monkeys' items to their original and reset inspectCount
func (t Troop) initialiseTroop(allItems [][]int) {
	for i, monkey := range t {
		monkey.items = allItems[i]
		monkey.inspectCount = 0
	}
}

func (t Troop) doMonkeyBusiness(n int, worryFunc func(worryLevel int) int, allItems [][]int) int {
	t.initialiseTroop(allItems)
	t.throwItems(n, worryFunc)
	sort.Slice(t, func(i, j int) bool {
		return t[i].inspectCount > t[j].inspectCount
	})
	return t[0].inspectCount * t[1].inspectCount
}

func findSolutions(input []string) (int, int, error) {
	troop, allItems, err := parseInput(input)
	if err != nil {
		return -1, -1, err
	}
	part1 := troop.doMonkeyBusiness(20, func(worryLevel int) int { return worryLevel / 3 }, allItems)

	lcm := 1
	for _, monkey := range troop {
		lcm *= monkey.test
	}
	part2 := troop.doMonkeyBusiness(10000, func(worryLevel int) int { return worryLevel % lcm }, allItems)

	return part1, part2, nil
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
