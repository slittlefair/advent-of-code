package main

import (
	"Advent-of-Code"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

var guards = make(map[string][60]int)

func main() {
	lines := helpers.ReadFile()
	sort.Sort(sort.Reverse(sort.StringSlice(lines)))

	reGuard := regexp.MustCompile("Guard \\#(\\d+)")
	reAsleep := regexp.MustCompile("asleep")
	reWakes := regexp.MustCompile("wakes")
	reMinute := regexp.MustCompile(":[\\d]+")

	var asleepMin, awakeMin int
	var sleepMins [60]int
	for _, entry := range lines {
		if ok := reWakes.MatchString(entry); ok {
			minute := reMinute.FindString(entry)[1:]
			m, err := strconv.Atoi(minute)
			helpers.Check(err)
			awakeMin = m
		} else if ok := reAsleep.MatchString(entry); ok {
			minute := reMinute.FindString(entry)[1:]
			m, err := strconv.Atoi(minute)
			helpers.Check(err)
			asleepMin = m
			for i := asleepMin; i < awakeMin; i++ {
				sleepMins[i]++
			}
		} else if ok := reGuard.MatchString(entry); ok {
			id := reGuard.FindString(entry)[7:]
			if val, ok := guards[id]; ok {
				for i := 0; i < len(sleepMins); i++ {
					sleepMins[i] = sleepMins[i] + val[i]
				}
				guards[id] = sleepMins
			} else {
				guards[id] = sleepMins
			}
			asleepMin, awakeMin = 0, 0
			sleepMins = [60]int{}
		}
	}
	var heaviestSleeper string
	var mostTimeSlept int
	for key, val := range guards {
		var time int
		for _, num := range val {
			time = time + num
		}
		if time > mostTimeSlept {
			heaviestSleeper = key
			mostTimeSlept = time
		}
	}
	var modalMinute, modalMinuteValue int
	for i, val := range guards[heaviestSleeper] {
		if val > modalMinuteValue {
			modalMinute = i
			modalMinuteValue = val
		}
	}
	guardID, err := strconv.Atoi(heaviestSleeper)
	helpers.Check(err)
	fmt.Println("Part 1:", guardID*modalMinute)

	// Part 2
	var commonestSleeper string
	var commonestMinute, commonestMinuteValue int
	for key, val := range guards {
		for i, num := range val {
			if num > commonestMinuteValue {
				commonestMinuteValue = num
				commonestMinute = i
				commonestSleeper = key
			}
		}
	}
	guardID, err = strconv.Atoi(commonestSleeper)
	helpers.Check(err)
	fmt.Println("Part 2:", guardID*commonestMinute)
}
