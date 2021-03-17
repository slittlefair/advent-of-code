package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
	"strconv"
)

func validateYr(val string, min int, max int) (bool, error) {
	year, err := strconv.Atoi(val)
	if err != nil {
		return false, err
	}
	return year >= min && year <= max, nil
}

func validateHgt(val string) (bool, error) {
	re := regexp.MustCompile(`(\d+)(cm|in)`)
	match := re.FindStringSubmatch(val)
	if len(match) == 0 {
		return false, nil
	}

	if match[2] == "cm" {
		cm, err := strconv.Atoi(match[1])
		if err != nil {
			return false, err
		}
		return cm >= 150 && cm <= 193, nil
	}

	in, err := strconv.Atoi(match[1])
	if err != nil {
		return false, err
	}
	return in >= 59 && in <= 76, nil
}

func validateHcl(val string) bool {
	re := regexp.MustCompile(`#[0-9a-f]{6}`)
	return len(val) == 7 && re.MatchString(val)
}

func validateEcl(val string) bool {
	re := regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
	return re.MatchString(val)
}

func validatePid(val string) bool {
	re := regexp.MustCompile((`[0-9]{9}`))
	return len(val) == 9 && re.MatchString((val))
}

func allFieldsValid(fields map[string]string) (bool, error) {
	byrValid, err := validateYr(fields["byr"], 1920, 2002)
	if err != nil {
		return false, err
	}
	if !byrValid {
		return false, nil
	}

	iyrValid, err := validateYr(fields["iyr"], 2010, 2020)
	if err != nil {
		return false, err
	}
	if !iyrValid {
		return false, nil
	}

	eyrValid, err := validateYr(fields["eyr"], 2020, 2030)
	if err != nil {
		return false, err
	}
	if !eyrValid {
		return false, nil
	}

	hgtValid, err := validateHgt(fields["hgt"])
	if err != nil {
		return false, err
	}
	if !hgtValid {
		return false, nil
	}

	hclValid := validateHcl(fields["hcl"])
	if !hclValid {
		return false, nil
	}

	eclValid := validateEcl(fields["ecl"])
	if !eclValid {
		return false, nil
	}

	pidValid := validatePid((fields["pid"]))
	if !pidValid {
		return false, nil
	}

	return true, nil
}

func solution(entries []string) (int, int, error) {
	fields := map[string]string{}
	validPassportsPart1 := 0
	validPassportsPart2 := 0
	re := regexp.MustCompile(`(eyr|hcl|pid|ecl|byr|hgt|iyr):(\S+)`)
	for i, entry := range entries {
		matches := re.FindAllStringSubmatch(entry, -1)
		for _, match := range matches {
			fields[match[1]] = match[2]
		}
		if len(entry) == 0 || i == len(entries)-1 {
			if len(fields) == 7 {
				validPassportsPart1++
				allValid, err := allFieldsValid(fields)
				if err != nil {
					return 0, 0, err
				}
				if allValid {
					validPassportsPart2++
				}
			}
			fields = map[string]string{}
		}
	}
	return validPassportsPart1, validPassportsPart2, nil
}

func main() {
	entries := helpers.ReadFile()
	validPassportsPart1, validPassportsPart2, err := solution(entries)
	if err != nil {
		return
	}
	fmt.Println("Part 1:", validPassportsPart1)
	fmt.Println("Part 2:", validPassportsPart2)
}
