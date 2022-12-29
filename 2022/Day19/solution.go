package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
)

type State struct {
	minute   int
	ore      int
	clay     int
	obsidian int
	geode    int
	robots   [4]int
}

type Material int

const (
	Ore Material = iota
	Clay
	Obsidian
	Geode
)

type Robot struct {
	robotType    Material
	oreCost      int
	clayCost     int
	obsidianCost int
}

type Blueprints []*Blueprint

type Cache map[State]bool

type Blueprint struct {
	id            int
	oreRobot      Robot
	clayRobot     Robot
	obsidianRobot Robot
	geodeRobot    Robot
	state         State
	qualityLevel  int
	cache         Cache
}

func parseInput(input []string) (Blueprints, error) {
	bps := Blueprints{}
	for _, line := range input {
		bp := &Blueprint{
			oreRobot:      Robot{robotType: Ore},
			clayRobot:     Robot{robotType: Clay},
			obsidianRobot: Robot{robotType: Obsidian},
			geodeRobot:    Robot{robotType: Geode},
			state:         State{robots: [4]int{1, 0, 0, 0}},
			cache:         make(Cache),
		}
		_, err := fmt.Sscanf(
			line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&bp.id,
			&bp.oreRobot.oreCost,
			&bp.clayRobot.oreCost,
			&bp.obsidianRobot.oreCost,
			&bp.obsidianRobot.clayCost,
			&bp.geodeRobot.oreCost,
			&bp.geodeRobot.obsidianCost,
		)
		if err != nil {
			return nil, err
		}
		bps = append(bps, bp)
	}
	return bps, nil
}

func (s *State) canAffordRobot(r Robot) bool {
	if r.oreCost > s.ore {
		return false
	}
	if r.clayCost > s.clay {
		return false
	}
	if r.obsidianCost > s.obsidian {
		return false
	}
	return true
}

func (s *State) buildRobot(r Robot) {
	s.ore -= r.oreCost
	s.clay -= r.clayCost
	s.obsidian -= r.obsidianCost
	s.robots[r.robotType]++
}

func (s *State) collectMinerals() {
	s.ore += s.robots[0]
	s.clay += s.robots[1]
	s.obsidian += s.robots[2]
	s.geode += s.robots[3]
}

func (b Blueprint) seenState(state State) bool {
	_, ok := b.cache[state]
	return ok
}

func (s State) maxGeodesCanBuild() int {
	sum := s.geode
	for i := 0; i < 24-s.minute; i++ {
		sum += s.robots[3] + i
	}
	return sum
}

func (b *Blueprint) oneMinute(state State) {
	if _, ok := b.cache[state]; ok {
		// fmt.Println("seen", state)
		return
	}
	if state.maxGeodesCanBuild() <= b.qualityLevel {
		return
	}
	state.minute++
	b.cache[state] = true
	if state.minute == 24 {
		// fmt.Printf("%+#v\n", state)
		// if state.geode > 1 {
		// 	fmt.Println(state)
		// }
		if maths.Max(b.qualityLevel, b.id*state.geode) != b.qualityLevel {
			fmt.Println("new best", b.id*state.geode)
			fmt.Printf("%+#v\n", state)
		}
		b.qualityLevel = maths.Max(b.qualityLevel, b.id*state.geode)
		return
	}
	// if state.minute == 20 {
	// 	fmt.Printf("XXX %+#v\n", state)
	// }
	state.collectMinerals()
	for _, robot := range []Robot{b.geodeRobot, b.obsidianRobot, b.clayRobot, b.oreRobot} {
		if state.canAffordRobot(robot) {
			state.buildRobot(robot)
		}
		b.oneMinute(state)
		// b = &newBlueprint
	}
	b.oneMinute(state)
	// newBlueprint := Blueprint{
	// 	id:            b.id,
	// 	oreRobot:      b.oreRobot,
	// 	clayRobot:     b.clayRobot,
	// 	obsidianRobot: b.obsidianRobot,
	// 	geodeRobot:    b.geodeRobot,
	// 	state:         newState,
	// 	QualityLevel:  b.QualityLevel,
	// }
	// newBlueprint.oneMinute()
	// b = &newBlueprint
}

func (bps Blueprints) findSolution() int {
	ql := 0
	for _, b := range bps {
		b.oneMinute(b.state)
		ql += b.qualityLevel
	}
	return ql
}

func main() {
	input := file.Read()
	blueprints, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	part1 := blueprints.findSolution()
	fmt.Println("Part 1:", part1)
	for _, bp := range blueprints {
		fmt.Printf("%+#v\n", bp.state)
	}
}
