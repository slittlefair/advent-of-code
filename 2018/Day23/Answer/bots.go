package day23

import "fmt"

type Bots map[Coordinate][]int

func NewBots(input []string) Bots {
	m := make(Bots)

	for _, data := range input {
		var r int
		var c Coordinate

		_, err := fmt.Sscanf(data, "pos=<%d,%d,%d>, r=%d", &c.X, &c.Y, &c.Z, &r)
		if err != nil {
			panic(err)
		}
		m[c] = append(m[c], r)
	}

	return m
}

func (m Bots) HaveInRange(pos Coordinate) int {
	var sum int

	for c, rs := range m {
		for _, r := range rs {
			if pos.Distance(c) <= r {
				sum++
			}
		}
	}

	return sum
}
