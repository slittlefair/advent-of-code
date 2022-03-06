package dijkstra

import (
	"Advent-of-Code/graph"
	"container/heap"
	"fmt"
)

type Path struct {
	Value int
	Nodes []graph.Co
}

type PriorityQueue []*Path

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Value < pq[j].Value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Path)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[:n-1]
	return node
}

// PrintGrid is used for debugging
func (g Graph) PrintGrid(riskMode bool) {
	for y := 0; y <= g.MaxY; y++ {
		for x := 0; x <= g.MaxX; x++ {
			if riskMode {
				fmt.Print(g.Grid[graph.Co{X: x, Y: y}])
			} else {
				if _, ok := g.Grid[graph.Co{X: x, Y: y}]; ok {
					fmt.Printf(".")
				} else {
					fmt.Printf("#")
				}
			}
		}
		fmt.Println()
	}
}

type Edge struct {
	Node   graph.Co
	Weight int
}

type Graph struct {
	Grid       map[graph.Co]int
	Nodes      map[graph.Co][]Edge
	MaxX, MaxY int
}

func NewGraph(maxX, maxY int) *Graph {
	return &Graph{
		Grid:  make(map[graph.Co]int),
		Nodes: make(map[graph.Co][]Edge),
		MaxX:  maxX,
		MaxY:  maxY,
	}
}

func (g *Graph) AddEdge(origin, destination graph.Co, weight int) {
	g.Nodes[origin] = append(g.Nodes[origin], Edge{Node: destination, Weight: weight})
	if _, ok := g.Grid[destination]; !ok {
		g.Grid[destination] = weight
	}
}

func (g *Graph) GetEdges(node graph.Co) []Edge {
	return g.Nodes[node]
}

func (g *Graph) ExtendGrid(factor int) {
	newGrid := map[graph.Co]int{}
	for y := 0; y < factor; y++ {
		for x := 0; x < factor; x++ {
			for co, val := range g.Grid {
				risk := val + x + y
				for risk > 9 {
					risk -= 9
				}
				newGrid[graph.Co{X: co.X + ((g.MaxX + 1) * x), Y: co.Y + ((g.MaxY + 1) * y)}] = risk
			}
		}
	}
	g.Grid = newGrid
	g.MaxX = factor*g.MaxX + factor - 1
	g.MaxY = factor*g.MaxY + factor - 1
}

func (g Graph) GetPath(origin, destination graph.Co) (*Path, error) {
	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &Path{
		Value: 0,
		Nodes: []graph.Co{origin},
	})
	visited := map[graph.Co]struct{}{}
	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*Path)
		node := p.Nodes[len(p.Nodes)-1]
		if _, ok := visited[node]; ok {
			continue
		}
		if node == destination {
			return p, nil
		}
		for _, n := range g.GetEdges(node) {
			if _, ok := visited[n.Node]; !ok {
				heap.Push(&pq, &Path{
					Value: p.Value + n.Weight,
					Nodes: append([]graph.Co{}, append(p.Nodes, n.Node)...),
				})
			}
		}
		visited[node] = struct{}{}
	}
	return nil, fmt.Errorf("no path found")
}
