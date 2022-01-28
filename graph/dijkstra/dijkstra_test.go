package dijkstra

import (
	"Advent-of-Code/graph"
	"fmt"
	"reflect"
	"testing"
)

func TestPriorityQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		pq   PriorityQueue
		want int
	}{
		{
			name: "returns 0 for an empty queue",
			pq:   PriorityQueue{},
			want: 0,
		},
		{
			name: "returns the length of the given queue",
			pq:   PriorityQueue{&Path{}, &Path{}, &Path{}, &Path{}, &Path{}, &Path{}, &Path{}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Len(); got != tt.want {
				t.Errorf("PriorityQueue.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		pq   PriorityQueue
		args args
		want bool
	}{
		{
			name: "returns true if the value of the first item is less then the value of the second",
			pq: PriorityQueue{
				&Path{Value: 3},
				&Path{Value: 1},
				&Path{Value: 7},
				&Path{Value: 4},
				&Path{Value: 5},
				&Path{Value: 2},
				&Path{Value: 6},
			},
			args: args{i: 4, j: 6},
			want: true,
		},
		{
			name: "returns false if the value of the first item is not less then the value of the second",
			pq: PriorityQueue{
				&Path{Value: 3},
				&Path{Value: 1},
				&Path{Value: 7},
				&Path{Value: 4},
				&Path{Value: 5},
				&Path{Value: 2},
				&Path{Value: 6},
			},
			args: args{i: 4, j: 5},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("PriorityQueue.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		pq   PriorityQueue
		args args
		want PriorityQueue
	}{
		{
			name: "swaps the items at the given indices",
			pq: PriorityQueue{
				&Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
				&Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
				&Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
				&Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
				&Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
				&Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
				&Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
			},
			args: args{i: 2, j: 5},
			want: PriorityQueue{
				&Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
				&Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
				&Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
				&Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
				&Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
				&Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
				&Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := tt.pq
			pq.Swap(tt.args.i, tt.args.j)
			if !reflect.DeepEqual(pq, tt.want) {
				t.Errorf("PriorityQueue.Swap() = %v, want %v", pq, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	tests := []struct {
		name string
		pq   *PriorityQueue
		x    interface{}
		want *PriorityQueue
	}{
		{
			name: "pushes a Path to the PriorityQueue",
			pq: &PriorityQueue{
				&Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
				&Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
				&Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
				&Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
				&Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
				&Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
				&Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
			},
			x: &Path{Value: 9, Nodes: []graph.Co{{X: 9, Y: 9}}},
			want: &PriorityQueue{
				&Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
				&Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
				&Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
				&Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
				&Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
				&Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
				&Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
				&Path{Value: 9, Nodes: []graph.Co{{X: 9, Y: 9}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := tt.pq
			pq.Push(tt.x)
			if !reflect.DeepEqual(pq, tt.want) {
				t.Errorf("PriorityQueue.Push() = %v, want %v", pq, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	tests := []struct {
		name  string
		pq    *PriorityQueue
		want  interface{}
		want1 *PriorityQueue
	}{
		{
			name: "removes the last item from the PriorityQueue",
			pq: &PriorityQueue{
				&Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
				&Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
				&Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
				&Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
				&Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
				&Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
				&Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
				&Path{Value: 9, Nodes: []graph.Co{{X: 9, Y: 9}}},
			},
			want: &Path{Value: 9, Nodes: []graph.Co{{X: 9, Y: 9}}},
			want1: &PriorityQueue{
				&Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
				&Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
				&Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
				&Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
				&Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
				&Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
				&Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := tt.pq
			if got := pq.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriorityQueue.Pop() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(pq, tt.want1) {
				t.Errorf("PriorityQueue.Pop() = %v, want %v", pq, tt.want1)
			}
		})
	}
}

func TestNewGraph(t *testing.T) {
	type args struct {
		maxX int
		maxY int
	}
	tests := []struct {
		name string
		args args
		want *Graph
	}{
		{
			name: "returns a constructed graph from given maxX and maxY",
			args: args{maxX: 5, maxY: 87},
			want: &Graph{
				Grid:  make(map[graph.Co]int),
				Nodes: make(map[graph.Co][]Edge),
				MaxX:  5,
				MaxY:  87,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGraph(tt.args.maxX, tt.args.maxY); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {
	type args struct {
		origin      graph.Co
		destination graph.Co
		weight      int
	}
	tests := []struct {
		name string
		g    *Graph
		args args
		want *Graph
	}{
		{
			name: "adds an edge to the graph that is not in the grid or nodes",
			g: &Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 9,
					{X: 1, Y: 0}: 8,
					{X: 1, Y: 1}: 4,
				},
				Nodes: map[graph.Co][]Edge{
					{X: 0, Y: 0}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 9}},
					{X: 1, Y: 0}: {{Node: graph.Co{X: 1, Y: 1}, Weight: 4}},
					{X: 1, Y: 1}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 8}},
				},
				MaxX: 3,
				MaxY: 3,
			},
			args: args{
				origin:      graph.Co{X: 10, Y: 0},
				destination: graph.Co{X: 9, Y: 1},
				weight:      7,
			},
			want: &Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 9,
					{X: 1, Y: 0}: 8,
					{X: 1, Y: 1}: 4,
					{X: 9, Y: 1}: 7,
				},
				Nodes: map[graph.Co][]Edge{
					{X: 0, Y: 0}:  {{Node: graph.Co{X: 0, Y: 1}, Weight: 9}},
					{X: 1, Y: 0}:  {{Node: graph.Co{X: 1, Y: 1}, Weight: 4}},
					{X: 1, Y: 1}:  {{Node: graph.Co{X: 1, Y: 0}, Weight: 8}},
					{X: 10, Y: 0}: {{Node: graph.Co{X: 9, Y: 1}, Weight: 7}},
				},
				MaxX: 3,
				MaxY: 3,
			},
		},
		{
			name: "adds an edge to the graph that is not in the grid but is in nodes",
			g: &Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 9,
					{X: 1, Y: 0}: 8,
					{X: 1, Y: 1}: 4,
				},
				Nodes: map[graph.Co][]Edge{
					{X: 0, Y: 0}:  {{Node: graph.Co{X: 0, Y: 1}, Weight: 9}},
					{X: 1, Y: 0}:  {{Node: graph.Co{X: 1, Y: 1}, Weight: 4}},
					{X: 1, Y: 1}:  {{Node: graph.Co{X: 1, Y: 0}, Weight: 8}},
					{X: 10, Y: 0}: {{Node: graph.Co{X: 9, Y: 1}, Weight: 7}},
				},
				MaxX: 3,
				MaxY: 3,
			},
			args: args{
				origin:      graph.Co{X: 10, Y: 0},
				destination: graph.Co{X: 11, Y: 7},
				weight:      11,
			},
			want: &Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}:  1,
					{X: 0, Y: 1}:  9,
					{X: 1, Y: 0}:  8,
					{X: 1, Y: 1}:  4,
					{X: 11, Y: 7}: 11,
				},
				Nodes: map[graph.Co][]Edge{
					{X: 0, Y: 0}:  {{Node: graph.Co{X: 0, Y: 1}, Weight: 9}},
					{X: 1, Y: 0}:  {{Node: graph.Co{X: 1, Y: 1}, Weight: 4}},
					{X: 1, Y: 1}:  {{Node: graph.Co{X: 1, Y: 0}, Weight: 8}},
					{X: 10, Y: 0}: {{Node: graph.Co{X: 9, Y: 1}, Weight: 7}, {Node: graph.Co{X: 11, Y: 7}, Weight: 11}},
				},
				MaxX: 3,
				MaxY: 3,
			},
		},
		{
			name: "adds an edge to the graph that is in grid and in nodes",
			g: &Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}:  1,
					{X: 0, Y: 1}:  9,
					{X: 1, Y: 0}:  8,
					{X: 1, Y: 1}:  4,
					{X: 11, Y: 7}: 11,
				},
				Nodes: map[graph.Co][]Edge{
					{X: 0, Y: 0}:  {{Node: graph.Co{X: 0, Y: 1}, Weight: 9}},
					{X: 1, Y: 0}:  {{Node: graph.Co{X: 1, Y: 1}, Weight: 4}},
					{X: 1, Y: 1}:  {{Node: graph.Co{X: 1, Y: 0}, Weight: 8}},
					{X: 10, Y: 0}: {{Node: graph.Co{X: 9, Y: 1}, Weight: 7}},
				},
				MaxX: 3,
				MaxY: 3,
			},
			args: args{
				origin:      graph.Co{X: 10, Y: 0},
				destination: graph.Co{X: 11, Y: 7},
				weight:      11,
			},
			want: &Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}:  1,
					{X: 0, Y: 1}:  9,
					{X: 1, Y: 0}:  8,
					{X: 1, Y: 1}:  4,
					{X: 11, Y: 7}: 11,
				},
				Nodes: map[graph.Co][]Edge{
					{X: 0, Y: 0}:  {{Node: graph.Co{X: 0, Y: 1}, Weight: 9}},
					{X: 1, Y: 0}:  {{Node: graph.Co{X: 1, Y: 1}, Weight: 4}},
					{X: 1, Y: 1}:  {{Node: graph.Co{X: 1, Y: 0}, Weight: 8}},
					{X: 10, Y: 0}: {{Node: graph.Co{X: 9, Y: 1}, Weight: 7}, {Node: graph.Co{X: 11, Y: 7}, Weight: 11}},
				},
				MaxX: 3,
				MaxY: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.g
			g.AddEdge(tt.args.origin, tt.args.destination, tt.args.weight)
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("Graph.AddEdge() = %v, want %v", g, tt.want)
			}
		})
	}
}

func TestGraph_GetEdges(t *testing.T) {
	tests := []struct {
		name string
		g    *Graph
		node graph.Co
		want []Edge
	}{
		{
			name: "returns the edges in the graph for the given node",
			g: &Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}:  1,
					{X: 0, Y: 1}:  9,
					{X: 1, Y: 0}:  8,
					{X: 1, Y: 1}:  4,
					{X: 11, Y: 7}: 11,
				},
				Nodes: map[graph.Co][]Edge{
					{X: 0, Y: 0}:  {{Node: graph.Co{X: 0, Y: 1}, Weight: 9}},
					{X: 1, Y: 0}:  {{Node: graph.Co{X: 1, Y: 1}, Weight: 4}},
					{X: 1, Y: 1}:  {{Node: graph.Co{X: 1, Y: 0}, Weight: 8}},
					{X: 10, Y: 0}: {{Node: graph.Co{X: 9, Y: 1}, Weight: 7}, {Node: graph.Co{X: 11, Y: 7}, Weight: 11}},
				},
				MaxX: 3,
				MaxY: 3,
			},
			node: graph.Co{X: 1, Y: 1},
			want: []Edge{{Node: graph.Co{X: 1, Y: 0}, Weight: 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.g
			if got := g.GetEdges(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.GetEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_ExtendGrid(t *testing.T) {
	tests := []struct {
		name   string
		g      *Graph
		factor int
		want   *Graph
	}{
		{
			name: "does not affect a graph if given factor is 1",
			g: &Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 2,
					{X: 1, Y: 0}: 7,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 9,
				},
				Nodes: map[graph.Co][]Edge{
					{X: 0, Y: 0}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 7}},
					{X: 1, Y: 0}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 2}},
					{X: 0, Y: 1}: {{Node: graph.Co{X: 1, Y: 1}, Weight: 9}},
					{X: 1, Y: 1}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}},
				},
				MaxX: 1,
				MaxY: 1,
			},
			factor: 1,
			want: &Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 2,
					{X: 1, Y: 0}: 7,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 9,
				},
				Nodes: map[graph.Co][]Edge{
					{X: 0, Y: 0}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 7}},
					{X: 1, Y: 0}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 2}},
					{X: 0, Y: 1}: {{Node: graph.Co{X: 1, Y: 1}, Weight: 9}},
					{X: 1, Y: 1}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}},
				},
				MaxX: 1,
				MaxY: 1,
			},
		},
		{
			name: "extends a graph by a factor greater than 1",
			g: &Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 2,
					{X: 1, Y: 0}: 7,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 9,
				},
				Nodes: map[graph.Co][]Edge{
					{X: 0, Y: 0}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 7}},
					{X: 1, Y: 0}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 2}},
					{X: 0, Y: 1}: {{Node: graph.Co{X: 1, Y: 1}, Weight: 9}},
					{X: 1, Y: 1}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}},
				},
				MaxX: 1,
				MaxY: 1,
			},
			factor: 3,
			want: &Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 2,
					{X: 1, Y: 0}: 7,
					{X: 2, Y: 0}: 3,
					{X: 3, Y: 0}: 8,
					{X: 4, Y: 0}: 4,
					{X: 5, Y: 0}: 9,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 9,
					{X: 2, Y: 1}: 2,
					{X: 3, Y: 1}: 1,
					{X: 4, Y: 1}: 3,
					{X: 5, Y: 1}: 2,
					{X: 0, Y: 2}: 3,
					{X: 1, Y: 2}: 8,
					{X: 2, Y: 2}: 4,
					{X: 3, Y: 2}: 9,
					{X: 4, Y: 2}: 5,
					{X: 5, Y: 2}: 1,
					{X: 0, Y: 3}: 2,
					{X: 1, Y: 3}: 1,
					{X: 2, Y: 3}: 3,
					{X: 3, Y: 3}: 2,
					{X: 4, Y: 3}: 4,
					{X: 5, Y: 3}: 3,
					{X: 0, Y: 4}: 4,
					{X: 1, Y: 4}: 9,
					{X: 2, Y: 4}: 5,
					{X: 3, Y: 4}: 1,
					{X: 4, Y: 4}: 6,
					{X: 5, Y: 4}: 2,
					{X: 0, Y: 5}: 3,
					{X: 1, Y: 5}: 2,
					{X: 2, Y: 5}: 4,
					{X: 3, Y: 5}: 3,
					{X: 4, Y: 5}: 5,
					{X: 5, Y: 5}: 4,
				},
				Nodes: map[graph.Co][]Edge{
					{X: 0, Y: 0}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 7}},
					{X: 1, Y: 0}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 2}},
					{X: 0, Y: 1}: {{Node: graph.Co{X: 1, Y: 1}, Weight: 9}},
					{X: 1, Y: 1}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}},
				},
				MaxX: 5,
				MaxY: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.g
			g.ExtendGrid(tt.factor)
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("Graph.ExtendGrid() = %v, want %v", g, tt.want)
			}
		})
	}
}

