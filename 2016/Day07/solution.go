package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
)

func containsABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

func tlsValidation(s string, re *regexp.Regexp) bool {
	for _, hnMatch := range re.FindAllString(s, -1) {
		if containsABBA(hnMatch) {
			return false
		}
	}
	for _, match := range re.Split(s, -1) {
		if containsABBA(match) {
			return true
		}
	}
	return false
}

func compileBAB(ip []string) [][]byte {
	bab := [][]byte{}
	for _, s := range ip {
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+2] && s[i] != s[i+1] {
				bab = append(bab, []byte{s[i+1], s[i], s[i+1]})
			}
		}
	}
	return bab
}

func hasBABMatch(sn []string, bab [][]byte) bool {
	for _, s := range sn {
		for i := 0; i < len(s)-2; i++ {
			for _, b := range bab {
				if s[i] == b[0] && s[i+1] == b[1] && s[i+2] == b[2] {
					return true
				}
			}
		}
	}
	return false
}

func sslValidation(s string, re *regexp.Regexp) bool {
	return hasBABMatch(re.Split(s, -1), compileBAB(re.FindAllString(s, -1)))
}

func countValidIPs(input []string) (int, int) {
	var tlsCount, sslCount int
	re := regexp.MustCompile(`\[\w+\]`)
	for _, ip := range input {
		if tlsValidation(ip, re) {
			tlsCount++
		}
		if sslValidation(ip, re) {
			sslCount++
		}
	}
	return tlsCount, sslCount
}

func main() {
	input := file.Read()
	tlsValid, sslValid := countValidIPs(input)
	fmt.Println("Part 1:", tlsValid)
	fmt.Println("Part 2:", sslValid)
}
