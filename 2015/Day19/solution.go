package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strings"
)

type Replacements map[string][]string

type Medicine struct {
	Replacements         Replacements
	Molecule             string
	NewMolecules         []string
	DistinctNewMolecules map[string]bool
}

func parseInput(input []string) *Medicine {
	med := &Medicine{
		Replacements:         make(Replacements),
		Molecule:             input[len(input)-1],
		DistinctNewMolecules: make(map[string]bool),
	}
	for i := 0; i < len(input)-2; i++ {
		line := input[i]
		split := strings.Split(line, "=>")
		key := strings.TrimSpace(split[0])
		val := strings.TrimSpace(split[1])
		if _, ok := med.Replacements[key]; !ok {
			med.Replacements[key] = []string{val}
		} else {
			med.Replacements[key] = append(med.Replacements[key], val)
		}
	}
	return med
}

func (m Medicine) FindIndicesOfSringInMolecule(s string) []int {
	index := strings.Index(m.Molecule, s)
	indices := []int{}
	offset := 0
	for index > -1 {
		indices = append(indices, index+offset)
		offset += len(m.Molecule[:index+len(s)])
		m.Molecule = m.Molecule[index+len(s):]
		index = strings.Index(m.Molecule, s)
	}
	return indices
}

func (m *Medicine) ReplaceAndFindNewMolecules() {
	for r, reps := range m.Replacements {
		indices := m.FindIndicesOfSringInMolecule(r)
		for _, re := range reps {
			for _, i := range indices {
				m.DistinctNewMolecules[m.Molecule[:i]+re+m.Molecule[i+len(r):]] = true
			}
		}
	}
}

func main() {
	input := helpers.ReadFile()
	med := parseInput(input)
	med.ReplaceAndFindNewMolecules()
	fmt.Println("Part 1:", len(med.DistinctNewMolecules))
}