var exampleGraph = Graph{
	Nodes: map[graph.Co][]Edge{
		{X: 0, Y: 0}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 1}, {Node: graph.Co{X: 0, Y: 1}, Weight: 99}},
		{X: 1, Y: 0}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 5}, {Node: graph.Co{X: 1, Y: 1}, Weight: 99}, {Node: graph.Co{X: 2, Y: 0}, Weight: 1}},
		{X: 2, Y: 0}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 1}, {Node: graph.Co{X: 2, Y: 1}, Weight: 99}, {Node: graph.Co{X: 3, Y: 0}, Weight: 1}},
		{X: 3, Y: 0}: {{Node: graph.Co{X: 2, Y: 0}, Weight: 1}, {Node: graph.Co{X: 3, Y: 1}, Weight: 99}, {Node: graph.Co{X: 4, Y: 0}, Weight: 1}},
		{X: 4, Y: 0}: {{Node: graph.Co{X: 3, Y: 0}, Weight: 1}, {Node: graph.Co{X: 4, Y: 1}, Weight: 1}},

		{X: 0, Y: 1}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 5}, {Node: graph.Co{X: 1, Y: 1}, Weight: 99}, {Node: graph.Co{X: 0, Y: 2}, Weight: 1}},
		{X: 1, Y: 1}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 1}, {Node: graph.Co{X: 2, Y: 1}, Weight: 99}, {Node: graph.Co{X: 1, Y: 2}, Weight: 1}, {Node: graph.Co{X: 0, Y: 1}, Weight: 99}},
		{X: 2, Y: 1}: {{Node: graph.Co{X: 2, Y: 0}, Weight: 1}, {Node: graph.Co{X: 3, Y: 1}, Weight: 99}, {Node: graph.Co{X: 2, Y: 2}, Weight: 1}, {Node: graph.Co{X: 1, Y: 1}, Weight: 99}},
		{X: 3, Y: 1}: {{Node: graph.Co{X: 3, Y: 0}, Weight: 1}, {Node: graph.Co{X: 4, Y: 1}, Weight: 1}, {Node: graph.Co{X: 3, Y: 2}, Weight: 1}, {Node: graph.Co{X: 2, Y: 1}, Weight: 99}},
		{X: 4, Y: 1}: {{Node: graph.Co{X: 4, Y: 0}, Weight: 1}, {Node: graph.Co{X: 4, Y: 2}, Weight: 1}, {Node: graph.Co{X: 3, Y: 1}, Weight: 99}},

		{X: 0, Y: 2}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 99}, {Node: graph.Co{X: 1, Y: 2}, Weight: 1}, {Node: graph.Co{X: 0, Y: 3}, Weight: 1}},
		{X: 1, Y: 2}: {{Node: graph.Co{X: 1, Y: 1}, Weight: 99}, {Node: graph.Co{X: 2, Y: 2}, Weight: 1}, {Node: graph.Co{X: 1, Y: 3}, Weight: 99}, {Node: graph.Co{X: 0, Y: 2}, Weight: 1}},
		{X: 2, Y: 2}: {{Node: graph.Co{X: 2, Y: 1}, Weight: 99}, {Node: graph.Co{X: 3, Y: 2}, Weight: 1}, {Node: graph.Co{X: 2, Y: 3}, Weight: 99}, {Node: graph.Co{X: 1, Y: 2}, Weight: 1}},
		{X: 3, Y: 2}: {{Node: graph.Co{X: 3, Y: 1}, Weight: 99}, {Node: graph.Co{X: 4, Y: 2}, Weight: 1}, {Node: graph.Co{X: 3, Y: 3}, Weight: 99}, {Node: graph.Co{X: 2, Y: 2}, Weight: 1}},
		{X: 4, Y: 2}: {{Node: graph.Co{X: 4, Y: 1}, Weight: 1}, {Node: graph.Co{X: 4, Y: 3}, Weight: 99}, {Node: graph.Co{X: 3, Y: 2}, Weight: 1}},

		{X: 0, Y: 3}: {{Node: graph.Co{X: 0, Y: 2}, Weight: 1}, {Node: graph.Co{X: 1, Y: 3}, Weight: 99}, {Node: graph.Co{X: 0, Y: 4}, Weight: 1}},
		{X: 1, Y: 3}: {{Node: graph.Co{X: 1, Y: 2}, Weight: 1}, {Node: graph.Co{X: 2, Y: 3}, Weight: 99}, {Node: graph.Co{X: 1, Y: 4}, Weight: 99}, {Node: graph.Co{X: 0, Y: 3}, Weight: 1}},
		{X: 2, Y: 3}: {{Node: graph.Co{X: 2, Y: 2}, Weight: 1}, {Node: graph.Co{X: 3, Y: 3}, Weight: 99}, {Node: graph.Co{X: 2, Y: 4}, Weight: 1}, {Node: graph.Co{X: 1, Y: 3}, Weight: 99}},
		{X: 3, Y: 3}: {{Node: graph.Co{X: 3, Y: 2}, Weight: 1}, {Node: graph.Co{X: 4, Y: 3}, Weight: 99}, {Node: graph.Co{X: 3, Y: 4}, Weight: 1}, {Node: graph.Co{X: 2, Y: 3}, Weight: 99}},
		{X: 4, Y: 3}: {{Node: graph.Co{X: 4, Y: 2}, Weight: 1}, {Node: graph.Co{X: 4, Y: 4}, Weight: 1}, {Node: graph.Co{X: 4, Y: 3}, Weight: 99}},

		{X: 0, Y: 4}: {{Node: graph.Co{X: 0, Y: 3}, Weight: 1}, {Node: graph.Co{X: 1, Y: 4}, Weight: 99}, {Node: graph.Co{X: 0, Y: 5}, Weight: 1}},
		{X: 1, Y: 4}: {{Node: graph.Co{X: 1, Y: 3}, Weight: 99}, {Node: graph.Co{X: 2, Y: 4}, Weight: 1}, {Node: graph.Co{X: 1, Y: 5}, Weight: 1}, {Node: graph.Co{X: 0, Y: 4}, Weight: 1}},
		{X: 2, Y: 4}: {{Node: graph.Co{X: 2, Y: 3}, Weight: 99}, {Node: graph.Co{X: 3, Y: 4}, Weight: 1}, {Node: graph.Co{X: 2, Y: 5}, Weight: 1}, {Node: graph.Co{X: 1, Y: 4}, Weight: 99}},
		{X: 3, Y: 4}: {{Node: graph.Co{X: 3, Y: 3}, Weight: 99}, {Node: graph.Co{X: 4, Y: 4}, Weight: 1}, {Node: graph.Co{X: 3, Y: 5}, Weight: 99}, {Node: graph.Co{X: 2, Y: 4}, Weight: 1}},
		{X: 4, Y: 4}: {{Node: graph.Co{X: 4, Y: 3}, Weight: 99}, {Node: graph.Co{X: 4, Y: 5}, Weight: 1}, {Node: graph.Co{X: 4, Y: 4}, Weight: 1}},

		{X: 0, Y: 5}: {{Node: graph.Co{X: 0, Y: 4}, Weight: 1}, {Node: graph.Co{X: 1, Y: 5}, Weight: 1}},
		{X: 1, Y: 5}: {{Node: graph.Co{X: 0, Y: 5}, Weight: 1}, {Node: graph.Co{X: 1, Y: 4}, Weight: 99}, {Node: graph.Co{X: 2, Y: 5}, Weight: 1}},
		{X: 2, Y: 5}: {{Node: graph.Co{X: 1, Y: 5}, Weight: 1}, {Node: graph.Co{X: 2, Y: 4}, Weight: 1}, {Node: graph.Co{X: 3, Y: 5}, Weight: 99}},
		{X: 3, Y: 5}: {{Node: graph.Co{X: 2, Y: 5}, Weight: 1}, {Node: graph.Co{X: 3, Y: 4}, Weight: 1}, {Node: graph.Co{X: 4, Y: 5}, Weight: 1}},
		{X: 4, Y: 5}: {{Node: graph.Co{X: 3, Y: 5}, Weight: 99}, {Node: graph.Co{X: 4, Y: 4}, Weight: 1}},
	},
	Grid: map[graph.Co]int{
		{X: 0, Y: 0}: 5,
		{X: 1, Y: 0}: 1,
		{X: 2, Y: 0}: 1,
		{X: 3, Y: 0}: 1,
		{X: 4, Y: 0}: 1,
		{X: 0, Y: 1}: 99,
		{X: 1, Y: 1}: 99,
		{X: 2, Y: 1}: 99,
		{X: 3, Y: 1}: 99,
		{X: 4, Y: 1}: 1,
		{X: 0, Y: 2}: 1,
		{X: 1, Y: 2}: 1,
		{X: 2, Y: 2}: 1,
		{X: 3, Y: 2}: 1,
		{X: 4, Y: 2}: 1,
		{X: 0, Y: 3}: 1,
		{X: 1, Y: 3}: 99,
		{X: 2, Y: 3}: 99,
		{X: 3, Y: 3}: 99,
		{X: 4, Y: 3}: 99,
		{X: 0, Y: 4}: 1,
		{X: 1, Y: 4}: 99,
		{X: 2, Y: 4}: 1,
		{X: 3, Y: 4}: 1,
		{X: 4, Y: 4}: 1,
		{X: 0, Y: 5}: 1,
		{X: 1, Y: 5}: 1,
		{X: 2, Y: 5}: 1,
		{X: 3, Y: 5}: 99,
		{X: 4, Y: 5}: 1,
	},
	MaxX: 4,
	MaxY: 5,
}

