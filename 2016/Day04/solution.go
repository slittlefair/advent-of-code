package main

import (
	"Advent-of-Code/cipher"
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int      { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key < p[j].Key
	}
	return p[i].Value > p[j].Value
}

func roomIsValid(line string, reString *regexp.Regexp) bool {
	matches := reString.FindAllString(line, -1)
	name := matches[:len(matches)-5]

	freq := map[string]int{}
	for _, l := range name {
		freq[l]++
	}
	p := make(PairList, len(freq))
	i := 0
	for k, v := range freq {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	topLetters := []string{}
	for i := 0; i < 5; i++ {
		topLetters = append(topLetters, p[i].Key)
	}
	checksum := matches[len(matches)-5:]

	return reflect.DeepEqual(topLetters, checksum)
}

func getValidRooms(input []string, reNum *regexp.Regexp) ([]string, int) {
	validRooms := []string{}
	sum := 0
	reString := regexp.MustCompile(`[a-z]`)
	for _, line := range input {
		if roomIsValid(line, reString) {
			validRooms = append(validRooms, strings.Split(line, "[")[0])
			num, _ := strconv.Atoi(reNum.FindString(line))
			sum += num
		}
	}
	return validRooms, sum
}

func validateRooms(input []string, reNum *regexp.Regexp) (string, error) {
	for _, room := range input {
		num := reNum.FindString(room)
		n, _ := strconv.Atoi(num)
		decrypted := cipher.CaesarCipher(room, n)
		if strings.Contains(decrypted, "northpole") {
			return num, nil
		}
	}
	return "", fmt.Errorf(`could not find "northpole" in any room`)
}

func main() {
	input := file.Read()
	validRooms, sum := getValidRooms(input, regex.MatchNums)
	fmt.Println("Part 1:", sum)

	id, err := validateRooms(validRooms, regex.MatchNums)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", id)
}
