package main

import (
	"Advent-of-Code/file"
	"crypto/md5"
	"fmt"
	"sort"
)

func hashContainsTriple(hash string) (bool, byte) {
	for i := 0; i < len(hash)-2; i++ {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] {
			return true, hash[i]
		}
	}
	return false, 0
}

func hashContainsQuintuple(hash string, s byte) bool {
	for i := 0; i < len(hash)-4; i++ {
		isMatch := true
		for j := i; j < i+5; j++ {
			if hash[j] != s {
				isMatch = false
			}
		}
		if isMatch {
			return true
		}
	}
	return false
}

func findKeys(salt string, part2 bool) int {
	triples := map[int]byte{}
	keys := []int{}
	i := 0
	for {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", salt, i))))
		// If it's part2 hash another 2016 times
		if part2 {
			for j := 0; j < 2016; j++ {
				hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
			}
		}
		// Delete any map entry for the current iteration (noop if it doesn't exist)
		delete(triples, i)
		// If any triples have a quintuple in the current hash then they're a valid key
		for k, v := range triples {
			if hashContainsQuintuple(hash, v) {
				keys = append(keys, k-1000)
				delete(triples, k)
			}
		}
		// If we've found at least 64 keys then we exit
		if len(keys) >= 64 {
			sort.Ints(keys)
			return keys[63]
		}
		// If the hash contains a triple then add the index+1000 to triples with the repeated
		// character. We'll then keep checking it for the next 1000 iterations and delete it if we
		// still haven't found a quintuple for the given character.
		if containsTriple, val := hashContainsTriple(hash); containsTriple {
			triples[i+1000] = val
		}
		i++
	}
}

func main() {
	input := file.Read()[0]
	fmt.Println("Part 1:", findKeys(input, false))
	fmt.Println("Part 2:", findKeys(input, true))
}
