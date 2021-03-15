// package main

// import (
// 	"container/heap"
// 	"flag"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"strconv"
// 	"strings"
// )

// var inputFn = flag.String("input", "input", "Input filename")
// var part = flag.String("part", "both", "Which problem part to run")
// var trace = flag.Bool("trace", false, "Enable debug tracing")
// var buffer = flag.Int("buffer", 150, "The amount of buffer to add to each dimension when generating the map")

// func main() {
// 	flag.Parse()
// 	if *trace {
// 		// log.SetLevel(log.DebugLevel)
// 	}

// 	if *part == "both" || *part == "a" {
// 		partA()
// 	}

// 	if *part == "both" || *part == "b" {
// 		partB()
// 	}
// }

// func partA() {
// 	scan := parseInput(*inputFn, 0)
// 	risk := scan.Risk()
// 	fmt.Println(risk)
// }

// func partB() {
// 	scan := parseInput(*inputFn, *buffer)
// 	fmt.Println(scan.path(point{0, 0}, scan.target))
// }

// func parseInput(fn string, buffer int) scan {
// 	input, err := ioutil.ReadFile(fn)
// 	if err != nil {
// 		log.Fatal("Error reading input file: ", err)
// 	}

// 	lines := strings.Split(string(input), "\n")
// 	depth, _ := strconv.Atoi(strings.Fields(lines[0])[1])
// 	target := strings.Split(strings.Fields(lines[1])[1], ",")
// 	tx, _ := strconv.Atoi(target[0])
// 	ty, _ := strconv.Atoi(target[1])
// 	return newScan(depth, buffer, point{tx, ty})
// }

// type terrain int

// const (
// 	rocky  terrain = 0
// 	wet    terrain = 1
// 	narrow terrain = 2
// )

// func (t terrain) String() string {
// 	switch t {
// 	case rocky:
// 		return "."
// 	case wet:
// 		return "="
// 	case narrow:
// 		return "|"
// 	}
// 	return "!"
// }

// type gear int

// const (
// 	torch    gear = 0
// 	climbing gear = 1
// 	neither  gear = 2
// )

// func (g gear) String() string {
// 	switch g {
// 	case torch:
// 		return "torch"
// 	case climbing:
// 		return "climbing gear"
// 	case neither:
// 		return "no gear"
// 	}
// 	return "wat even!!!"
// }

// const changeGearCost int = 7
// const moveCost int = 1

// type point struct {
// 	x int
// 	y int
// }

// func (p point) distance(t point) int {
// 	return abs(p.x-t.x) + abs(p.y-t.y)
// }

// func (p point) String() string {
// 	return fmt.Sprintf("(%d, %d)", p.x, p.y)
// }

// type move struct {
// 	p point
// 	g gear
// }

// type path []move

// func (p path) cost(s scan, g gear) int {
// 	res := 0
// 	steps := 0
// 	s.current = p[0].p
// 	if *trace {
// 		fmt.Println("Starting condition")
// 		// fmt.Println(s)
// 	}
// 	for _, to := range p[1:] {
// 		if s.current == to.p {
// 			res += changeGearCost
// 		} else {
// 			res += moveCost
// 		}
// 		steps++
// 		if *trace {
// 			fmt.Printf("Moving from %s to %s. ", s.current, to.p)
// 			fmt.Printf("Cost is %d after %d steps.  %s is now equipped.\n", res, steps, to.g)
// 			// fmt.Println(s)
// 		}
// 		g = to.g
// 		s.current = to.p
// 	}
// 	if *trace {
// 		fmt.Println("Reached target.")
// 	}
// 	if g != torch {
// 		if *trace {
// 			fmt.Println("Switching to torch.")
// 		}
// 		res += changeGearCost
// 	}
// 	if *trace {
// 		fmt.Printf("Total cost is %d after %d steps.\n", res, steps)
// 	}
// 	return res
// }

// type item struct {
// 	value    move
// 	priority int
// 	index    int
// }
// type priorityQueue []*item

// func (pq priorityQueue) Len() int { return len(pq) }
// func (pq priorityQueue) Less(i, j int) bool {
// 	return pq[i].priority < pq[j].priority
// }

// func (pq priorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// 	pq[i].index = i
// 	pq[j].index = j
// }

// func (pq *priorityQueue) Push(x interface{}) {
// 	n := len(*pq)
// 	i := x.(*item)
// 	i.index = n
// 	*pq = append(*pq, i)
// }

// func (pq *priorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	i := old[n-1]
// 	i.index = -1
// 	*pq = old[0 : n-1]
// 	return i
// }

// type scan struct {
// 	depth   int
// 	target  point
// 	current point
// 	regions map[point]int
// }

// func (s scan) String() string {
// 	var b strings.Builder
// 	maxX, maxY := 0, 0
// 	for k := range s.regions {
// 		if k.x > maxX {
// 			maxX = k.x
// 		}
// 		if k.y > maxY {
// 			maxY = k.y
// 		}
// 	}
// 	for y := 0; y <= maxY; y++ {
// 		for x := 0; x <= maxX; x++ {
// 			cur := point{x, y}
// 			if cur == s.current {
// 				b.WriteString("X")
// 			} else if cur == s.target {
// 				b.WriteString("T")
// 			} else {
// 				b.WriteString(s.pointType(cur).String())
// 			}
// 		}
// 		b.WriteString("\n")
// 	}
// 	return b.String()
// }

// func newScan(depth, buffer int, target point) scan {
// 	s := scan{depth: depth, target: target}
// 	width := s.target.x + buffer
// 	height := s.target.y + buffer
// 	s.regions = map[point]int{}
// 	s.regions[point{0, 0}] = 0
// 	s.regions[s.target] = 0
// 	for y := 0; y <= height; y++ {
// 		for x := 0; x <= width; x++ {
// 			idx := point{x, y}
// 			s.regions[idx] = s.erosionIndex(idx)
// 		}
// 	}
// 	return s
// }

