package main

import (
	"Advent-of-Code/file"
	"fmt"
	"strings"
)

type Monkey struct {
	id     string
	m1, m2 *Monkey
	value  int
}

type Monkeys map[string]Monkey

func parseInput(input []string) (Monkeys, error) {
	monkeys := Monkeys{}
	for len(monkeys) < len(input) {
		for _, line := range input {
			split := strings.Split(line, ": ")
			id := split[0]
			if _, ok := monkeys[id]; ok {
				continue
			}
			split = strings.Split(split[1], " ")
			if len(split) == 1 {
				var value int
				_, err := fmt.Sscanf(split[0], "%d", &value)
				if err != nil {
					return nil, err
				}
				monkeys[id] = Monkey{
					id:    id,
					value: value,
				}
				continue
			}
			if len(split) != 3 {
				return nil, fmt.Errorf("malformed line %s", line)
			}
			m1ID, op, m2ID := split[0], split[1], split[2]
			if _, ok := monkeys[m1ID]; !ok {
				continue
			}
			if _, ok := monkeys[m2ID]; !ok {
				continue
			}
			m1 := monkeys[m1ID]
			m2 := monkeys[m2ID]
			monkey := Monkey{
				id: id,
				m1: &m1,
				m2: &m2,
			}
			switch op {
			case "+":
				monkey.value = m1.value + m2.value
			case "-":
				monkey.value = m1.value - m2.value
			case "*":
				monkey.value = m1.value * m2.value
			case "/":
				monkey.value = m1.value / m2.value
			default:
				return nil, fmt.Errorf("invalid operator in line %s", line)
			}
			monkeys[id] = monkey
		}
	}
	return monkeys, nil
}

func parseInput2(input []string, humn int) (bool, error) {
	monkeys := Monkeys{
		"humn": Monkey{value: humn},
	}
	for len(monkeys) < len(input) {
		for _, line := range input {
			split := strings.Split(line, ": ")
			id := split[0]
			if _, ok := monkeys[id]; ok {
				continue
			}
			split = strings.Split(split[1], " ")
			if len(split) == 1 {
				var value int
				_, err := fmt.Sscanf(split[0], "%d", &value)
				if err != nil {
					return false, err
				}
				monkeys[id] = Monkey{
					id:    id,
					value: value,
				}
				continue
			}
			if len(split) != 3 {
				return false, fmt.Errorf("malformed line %s", line)
			}
			m1ID, op, m2ID := split[0], split[1], split[2]
			if _, ok := monkeys[m1ID]; !ok {
				continue
			}
			if _, ok := monkeys[m2ID]; !ok {
				continue
			}
			m1 := monkeys[m1ID]
			m2 := monkeys[m2ID]
			if id == "root" {
				return m1.value == m2.value, nil
			}
			monkey := Monkey{
				id: id,
				m1: &m1,
				m2: &m2,
			}

			switch op {
			case "+":
				monkey.value = m1.value + m2.value
			case "-":
				monkey.value = m1.value - m2.value
			case "*":
				monkey.value = m1.value * m2.value
			case "/":
				monkey.value = m1.value / m2.value
			default:
				return false, fmt.Errorf("invalid operator in line %s", line)
			}
			monkeys[id] = monkey
		}
	}
	return false, fmt.Errorf("valid value for humn not found")
}

func main() {
	input := file.Read()
	monkeys, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", monkeys["root"].value)
	// min, max := 0, 1000000
	// for i := min; i < max; i++ {
	// 	if maths.Modulo(i, 100) == 0 {
	// 		fmt.Println(i)
	// 	}
	// 	equal, err := parseInput2(input, i)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	if equal {
	// 		fmt.Println(monkeys["root"].m1.value, monkeys["root"].m2.value)
	// 		fmt.Println("Part 2:", i)
	// 		return
	// 	}
	// }
	// fmt.Println(fmt.Errorf("could not find solution in range %d to %d", min, max))
}
