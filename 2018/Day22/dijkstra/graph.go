package dijkstra

import (
	helpers "Advent-of-Code"
	"fmt"
	"strconv"
)

type node struct {
	key  string
	cost int
}

// Graph is a rappresentation of how the points in our graph are connected
// between each other
type Graph map[string]map[string]int

// TypeToEquipment maps a region type to the equipment that can be used in that region
var TypeToEquipment = map[int][]string{
	0: []string{"torch", "climbing"},
	1: []string{"climbing", "no gear"},
	2: []string{"torch", "no gear"},
}

// CalculateTimes calculate the time it takes to travel to a region from an adjacent region
func (g *Graph) CalculateTimes(co string, coType int, adjacentCo map[helpers.Co]int, equip string) {
	for key, val := range adjacentCo {
		for _, eq := range TypeToEquipment[val] {
			adjCoString := fmt.Sprintf("%v,%v,%v", strconv.Itoa(key.X), strconv.Itoa(key.Y), eq)
			if eq == equip {
				(*g)[co][adjCoString] = 1
			} else {
				(*g)[co][adjCoString] = 8
			}
		}
	}

}

// Path finds the shortest path between start and target, also returning the
// total cost of the found path.
func (g Graph) Path(start, target string) (path []string, cost int, err error) {
	if len(g) == 0 {
		err = fmt.Errorf("cannot find path in empty map")
		return
	}

	// ensure start and target are part of the graph
	if _, ok := g[start]; !ok {
		err = fmt.Errorf("cannot find start %v in graph", start)
		return
	}
	if _, ok := g[target]; !ok {
		err = fmt.Errorf("cannot find target %v in graph", target)
		return
	}

	explored := make(map[string]bool)   // set of nodes we already explored
	frontier := NewQueue()              // queue of the nodes to explore
	previous := make(map[string]string) // previously visited node

	// add starting point to the frontier as it'll be the first node visited
	frontier.Set(start, 0)

	// run until we visited every node in the frontier
	for !frontier.IsEmpty() {
		// get the node in the frontier with the lowest cost (or priority)
		aKey, aPriority := frontier.Next()
		n := node{aKey, aPriority}

		// when the node with the lowest cost in the frontier is target, we can
		// compute the cost and path and exit the loop
		if n.key == target {
			cost = n.cost

			nKey := n.key
			for nKey != start {
				path = append(path, nKey)
				nKey = previous[nKey]
			}

			break
		}

		// add the current node to the explored set
		explored[n.key] = true

		// loop all the neighboring nodes
		for nKey, nCost := range g[n.key] {
			// skip alreadt-explored nodes
			if explored[nKey] {
				continue
			}

			// if the node is not yet in the frontier add it with the cost
			if _, ok := frontier.Get(nKey); !ok {
				previous[nKey] = n.key
				frontier.Set(nKey, n.cost+nCost)
				continue
			}

			frontierCost, _ := frontier.Get(nKey)
			nodeCost := n.cost + nCost

			// only update the cost of this node in the frontier when
			// it's below what's currently set
			if nodeCost < frontierCost {
				previous[nKey] = n.key
				frontier.Set(nKey, nodeCost)
			}
		}
	}

	// add the origin at the end of the path
	path = append(path, start)

	// reverse the path because it was popilated
	// in reverse, form target to start
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return
}
