package main

import (
	helpers "Advent-of-Code"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func passwordComplete(password [8]string) bool {
	for _, p := range password {
		if p == "" {
			return false
		}
	}
	return true
}

func composePassword(password [8]string) string {
	p := ""
	for _, v := range password {
		p += v
	}
	return p
}

func printPassword(part int, password [8]string) string {
	p := fmt.Sprintf("\rPart %d: ", part)
	for _, v := range password {
		if v == "" {
			p += "_"
		} else {
			p += v
		}
	}
	return p
}

func findEasyPassword(id string) string {
	i := 0
	j := 0
	password := [8]string{}
	for {
		str := id + strconv.Itoa(i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(str)))
		if strings.HasPrefix(hash, "00000") {
			password[j] = string(hash[5])
			fmt.Print(printPassword(1, password))
			j++
		}
		if passwordComplete(password) {
			return composePassword(password)
		}
		i++
	}
}

func findDifficultPassword(id string) string {
	i := 0
	password := [8]string{}
	for {
		str := id + strconv.Itoa(i)
		i++
		hash := fmt.Sprintf("%x", md5.Sum([]byte(str)))
		if strings.HasPrefix(hash, "00000") {
			position, err := strconv.Atoi(string(hash[5]))
			if err != nil {
				continue
			}
			if position > 7 || password[position] != "" {
				continue
			}
			password[position] = string(hash[6])
			fmt.Print(printPassword(2, password))
		}
		if passwordComplete(password) {
			return composePassword(password)
		}
	}
}

func main() {
	input := helpers.ReadFile()[0]
	fmt.Print("Part 1: ________")
	fmt.Printf("\rPart 1: %s UNLOCKED\nPart 2: ________", findEasyPassword(input))
	fmt.Printf("\rPart 2: %s UNLOCKED\n", findDifficultPassword(input))
}