func TestGraph_GetPath(t *testing.T) {
	type args struct {
		origin      graph.Co
		destination graph.Co
	}
	tests := []struct {
		name    string
		args    args
		want    *Path
		wantErr bool
	}{
		{
			name: "returns an error if a path can't be found",
			args: args{
				origin:      graph.Co{X: 0, Y: 0},
				destination: graph.Co{X: 100, Y: 100},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "finds minimal path through graph",
			args: args{
				origin:      graph.Co{X: 0, Y: 0},
				destination: graph.Co{X: 4, Y: 5},
			},
			want: &Path{
				Value: 19,
				Nodes: []graph.Co{
					{X: 0, Y: 0},
					{X: 1, Y: 0},
					{X: 2, Y: 0},
					{X: 3, Y: 0},
					{X: 4, Y: 0},
					{X: 4, Y: 1},
					{X: 4, Y: 2},
					{X: 3, Y: 2},
					{X: 2, Y: 2},
					{X: 1, Y: 2},
					{X: 0, Y: 2},
					{X: 0, Y: 3},
					{X: 0, Y: 4},
					{X: 0, Y: 5},
					{X: 1, Y: 5},
					{X: 2, Y: 5},
					{X: 2, Y: 4},
					{X: 3, Y: 4},
					{X: 4, Y: 4},
					{X: 4, Y: 5},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := exampleGraph.GetPath(tt.args.origin, tt.args.destination)
			for co, edges := range exampleGraph.Nodes {
				for _, e := range edges {
					if exampleGraph.Grid[e.Node] != e.Weight {
						fmt.Println(co, e)
					}
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Graph.GetPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