// func (s scan) Risk() int {
// 	res := 0
// 	for y := 0; y <= s.target.y; y++ {
// 		for x := 0; x <= s.target.x; x++ {
// 			risk := s.erosionLevel(s.regions[point{x, y}]) % 3
// 			res += risk
// 		}
// 	}
// 	return res
// }

// func (s scan) pointType(p point) terrain {
// 	idx := s.erosionIndex(p)
// 	lvl := s.erosionLevel(idx)
// 	risk := lvl % 3
// 	return terrain(risk)
// }

// func (s scan) erosionIndex(p point) int {
// 	res := 0
// 	if _, ok := s.regions[p]; ok {
// 		res = s.regions[p]
// 	} else if p.y == 0 {
// 		res = p.x * 16807
// 	} else if p.x == 0 {
// 		res = p.y * 48271
// 	} else {
// 		res = s.erosionLevel(s.regions[point{x: p.x - 1, y: p.y}]) * s.erosionLevel(s.regions[point{x: p.x, y: p.y - 1}])
// 	}
// 	return res
// }

// func (s scan) erosionLevel(index int) int {
// 	return (index + s.depth) % 20183
// }

// func (s scan) path(p, t point) int {
// 	frontier := make(priorityQueue, 1)
// 	initialMove := move{p, torch}
// 	first := item{value: initialMove, priority: 0, index: 0}
// 	frontier[0] = &first
// 	heap.Init(&frontier)
// 	cameFrom := map[move]move{}
// 	costSoFar := map[move]int{}

// 	cameFrom[initialMove] = initialMove
// 	costSoFar[initialMove] = 0

// 	for frontier.Len() > 0 {
// 		i := heap.Pop(&frontier).(*item)
// 		cur := i.value

// 		if cur.p == t {
// 			curPath := path{cur}
// 			c := 0
// 			for n := cameFrom[cur]; n.p != p; n = cameFrom[n] {
// 				c++
// 				curPath = append(path{n}, curPath...)
// 			}
// 			curPath = append(path{move{p: p, g: torch}}, curPath...)
// 			cost := costSoFar[cur]
// 			if cur.g != torch {
// 				cost += changeGearCost
// 			}
// 			if *trace {
// 				fmt.Println("Calculated cost from reversing and walking path: %d", curPath.cost(s, torch))
// 			}
// 			return cost
// 		}

// 		for _, next := range s.neighbors(cur) {
// 			newCost := moveCost + costSoFar[cur]
// 			if next.p == t && cur.g != torch {
// 				newCost = 9999
// 			} else if cur.p == next.p && cur.g != next.g {
// 				newCost = changeGearCost + costSoFar[cur]
// 			}
// 			if cost, ok := costSoFar[next]; !ok || newCost < cost {
// 				costSoFar[next] = newCost
// 				priority := newCost + t.distance(next.p)
// 				nexti := item{value: next, priority: priority}
// 				for _, v := range frontier {
// 					if v.value == next {
// 						heap.Remove(&frontier, v.index)
// 						break
// 					}
// 				}
// 				heap.Push(&frontier, &nexti)
// 				cameFrom[next] = cur
// 			}
// 		}
// 	}
// 	return 0
// }

// func (s scan) cost(p, t point, g gear) (int, gear) {
// 	if _, ok := s.regions[p]; !ok {
// 		log.Fatalf("Unable to find source point %v in region map, moving to %v", p, t)
// 	}
// 	if _, ok := s.regions[t]; !ok {
// 		log.Fatalf("Unable to find target point %v in region map, moving from %v", t, p)
// 	}
// 	if t == s.target && g != torch {
// 		return moveCost + changeGearCost, torch
// 	}

// 	from := s.pointType(p)
// 	to := s.pointType(t)
// 	if from == to {
// 		return moveCost, g
// 	}

// 	if (from == rocky && to == wet) || (from == wet && to == rocky) {
// 		if g == climbing {
// 			return moveCost, g
// 		}
// 		return moveCost + changeGearCost, climbing
// 	}

// 	if (from == rocky && to == narrow) || (from == narrow && to == rocky) {
// 		if g == torch {
// 			return moveCost, g
// 		}
// 		return moveCost + changeGearCost, torch
// 	}

// 	// must be moving between wet and narrow
// 	if (from == wet && to == narrow) || (from == narrow && to == wet) {
// 		if g == neither {
// 			return moveCost, g
// 		}

// 		return moveCost + changeGearCost, neither
// 	}
// 	log.Fatalf("WTF!!!! From: %s, To: %s", from, to)
// 	return 999999, 99999
// }

// func (s scan) neighbors(m move) []move {
// 	res := []move{}
// 	for _, i := range []int{-1, 1} {
// 		if m.p.x+i >= 0 {
// 			np := point{m.p.x + i, m.p.y}
// 			if c, g := s.cost(m.p, np, m.g); c == moveCost {
// 				res = append(res, move{np, g})
// 			} else {
// 				res = append(res, move{m.p, g})
// 			}
// 		}
// 		if m.p.y+i >= 0 {
// 			np := point{m.p.x, m.p.y + i}
// 			if c, g := s.cost(m.p, np, m.g); c == moveCost {
// 				res = append(res, move{np, g})
// 			} else {
// 				res = append(res, move{m.p, g})
// 			}
// 		}
// 	}
// 	return res
// }

// func abs(i int) int {
// 	if i < 0 {
// 		return -1 * i
// 	}
// 	return i
// }
package main
