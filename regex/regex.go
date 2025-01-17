package regex

import "regexp"

// Match all numbers in a string
var MatchNums = regexp.MustCompile(`[-]?\d+`)

// Match all lower case letters in a string
var MatchLettersLower = regexp.MustCompile(`[a-z]+`)

// Match all upper case letters in a string
var MatchLettersUpper = regexp.MustCompile(`[A-Z]+`)

// Match all letters in a string
var MatchLettersAll = regexp.MustCompile(`[a-zA-Z]+`)

// Match all words in a string
var MatchWords = regexp.MustCompile(`\w+`)
