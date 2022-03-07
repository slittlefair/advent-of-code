package main

import (
	"Advent-of-Code/file"
	"crypto/md5"
	"fmt"
	"html"
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

func getEmoji(unicode string) string {
	xx := fmt.Sprintf("\\U000%s", unicode)
	// Hex String
	h := strings.ReplaceAll(xx, "\\U", "0x")
	// Hex to Int
	i, _ := strconv.ParseInt(h, 0, 64)
	// Unescape the string (HTML Entity -> String).
	return html.UnescapeString(string(rune(i)))
}

func lockEmoji() string {
	return fmt.Sprintf("%s locked", getEmoji("1F512"))
}

func unlockEmoji() string {
	return fmt.Sprintf("%s UNLOCKED", getEmoji("1F513"))
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
			fmt.Printf("%s %s", printPassword(1, password), lockEmoji())
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
			fmt.Printf("%s %s", printPassword(2, password), lockEmoji())
		}
		if passwordComplete(password) {
			return composePassword(password)
		}
	}
}

func main() {
	input := file.Read()[0]
	fmt.Printf("Part 1: ________ %s", lockEmoji())
	fmt.Printf("\rPart 1: %s %s\nPart 2: ________ %s", findEasyPassword(input), unlockEmoji(), lockEmoji())
	fmt.Printf("\rPart 2: %s %s\n", findDifficultPassword(input), unlockEmoji())
}
