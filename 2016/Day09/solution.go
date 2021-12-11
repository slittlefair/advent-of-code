package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
)

func handleMarker(s string, startingIndex int, re *regexp.Regexp) (string, int) {
	values := re.FindAllString(s, 2)
	fmt.Println("handleMarker", s)
	numValues := helpers.StringSliceToIntSlice(values)
	var endIndex int
	for i := 0; i < len(s); i++ {
		if string(s[i]) == ")" {
			endIndex = i + 1
			break
		}
	}
	repeatedString := s[endIndex : endIndex+numValues[0]]
	newString := ""
	fmt.Println(repeatedString, numValues)
	for i := 0; i < numValues[1]; i++ {
		newString += repeatedString
		fmt.Println(newString)
	}
	return newString, startingIndex + endIndex + numValues[0]
}

func decompress(s string, re *regexp.Regexp) string {
	newString := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		if string(char) == "(" {
			postMarker, j := handleMarker(s[i:], i, re)
			i = j
			newString += postMarker
		} else {
			newString += string(char)
		}
	}
	return newString
}

func main() {
	input := helpers.ReadFile()
	re := regexp.MustCompile(`\d+`)
	for _, i := range input {
		newString := decompress(i, re)
		fmt.Println(i, newString, len(newString))
	}
}
