package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"Advent-of-Code/slice"
	"fmt"
	"strings"
)

const (
	elevator = iota
	microchip
	generator
)

var item_name = map[int]string{
	0: "elevator",
	1: "microchip",
	2: "generator",
}

type item struct {
	id       string
	name     string
	itemType int
	floor    int
	partner  *item
}

type items map[string]*item

func makePartnerID(itm *item) string {
	if itm.itemType == microchip {
		return fmt.Sprintf("%s-%s", itm.name, "generator")
	}
	return fmt.Sprintf("%s-%s", itm.name, "microchip")
}

func makeItem(name string, itemType, floor int) *item {
	var id string
	if name == "elevator" {
		id = "elevator"
	} else {
		id = fmt.Sprintf("%s-%s", name, item_name[itemType])
	}
	return &item{
		id:       id,
		name:     name,
		itemType: itemType,
		floor:    floor,
	}
}

func (is items) floorIsUnstable(floor int) bool {
	floorChips := []*item{}
	floorGenerators := []*item{}
	for _, i := range is {
		if i.itemType == elevator {
			continue
		}
		if i.floor == floor {
			if i.itemType == microchip {
				floorChips = append(floorChips, i)
			} else {
				floorGenerators = append(floorGenerators, i)
			}
		}
	}
	if len(floorGenerators) == 0 {
		return false
	}
	for _, i := range floorChips {
		if !slice.Contains(floorGenerators, i.partner) {
			return true
		}
	}
	return false
}

func (is items) moveItem(itm *item, newFloor int, formations []items) (items, bool) {
	if newFloor < 1 || newFloor > 4 {
		return is, false
	}
	oldFloor := itm.floor
	newItems := items{}
	for k, v := range is {
		newItems[k] = v
	}
	newItems[itm.id] = &item{
		id:       itm.id,
		name:     itm.name,
		itemType: itm.itemType,
		floor:    newFloor,
		partner:  itm.partner,
	}
	if newItems.floorIsUnstable(oldFloor) || newItems.floorIsUnstable(newFloor) {
		return is, false
	}
	newItems["elevator"].floor = newFloor
	if newItems.seenFormation(formations) {
		return is, false
	}
	return newItems, true
}

func parseInput(input []string) items {
	f := items{}
	for i, line := range input {
		split := strings.Split(line, " ")
		for j, w := range split {
			itm := &item{}
			if w == "microchip" || w == "microchip." {
				itm = makeItem(strings.Split(split[j-1], "-")[0], microchip, i+1)
				f[itm.id] = itm
			} else if w == "generator" || w == "generator." {
				itm = makeItem(split[j-1], generator, i+1)
				f[itm.id] = itm
			}
		}
	}
	for _, itm := range f {
		itm.partner = f[makePartnerID(itm)]
	}
	f["elevator"] = makeItem("elevator", elevator, 1)
	return f
}

func (is items) foundSolution() bool {
	for _, i := range is {
		if i.floor != 4 {
			return false
		}
	}
	return true
}

func (is items) print() {
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		for _, itm := range is {
			if itm.floor == i {
				fmt.Printf("%+#v\n", itm)
			}
		}
	}
	fmt.Println()
}

func (is items) seenFormation(formations []items) bool {
	for _, i := range is {
		fmt.Printf("%s: %d, ", i.id, i.floor)
	}
	fmt.Println()
	for _, f := range formations {
		for _, v := range f {
			fmt.Printf("%s: %d, ", v.id, v.floor)
		}
		fmt.Println()
		for k, v := range f {
			fmt.Printf("checking %s, %d == %d\n", k, is[k].floor, v.floor)
			if is[k].floor != v.floor {
				return false
			}
		}
	}
	return true
}

func (is items) findSolutions(formations []items, lowestScore int) ([]items, int) {
	if is.foundSolution() {
		return formations, maths.Min(lowestScore, len(formations))
	}
	currentFloor := is["elevator"].floor
	for _, itm := range is {
		if itm.itemType != elevator && itm.floor == currentFloor && len(formations) < 30 {
			newItems, moved := is.moveItem(itm, currentFloor+1, formations)
			if moved {
				formations = append(formations, newItems)
				formations, lowestScore = newItems.findSolutions(formations, lowestScore)
			}
			newItems, moved = is.moveItem(itm, currentFloor-1, formations)
			if moved {
				formations = append(formations, newItems)
				formations, lowestScore = newItems.findSolutions(formations, lowestScore)
			}
		}
	}
	return formations, lowestScore
}

func main() {
	input := file.Read()
	is := parseInput(input)
	formations, lowestScore := is.findSolutions([]items{is}, maths.Infinity)
	fmt.Println(lowestScore)
	formations[len(formations)-1].print()
}
