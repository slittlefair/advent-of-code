package dijkstra_test

import (
	"Advent-of-Code/graph"
	"Advent-of-Code/graph/dijkstra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		pq   dijkstra.PriorityQueue
		want int
	}{
		{
			name: "returns 0 for an empty queue",
			pq:   dijkstra.PriorityQueue{},
			want: 0,
		},
		{
			name: "returns the length of the given queue",
			pq:   dijkstra.PriorityQueue{&dijkstra.Path{}, &dijkstra.Path{}, &dijkstra.Path{}, &dijkstra.Path{}, &dijkstra.Path{}, &dijkstra.Path{}, &dijkstra.Path{}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.pq.Len()
			assert.Equal(t, tt.want, got)
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
		pq   dijkstra.PriorityQueue
		args args
		want bool
	}{
		{
			name: "returns true if the value of the first item is less then the value of the second",
			pq: dijkstra.PriorityQueue{
				&dijkstra.Path{Value: 3},
				&dijkstra.Path{Value: 1},
				&dijkstra.Path{Value: 7},
				&dijkstra.Path{Value: 4},
				&dijkstra.Path{Value: 5},
				&dijkstra.Path{Value: 2},
				&dijkstra.Path{Value: 6},
			},
			args: args{i: 4, j: 6},
			want: true,
		},
		{
			name: "returns false if the value of the first item is not less then the value of the second",
			pq: dijkstra.PriorityQueue{
				&dijkstra.Path{Value: 3},
				&dijkstra.Path{Value: 1},
				&dijkstra.Path{Value: 7},
				&dijkstra.Path{Value: 4},
				&dijkstra.Path{Value: 5},
				&dijkstra.Path{Value: 2},
				&dijkstra.Path{Value: 6},
			},
			args: args{i: 4, j: 5},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.pq.Less(tt.args.i, tt.args.j)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPriorityQueue_Swap(t *testing.T) {
	t.Run("swaps the items at the given indices", func(t *testing.T) {
		pq := dijkstra.PriorityQueue{
			&dijkstra.Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
			&dijkstra.Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
			&dijkstra.Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
			&dijkstra.Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
			&dijkstra.Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
			&dijkstra.Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
			&dijkstra.Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
		}
		want := dijkstra.PriorityQueue{
			&dijkstra.Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
			&dijkstra.Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
			&dijkstra.Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
			&dijkstra.Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
			&dijkstra.Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
			&dijkstra.Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
			&dijkstra.Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
		}
		pq.Swap(2, 5)
		assert.Equal(t, want, pq)
	})
}

func TestPriorityQueue_Push(t *testing.T) {
	t.Run("pushes a Path to the PriorityQueue", func(t *testing.T) {
		pq := dijkstra.PriorityQueue{
			&dijkstra.Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
			&dijkstra.Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
			&dijkstra.Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
			&dijkstra.Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
			&dijkstra.Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
			&dijkstra.Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
			&dijkstra.Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
		}
		want := dijkstra.PriorityQueue{
			&dijkstra.Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
			&dijkstra.Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
			&dijkstra.Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
			&dijkstra.Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
			&dijkstra.Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
			&dijkstra.Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
			&dijkstra.Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
			&dijkstra.Path{Value: 9, Nodes: []graph.Co{{X: 9, Y: 9}}},
		}
		pq.Push(&dijkstra.Path{Value: 9, Nodes: []graph.Co{{X: 9, Y: 9}}})
		assert.Equal(t, want, pq)
	})
}

func TestPriorityQueue_Pop(t *testing.T) {
	t.Run("removes the last item from the PriorityQueue", func(t *testing.T) {
		pq := &dijkstra.PriorityQueue{
			&dijkstra.Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
			&dijkstra.Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
			&dijkstra.Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
			&dijkstra.Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
			&dijkstra.Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
			&dijkstra.Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
			&dijkstra.Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
			&dijkstra.Path{Value: 9, Nodes: []graph.Co{{X: 9, Y: 9}}},
		}
		want := &dijkstra.Path{Value: 9, Nodes: []graph.Co{{X: 9, Y: 9}}}
		want1 := &dijkstra.PriorityQueue{
			&dijkstra.Path{Value: 3, Nodes: []graph.Co{{X: 0, Y: 1}}},
			&dijkstra.Path{Value: 1, Nodes: []graph.Co{{X: 1, Y: 1}}},
			&dijkstra.Path{Value: 2, Nodes: []graph.Co{{X: 7, Y: 1}, {X: 6, Y: 4}}},
			&dijkstra.Path{Value: 4, Nodes: []graph.Co{{X: 0, Y: 4}, {X: 10, Y: 9}}},
			&dijkstra.Path{Value: 5, Nodes: []graph.Co{{X: 0, Y: 5}}},
			&dijkstra.Path{Value: 7, Nodes: []graph.Co{{X: 2, Y: 1}}},
			&dijkstra.Path{Value: 6, Nodes: []graph.Co{{X: 8, Y: 1}}},
		}
		got := pq.Pop()
		assert.Equal(t, want, got)
		assert.Equal(t, want1, pq)
	})
}

func TestNewGraph(t *testing.T) {
	t.Run("returns a constructed graph from given maxX and maxY", func(t *testing.T) {
		got := dijkstra.NewGraph(5, 87)
		want := &dijkstra.Graph{
			Grid:  make(map[graph.Co]int),
			Nodes: make(map[graph.Co][]dijkstra.Edge),
			MaxX:  5,
			MaxY:  87,
		}
		assert.Equal(t, want, got)
	})
}

func TestGraph_AddEdge(t *testing.T) {
	type args struct {
		origin      graph.Co
		destination graph.Co
		weight      int
	}
	tests := []struct {
		name string
		g    *dijkstra.Graph
		args args
		want *dijkstra.Graph
	}{
		{
			name: "adds an edge to the graph that is not in the grid or nodes",
			g: &dijkstra.Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 9,
					{X: 1, Y: 0}: 8,
					{X: 1, Y: 1}: 4,
				},
				Nodes: map[graph.Co][]dijkstra.Edge{
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
			want: &dijkstra.Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 9,
					{X: 1, Y: 0}: 8,
					{X: 1, Y: 1}: 4,
					{X: 9, Y: 1}: 7,
				},
				Nodes: map[graph.Co][]dijkstra.Edge{
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
			g: &dijkstra.Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 9,
					{X: 1, Y: 0}: 8,
					{X: 1, Y: 1}: 4,
				},
				Nodes: map[graph.Co][]dijkstra.Edge{
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
			want: &dijkstra.Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}:  1,
					{X: 0, Y: 1}:  9,
					{X: 1, Y: 0}:  8,
					{X: 1, Y: 1}:  4,
					{X: 11, Y: 7}: 11,
				},
				Nodes: map[graph.Co][]dijkstra.Edge{
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
			g: &dijkstra.Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}:  1,
					{X: 0, Y: 1}:  9,
					{X: 1, Y: 0}:  8,
					{X: 1, Y: 1}:  4,
					{X: 11, Y: 7}: 11,
				},
				Nodes: map[graph.Co][]dijkstra.Edge{
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
			want: &dijkstra.Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}:  1,
					{X: 0, Y: 1}:  9,
					{X: 1, Y: 0}:  8,
					{X: 1, Y: 1}:  4,
					{X: 11, Y: 7}: 11,
				},
				Nodes: map[graph.Co][]dijkstra.Edge{
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
			assert.Equal(t, tt.want, g)
		})
	}
}

