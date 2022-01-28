package main

import (
	"Advent-of-Code/file"
	"fmt"
	"strings"
	"unicode"
)

type Replacements map[string][]string

type Medicine struct {
	Replacements         Replacements
	Molecule             string
	DistinctNewMolecules map[string]bool
}

func parseInput(input []string) *Medicine {
	med := &Medicine{
		Replacements:         make(Replacements),
		Molecule:             input[len(input)-1],
		DistinctNewMolecules: map[string]bool{},
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

func (m Medicine) FindIndicesOfStringInMolecule(s string) []int {
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
		indices := m.FindIndicesOfStringInMolecule(r)
		for _, re := range reps {
			for _, i := range indices {
				m.DistinctNewMolecules[m.Molecule[:i]+re+m.Molecule[i+len(r):]] = true
			}
		}
	}
}

func countUpper(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsUpper(char) {
			count++
		}
	}
	return count
}

// Each "token" ("H", "O", "Ci" etc.) in the molecule can be converted via a replacement except for
// "Rn", "Ar", and "Y". So to reverse engineer from molecule to "e" we just need to keep replacing
// tokens to shorten the length. We can do this to all except for those special tokens, so get the
// number of tokens (each token starts with a capital so count the number of upper case characters),
// subtract the number of Rn and Ar, subtract 2 lots of Y to include the characters that follow
// them, and subtract 1 since we start with e.
func (m Medicine) GetNumberOfSubs() int {
	RnCount := strings.Count(m.Molecule, "Rn")
	ArCount := strings.Count(m.Molecule, "Ar")
	YCount := strings.Count(m.Molecule, "Y")
	return countUpper(m.Molecule) - RnCount - ArCount - 2*YCount - 1
}

func main() {
	input := file.Read()
	med := parseInput(input)
	med.ReplaceAndFindNewMolecules()
	fmt.Println("Part 1:", len(med.DistinctNewMolecules))
	fmt.Println("Part 2:", med.GetNumberOfSubs())
}
