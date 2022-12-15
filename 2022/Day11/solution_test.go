package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertTroopEqual(t *testing.T, troop, want Troop, allItems [][]int) {
	for i, monkey := range troop {
		assert.Equal(t, want[i].items, allItems[i])
		assert.Equal(t, want[i].id, monkey.id)
		assert.Equal(t, want[i].test, monkey.test)
		assert.Equal(t, want[i].ifTrue, monkey.ifTrue)
		assert.Equal(t, want[i].ifFalse, monkey.ifFalse)
		assert.Equal(t, want[i].inspectCount, monkey.inspectCount)
	}
}

func TestCreateMonkey(t *testing.T) {
	errTests := []struct {
		name  string
		input []string
	}{
		{
			name:  "input is not 6 lines",
			input: []string{"create this monkey"},
		},
		{
			name: "monkey line is in an incorrect format",
			input: []string{
				"First Monkey:",
				"   Starting items: 79, 98",
				"   Operation: new = old * 19",
				"   Test: divisible by 23",
				"     If true: throw to monkey 2",
				"     If false: throw to monkey 3",
			},
		},
		{
			name: "operation line is in an incorrect format",
			input: []string{
				"Monkey 0:",
				"   Starting items: 79, 98",
				"   Op: new = old * 19",
				"   Test: divisible by 23",
				"     If true: throw to monkey 2",
				"     If false: throw to monkey 3",
			},
		},
		{
			name: "operation line contains an invalid operator",
			input: []string{
				"Monkey 0:",
				"   Starting items: 79, 98",
				"   Operation: new = old - 19",
				"   Test: divisible by 23",
				"     If true: throw to monkey 2",
				"     If false: throw to monkey 3",
			},
		},
		{
			name: "test line is in an incorrect format",
			input: []string{
				"Monkey 0:",
				"   Starting items: 79, 98",
				"   Operation: new = old * 19",
				"   Test: divide by 23",
				"     If true: throw to monkey 2",
				"     If false: throw to monkey 3",
			},
		},
		{
			name: `"if true" line is in an incorrect format`,
			input: []string{
				"Monkey 0:",
				"   Starting items: 79, 98",
				"   Operation: new = old * 19",
				"   Test: divisible by 23",
				"     If true: throw to number 2",
				"     If false: throw to monkey 3",
			},
		},
		{
			name: `"if false" line is in an incorrect format`,
			input: []string{
				"Monkey 0:",
				"   Starting items: 79, 98",
				"   Operation: new = old * 19",
				"   Test: divisible by 23",
				"     If true: throw to monkey 2",
				"     If false: throw to numero 3",
			},
		},
	}

	for _, tt := range errTests {
		t.Run(fmt.Sprintf("returns an error if %s", tt.name), func(t *testing.T) {
			monkey, allItems, err := createMonkey(tt.input)
			assert.Nil(t, monkey)
			assert.Nil(t, allItems)
			assert.Error(t, err)
		})
	}

	tests := []struct {
		input []string
		want  *Monkey
	}{
		{
			input: []string{
				"Monkey 0:",
				"   Starting items: 79, 98",
				"   Operation: new = old * 19",
				"   Test: divisible by 23",
				"     If true: throw to monkey 2",
				"     If false: throw to monkey 3",
			},
			want: &Monkey{
				id:      0,
				items:   []int{79, 98},
				test:    23,
				ifTrue:  2,
				ifFalse: 3,
			},
		},
		{
			input: []string{
				"Monkey 1:",
				"  Starting items: 54, 65, 75, 74",
				"  Operation: new = old + 6",
				"  Test: divisible by 19",
				"    If true: throw to monkey 2",
				"    If false: throw to monkey 0",
			},
			want: &Monkey{
				id:      1,
				items:   []int{54, 65, 75, 74},
				test:    19,
				ifTrue:  2,
				ifFalse: 0,
			},
		},
		{
			input: []string{
				"Monkey 2:",
				"  Starting items: 79, 60, 97",
				"  Operation: new = old * old",
				"  Test: divisible by 13",
				"    If true: throw to monkey 1",
				"    If false: throw to monkey 3",
			},
			want: &Monkey{
				id:      2,
				items:   []int{79, 60, 97},
				test:    13,
				ifTrue:  1,
				ifFalse: 3,
			},
		},
		{
			input: []string{
				"Monkey 3:",
				"  Starting items: 74",
				"  Operation: new = old + 3",
				"  Test: divisible by 17",
				"    If true: throw to monkey 0",
				"    If false: throw to monkey 1",
			},
			want: &Monkey{
				id:      3,
				items:   []int{74},
				test:    17,
				ifTrue:  0,
				ifFalse: 1,
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("creates a monkey from input, advent of code example monkey %d", i), func(t *testing.T) {
			monkey, allItems, err := createMonkey(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.want.items, allItems)
			assert.Equal(t, tt.want.id, monkey.id)
			assert.Equal(t, tt.want.test, monkey.test)
			assert.Equal(t, tt.want.ifTrue, monkey.ifTrue)
			assert.Equal(t, tt.want.ifFalse, monkey.ifFalse)
		})
	}
}

func TestParseInput(t *testing.T) {
	t.Run("returns an error in an input line is invalid", func(t *testing.T) {
		input := []string{
			"Monkey 0:",
			"  Starting items: 79, 98",
			"  Operation: new = old * 19",
			"  Test: divisible by 23",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 3",
			"",
			"Monkey 1:",
			"  Starting items: 54, 65, 75, 74",
			"  Operation: new = old + 6",
			"  Test: divide by 19",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 0",
			"",
		}
		troop, allItems, err := parseInput(input)
		assert.Nil(t, troop)
		assert.Nil(t, allItems)
		assert.Error(t, err)
	})

	t.Run("returns a troop parsed from input", func(t *testing.T) {
		input := []string{
			"Monkey 0:",
			"  Starting items: 79, 98",
			"  Operation: new = old * 19",
			"  Test: divisible by 23",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 3",
			"",
			"Monkey 1:",
			"  Starting items: 54, 65, 75, 74",
			"  Operation: new = old + 6",
			"  Test: divisible by 19",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 0",
			"",
			"Monkey 2:",
			"  Starting items: 79, 60, 97",
			"  Operation: new = old * old",
			"  Test: divisible by 13",
			"    If true: throw to monkey 1",
			"    If false: throw to monkey 3",
			"",
			"Monkey 3:",
			"  Starting items: 74",
			"  Operation: new = old + 3",
			"  Test: divisible by 17",
			"    If true: throw to monkey 0",
			"    If false: throw to monkey 1",
		}
		want := Troop{
			{
				id:      0,
				items:   []int{79, 98},
				test:    23,
				ifTrue:  2,
				ifFalse: 3,
			},
			{
				id:      1,
				items:   []int{54, 65, 75, 74},
				test:    19,
				ifTrue:  2,
				ifFalse: 0,
			},
			{
				id:      2,
				items:   []int{79, 60, 97},
				test:    13,
				ifTrue:  1,
				ifFalse: 3,
			},
			{
				id:      3,
				items:   []int{74},
				test:    17,
				ifTrue:  0,
				ifFalse: 1,
			},
		}
		troop, allItems, err := parseInput(input)
		assertTroopEqual(t, troop, want, allItems)
		assert.NoError(t, err)
	})
}

func TestThrowItems(t *testing.T) {
	allItems := [][]int{
		{79, 98},
		{54, 65, 75, 74},
		{79, 60, 97},
		{74},
	}
	troop := &Troop{
		{
			id:        0,
			operation: func(old int) int { return old * 19 },
			test:      23,
			ifTrue:    2,
			ifFalse:   3,
		},
		{
			id:        1,
			operation: func(old int) int { return old + 6 },
			test:      19,
			ifTrue:    2,
			ifFalse:   0,
		},
		{
			id:        2,
			operation: func(old int) int { return old * old },
			test:      13,
			ifTrue:    1,
			ifFalse:   3,
		},
		{
			id:        3,
			operation: func(old int) int { return old + 3 },
			test:      17,
			ifTrue:    0,
			ifFalse:   1,
		},
	}
	lcm := 1
	for _, monkey := range *troop {
		lcm *= monkey.test
	}

	part1Func := func(worryLevel int) int { return worryLevel / 3 }
	part2Func := func(worryLevel int) int { return worryLevel % lcm }

	tests := []struct {
		n         int
		worryFunc func(worryLevel int) int
		want      []int
	}{
		{
			n:         20,
			worryFunc: part1Func,
			want:      []int{101, 95, 7, 105},
		},
		{
			n:         1,
			worryFunc: part2Func,
			want:      []int{2, 4, 3, 6},
		},
		{
			n:         20,
			worryFunc: part2Func,
			want:      []int{99, 97, 8, 103},
		},
		{
			n:         1000,
			worryFunc: part2Func,
			want:      []int{5204, 4792, 199, 5192},
		},
		{
			n:         2000,
			worryFunc: part2Func,
			want:      []int{10419, 9577, 392, 10391},
		},
		{
			n:         3000,
			worryFunc: part2Func,
			want:      []int{15638, 14358, 587, 15593},
		},
		{
			n:         4000,
			worryFunc: part2Func,
			want:      []int{20858, 19138, 780, 20797},
		},
		{
			n:         5000,
			worryFunc: part2Func,
			want:      []int{26075, 23921, 974, 26000},
		},
		{
			n:         6000,
			worryFunc: part2Func,
			want:      []int{31294, 28702, 1165, 31204},
		},
		{
			n:         7000,
			worryFunc: part2Func,
			want:      []int{36508, 33488, 1360, 36400},
		},
		{
			n:         8000,
			worryFunc: part2Func,
			want:      []int{41728, 38268, 1553, 41606},
		},
		{
			n:         9000,
			worryFunc: part2Func,
			want:      []int{46945, 43051, 1746, 46807},
		},
		{
			n:         10000,
			worryFunc: part2Func,
			want:      []int{52166, 47830, 1938, 52013},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("correctly keeps track of inspected items after %d iteration", tt.n), func(t *testing.T) {
			for i, monkey := range *troop {
				monkey.items = allItems[i]
				monkey.inspectCount = 0
			}
			troop.throwItems(tt.n, tt.worryFunc)
			for i, monkey := range *troop {
				assert.Equal(t, tt.want[i], monkey.inspectCount)
			}
		})
	}
}

func TestInitialiseTroop(t *testing.T) {
	allItems := [][]int{
		{79, 98},
		{54, 65, 75, 74},
		{79, 60, 97},
		{74},
	}

	want := &Troop{
		{
			id:           0,
			operation:    func(old int) int { return old * 19 },
			test:         23,
			ifTrue:       2,
			ifFalse:      3,
			inspectCount: 0,
			items:        []int{79, 98},
		},
		{
			id:           1,
			operation:    func(old int) int { return old + 6 },
			test:         19,
			ifTrue:       2,
			ifFalse:      0,
			inspectCount: 0,
			items:        []int{54, 65, 75, 74},
		},
		{
			id:           2,
			operation:    func(old int) int { return old * old },
			test:         13,
			ifTrue:       1,
			ifFalse:      3,
			inspectCount: 0,
			items:        []int{79, 60, 97},
		},
		{
			id:           3,
			operation:    func(old int) int { return old + 3 },
			test:         17,
			ifTrue:       0,
			ifFalse:      1,
			inspectCount: 0,
			items:        []int{74},
		},
	}

	t.Run("initialises a starting troop with all items and inspect count", func(t *testing.T) {
		troop := &Troop{
			{
				id:        0,
				operation: func(old int) int { return old * 19 },
				test:      23,
				ifTrue:    2,
				ifFalse:   3,
			},
			{
				id:        1,
				operation: func(old int) int { return old + 6 },
				test:      19,
				ifTrue:    2,
				ifFalse:   0,
			},
			{
				id:        2,
				operation: func(old int) int { return old * old },
				test:      13,
				ifTrue:    1,
				ifFalse:   3,
			},
			{
				id:        3,
				operation: func(old int) int { return old + 3 },
				test:      17,
				ifTrue:    0,
				ifFalse:   1,
			},
		}
		troop.initialiseTroop(allItems)
		assertTroopEqual(t, *troop, *want, allItems)
	})

	t.Run("reinitialises a troop with all items and inspect count", func(t *testing.T) {
		troop := &Troop{
			{
				id:           0,
				operation:    func(old int) int { return old * 19 },
				test:         23,
				ifTrue:       2,
				ifFalse:      3,
				inspectCount: 10,
				items:        []int{21, 23},
			},
			{
				id:           1,
				operation:    func(old int) int { return old + 6 },
				test:         19,
				ifTrue:       2,
				ifFalse:      0,
				inspectCount: 122,
				items:        []int{},
			},
			{
				id:           2,
				operation:    func(old int) int { return old * old },
				test:         13,
				ifTrue:       1,
				ifFalse:      3,
				inspectCount: 9,
				items:        []int{98, 56, 43, 29},
			},
			{
				id:           3,
				operation:    func(old int) int { return old + 3 },
				test:         17,
				ifTrue:       0,
				ifFalse:      1,
				inspectCount: 76,
				items:        []int{100, 101, 109, 108, 11111},
			},
		}
		troop.initialiseTroop(allItems)
		assertTroopEqual(t, *troop, *want, allItems)
	})
}

func TestDoMonkeyBusiness(t *testing.T) {
	t.Run("does monkey business, advent of code example 1", func(t *testing.T) {
		troop := Troop{
			{
				id:        0,
				operation: func(old int) int { return old * 19 },
				test:      23,
				ifTrue:    2,
				ifFalse:   3,
			},
			{
				id:        1,
				operation: func(old int) int { return old + 6 },
				test:      19,
				ifTrue:    2,
				ifFalse:   0,
			},
			{
				id:        2,
				operation: func(old int) int { return old * old },
				test:      13,
				ifTrue:    1,
				ifFalse:   3,
			},
			{
				id:        3,
				operation: func(old int) int { return old + 3 },
				test:      17,
				ifTrue:    0,
				ifFalse:   1,
			},
		}
		allItems := [][]int{
			{79, 98},
			{54, 65, 75, 74},
			{79, 60, 97},
			{74},
		}

		got := troop.doMonkeyBusiness(20, func(worryLevel int) int { return worryLevel / 3 }, allItems)
		assert.Equal(t, 10605, got)
	})

	t.Run("does monkey business, advent of code example 2", func(t *testing.T) {
		troop := Troop{
			{
				id:        0,
				operation: func(old int) int { return old * 19 },
				test:      23,
				ifTrue:    2,
				ifFalse:   3,
			},
			{
				id:        1,
				operation: func(old int) int { return old + 6 },
				test:      19,
				ifTrue:    2,
				ifFalse:   0,
			},
			{
				id:        2,
				operation: func(old int) int { return old * old },
				test:      13,
				ifTrue:    1,
				ifFalse:   3,
			},
			{
				id:        3,
				operation: func(old int) int { return old + 3 },
				test:      17,
				ifTrue:    0,
				ifFalse:   1,
			},
		}
		allItems := [][]int{
			{79, 98},
			{54, 65, 75, 74},
			{79, 60, 97},
			{74},
		}

		lcm := 1
		for _, monkey := range troop {
			lcm *= monkey.test
		}

		got := troop.doMonkeyBusiness(10000, func(worryLevel int) int { return worryLevel % lcm }, allItems)
		assert.Equal(t, 2713310158, got)
	})
}

func TestFindSolutions(t *testing.T) {
	t.Run("returns an error if input can't be parsed", func(t *testing.T) {
		input := []string{
			"Monkey 0:",
			"  Starting items: 79, 98",
			"  Operation: new = old * 19",
			"  Test: divisible by 23",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 3",
			"",
			"Monkey 1:",
			"  Starting items: 54, 65, 75, 74",
			"  Operation: new = old + 6",
			"  Test: divide by 19",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 0",
			"",
		}
		got, got1, err := findSolutions(input)
		assert.Error(t, err)
		assert.Equal(t, -1, got)
		assert.Equal(t, -1, got1)
	})

	t.Run("returns solutions, advent of code example", func(t *testing.T) {
		input := []string{
			"Monkey 0:",
			"  Starting items: 79, 98",
			"  Operation: new = old * 19",
			"  Test: divisible by 23",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 3",
			"",
			"Monkey 1:",
			"  Starting items: 54, 65, 75, 74",
			"  Operation: new = old + 6",
			"  Test: divisible by 19",
			"    If true: throw to monkey 2",
			"    If false: throw to monkey 0",
			"",
			"Monkey 2:",
			"  Starting items: 79, 60, 97",
			"  Operation: new = old * old",
			"  Test: divisible by 13",
			"    If true: throw to monkey 1",
			"    If false: throw to monkey 3",
			"",
			"Monkey 3:",
			"  Starting items: 74",
			"  Operation: new = old + 3",
			"  Test: divisible by 17",
			"    If true: throw to monkey 0",
			"    If false: throw to monkey 1",
		}
		got, got1, err := findSolutions(input)
		assert.NoError(t, err)
		assert.Equal(t, 10605, got)
		assert.Equal(t, 2713310158, got1)
	})
}