func TestGraph_GetEdges(t *testing.T) {
	t.Run("returns the edges in the graph for the given node", func(t *testing.T) {
		g := &dijkstra.Graph{
			Grid: map[graph.Co]int{
				{X: 0, Y: 0}:  1,
				{X: 0, Y: 1}:  9,
				{X: 1, Y: 0}:  8,
				{X: 1, Y: 1}:  4,
				{X: 11, Y: 7}: 11,
			},
			Nodes: map[graph.Co][]dijkstra.Edge{
				{X: 0, Y: 0}:  {{Node: graph.Co{X: 0, Y: 1}, Weight: 9}},
				{X: 1, Y: 0}:  {{Node: graph.Co{X: 1, Y: 1}, Weight: 4}},
				{X: 1, Y: 1}:  {{Node: graph.Co{X: 1, Y: 0}, Weight: 8}},
				{X: 10, Y: 0}: {{Node: graph.Co{X: 9, Y: 1}, Weight: 7}, {Node: graph.Co{X: 11, Y: 7}, Weight: 11}},
			},
			MaxX: 3,
			MaxY: 3,
		}
		got := g.GetEdges(graph.Co{X: 1, Y: 1})
		assert.Equal(t, []dijkstra.Edge{{Node: graph.Co{X: 1, Y: 0}, Weight: 8}}, got)
	})
}

func TestGraph_ExtendGrid(t *testing.T) {
	tests := []struct {
		name   string
		g      *dijkstra.Graph
		factor int
		want   *dijkstra.Graph
	}{
		{
			name: "does not affect a graph if given factor is 1",
			g: &dijkstra.Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 2,
					{X: 1, Y: 0}: 7,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 9,
				},
				Nodes: map[graph.Co][]dijkstra.Edge{
					{X: 0, Y: 0}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 7}},
					{X: 1, Y: 0}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 2}},
					{X: 0, Y: 1}: {{Node: graph.Co{X: 1, Y: 1}, Weight: 9}},
					{X: 1, Y: 1}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}},
				},
				MaxX: 1,
				MaxY: 1,
			},
			factor: 1,
			want: &dijkstra.Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 2,
					{X: 1, Y: 0}: 7,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 9,
				},
				Nodes: map[graph.Co][]dijkstra.Edge{
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
			g: &dijkstra.Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 2,
					{X: 1, Y: 0}: 7,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 9,
				},
				Nodes: map[graph.Co][]dijkstra.Edge{
					{X: 0, Y: 0}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 7}},
					{X: 1, Y: 0}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 2}},
					{X: 0, Y: 1}: {{Node: graph.Co{X: 1, Y: 1}, Weight: 9}},
					{X: 1, Y: 1}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}},
				},
				MaxX: 1,
				MaxY: 1,
			},
			factor: 3,
			want: &dijkstra.Graph{
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
				Nodes: map[graph.Co][]dijkstra.Edge{
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
			assert.Equal(t, tt.want, g)
		})
	}
}

var exampleGraph = dijkstra.Graph{
	Nodes: map[graph.Co][]dijkstra.Edge{
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
		name               string
		args               args
		want               *dijkstra.Path
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if a path can't be found",
			args: args{
				origin:      graph.Co{X: 0, Y: 0},
				destination: graph.Co{X: 100, Y: 100},
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "finds minimal path through graph",
			args: args{
				origin:      graph.Co{X: 0, Y: 0},
				destination: graph.Co{X: 4, Y: 5},
			},
			want: &dijkstra.Path{
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
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := exampleGraph.GetPath(tt.args.origin, tt.args.destination)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
