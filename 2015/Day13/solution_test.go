package main

import (
	"reflect"
	"testing"
)

func TestGraph_AddNode(t *testing.T) {
	tests := []struct {
		name  string
		nodes []string
		arg   string
		want  []string
	}{
		{
			name: "doesn't add a node if it's already in the graph",
			nodes: []string{
				"Australia",
				"Belgium",
				"Cyprus",
			},
			arg: "Belgium",
			want: []string{
				"Australia",
				"Belgium",
				"Cyprus",
			},
		},
		{
			name: "does add a node if it's not already in the graph",
			nodes: []string{
				"Australia",
				"Belgium",
				"Cyprus",
			},
			arg: "Denmark",
			want: []string{
				"Australia",
				"Belgium",
				"Cyprus",
				"Denmark",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Nodes: tt.nodes,
			}
			g.AddNode(tt.arg)
			if !reflect.DeepEqual(g.Nodes, tt.want) {
				t.Errorf("Graph.AddNode() = %v, want %v", g.Nodes, tt.want)
			}
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {
	type fields struct {
		Edges []*Edge
		Nodes []string
	}
	type args struct {
		parent string
		child  string
		cost   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Graph
	}{
		{
			name: "adds an edge and two nodes to a graph",
			args: args{
				parent: "Australia",
				child:  "Belgium",
				cost:   11,
			},
			want: &Graph{
				Edges: []*Edge{
					{
						Parent: "Australia",
						Child:  "Belgium",
						Cost:   11,
					},
				},
				Nodes: []string{"Australia", "Belgium"},
			},
		},
		{
			name: "adds an edge and parent to a graph (child already exists)",
			args: args{
				parent: "Australia",
				child:  "Belgium",
				cost:   22,
			},
			fields: fields{
				Nodes: []string{"Belgium"},
			},
			want: &Graph{
				Edges: []*Edge{
					{
						Parent: "Australia",
						Child:  "Belgium",
						Cost:   22,
					},
				},
				Nodes: []string{"Belgium", "Australia"},
			},
		},
		{
			name: "adds an edge and child to a graph (parent already exists)",
			args: args{
				parent: "Australia",
				child:  "Belgium",
				cost:   33,
			},
			fields: fields{
				Nodes: []string{"Australia"},
			},
			want: &Graph{
				Edges: []*Edge{
					{
						Parent: "Australia",
						Child:  "Belgium",
						Cost:   33,
					},
				},
				Nodes: []string{"Australia", "Belgium"},
			},
		},
		{
			name: "changes the cost of an edge when the parent and child already exist",
			args: args{
				parent: "Belgium",
				child:  "Australia",
				cost:   12,
			},
			fields: fields{
				Edges: []*Edge{
					{
						Parent: "Australia",
						Child:  "Belgium",
						Cost:   44,
					},
				},
				Nodes: []string{"Australia", "Belgium"},
			},
			want: &Graph{
				Edges: []*Edge{
					{
						Parent: "Australia",
						Child:  "Belgium",
						Cost:   56,
					},
				},
				Nodes: []string{"Australia", "Belgium"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Edges: tt.fields.Edges,
				Nodes: tt.fields.Nodes,
			}
			g.AddEdge(tt.args.parent, tt.args.child, tt.args.cost)
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("Graph.AddEdge() = %v, want %v", g, tt.want)
			}
		})
	}
}

func TestGraph_AddMe(t *testing.T) {
	type fields struct {
		Edges []*Edge
		Nodes []string
		Paths [][]string
	}
	tests := []struct {
		name   string
		fields fields
		want   *Graph
	}{
		{
			name: `adds "Me" to a populated Graph`,
			fields: fields{
				Edges: []*Edge{
					{
						Parent: "Alf",
						Child:  "Bert",
						Cost:   5,
					},
				},
				Nodes: []string{"Alf", "Bert"},
				Paths: [][]string{
					{"Alf", "Bert"},
					{"Bert", "Alf"},
				},
			},
			want: &Graph{
				Edges: []*Edge{
					{
						Parent: "Alf",
						Child:  "Bert",
						Cost:   5,
					},
					{
						Parent: "Alf",
						Child:  "Me",
						Cost:   0,
					},
					{
						Parent: "Bert",
						Child:  "Me",
						Cost:   0,
					},
				},
				Nodes: []string{"Alf", "Bert", "Me"},
				Paths: [][]string{
					{"Alf", "Bert", "Me"},
					{"Bert", "Alf", "Me"},
					{"Me", "Bert", "Alf"},
					{"Bert", "Me", "Alf"},
					{"Me", "Alf", "Bert"},
					{"Alf", "Me", "Bert"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Edges: tt.fields.Edges,
				Nodes: tt.fields.Nodes,
				Paths: tt.fields.Paths,
			}
			g.AddMe()
			if !reflect.DeepEqual(g.Nodes, tt.want.Nodes) {
				t.Errorf("Graph.AddEdge().Nodes = %v, want %v", g.Nodes, tt.want.Nodes)
			}
			if !reflect.DeepEqual(g.Paths, tt.want.Paths) {
				t.Errorf("Graph.AddEdge().Paths = %v, want %v", g.Paths, tt.want.Paths)
			}
			for i, e := range g.Edges {
				if want := tt.want.Edges[i]; *want != *e {
					t.Errorf("Graph.AddEdge().Edge = %v, want %v", *e, *want)
				}
			}
		})
	}
}

func TestGraph_ParseInput(t *testing.T) {
	tests := []struct {
		name    string
		arg     []string
		want    *Graph
		wantErr bool
	}{
		{
			name: "returns an error if an input line cost can't be converted correctly",
			arg: []string{
				"Alice would gain 54 happiness units by sitting next to Bob.",
				"Alice would lose some happiness units by sitting next to Carol.",
				"Bob would gain 83 happiness units by sitting next to Alice.",
				"Bob would lose 7 happiness units by sitting next to Carol.",
				"Carol would lose 62 happiness units by sitting next to Alice.",
				"Carol would gain 60 happiness units by sitting next to Bob.",
			},
			want: &Graph{
				Edges: []*Edge{
					{
						Parent: "Alice",
						Child:  "Bob",
						Cost:   54,
					},
				},
				Nodes: []string{"Alice", "Bob"},
			},
			wantErr: true,
		},
		{
			name: "returns a populated graph from simple input",
			arg: []string{
				"Alice would gain 54 happiness units by sitting next to Bob.",
				"Alice would lose 79 happiness units by sitting next to Carol.",
				"Bob would gain 83 happiness units by sitting next to Alice.",
				"Bob would lose 7 happiness units by sitting next to Carol.",
				"Carol would lose 62 happiness units by sitting next to Alice.",
				"Carol would gain 60 happiness units by sitting next to Bob.",
			},
			want: &Graph{
				Edges: []*Edge{
					{
						Parent: "Alice",
						Child:  "Bob",
						Cost:   54 + 83,
					},
					{
						Parent: "Alice",
						Child:  "Carol",
						Cost:   -79 - 62,
					},
					{
						Parent: "Bob",
						Child:  "Carol",
						Cost:   -7 + 60,
					},
				},
				Nodes: []string{"Alice", "Bob", "Carol"},
				Paths: [][]string{
					{"Alice", "Bob", "Carol"},
					{"Bob", "Alice", "Carol"},
					{"Carol", "Bob", "Alice"},
					{"Bob", "Carol", "Alice"},
					{"Carol", "Alice", "Bob"},
					{"Alice", "Carol", "Bob"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{}
			if err := g.ParseInput(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("Graph.ParseInput() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(g.Nodes, tt.want.Nodes) {
				t.Errorf("Graph.AddEdge().Nodes = %v, want %v", g.Nodes, tt.want.Nodes)
			}
			if !reflect.DeepEqual(g.Paths, tt.want.Paths) {
				t.Errorf("Graph.AddEdge().Paths = %v, want %v", g.Paths, tt.want.Paths)
			}
			for i, e := range g.Edges {
				if want := tt.want.Edges[i]; *want != *e {
					t.Errorf("Graph.AddEdge().Edge = %v, want %v", *e, *want)
				}
			}
		})
	}
}

func TestGraph_GetDistanceOfPath(t *testing.T) {
	tests := []struct {
		name  string
		Edges []*Edge
		arg   []string
		want  int
	}{
		{
			name: "gets the distance of a path in the graph",
			Edges: []*Edge{
				{
					Parent: "Carol",
					Child:  "David",
					Cost:   4000,
				},
				{
					Parent: "Alice",
					Child:  "Bob",
					Cost:   1,
				},
				{
					Parent: "Alice",
					Child:  "David",
					Cost:   50000,
				},
				{
					Parent: "Ernest",
					Child:  "David",
					Cost:   300,
				},
				{
					Parent: "Carol",
					Child:  "Bob",
					Cost:   20,
				},
			},
			arg:  []string{"Alice", "Bob", "Carol", "David"},
			want: 54021,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Edges: tt.Edges,
			}
			if got := g.GetDistanceOfPath(tt.arg); got != tt.want {
				t.Errorf("Graph.GetDistanceOfPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_FindGreatestHappiness(t *testing.T) {
	tests := []struct {
		name  string
		Edges []*Edge
		Paths [][]string
		want  int
	}{
		{
			name: "finds the gratest happiness number, advent of code example",
			Edges: []*Edge{
				{
					Parent: "Alice",
					Child:  "Bob",
					Cost:   54 + 83,
				},
				{
					Parent: "Alice",
					Child:  "Carol",
					Cost:   -79 - 62,
				},
				{
					Parent: "Alice",
					Child:  "David",
					Cost:   -2 + 46,
				},
				{
					Parent: "Bob",
					Child:  "Carol",
					Cost:   -7 + 60,
				},
				{
					Parent: "Bob",
					Child:  "David",
					Cost:   -63 - 7,
				},
				{
					Parent: "Carol",
					Child:  "David",
					Cost:   55 + 41,
				},
			},
			Paths: [][]string{
				{"Alice", "Bob", "Carol", "David"},
				{"Alice", "Bob", "David", "Carol"},
				{"Alice", "Carol", "Bob", "David"},
				{"Alice", "Carol", "David", "Bob"},
				{"Alice", "David", "Carol", "Bob"},
				{"Alice", "David", "Bob", "Carol"},
				{"Bob", "Alice", "Carol", "David"},
				{"Bob", "Alice", "David", "Carol"},
				{"Bob", "Carol", "Alice", "David"},
				{"Bob", "Carol", "David", "Alice"},
				{"Bob", "David", "Carol", "Alice"},
				{"Bob", "David", "Alice", "Carol"},
				{"Carol", "Bob", "Alice", "David"},
				{"Carol", "Bob", "David", "Alice"},
				{"Carol", "Alice", "Bob", "David"},
				{"Carol", "Alice", "David", "Bob"},
				{"Carol", "David", "Alice", "Bob"},
				{"Carol", "David", "Bob", "Alice"},
				{"David", "Bob", "Carol", "Alice"},
				{"David", "Bob", "Alice", "Carol"},
				{"David", "Carol", "Bob", "Alice"},
				{"David", "Carol", "Alice", "Bob"},
				{"David", "Alice", "Carol", "Bob"},
				{"David", "Alice", "Bob", "Carol"},
			},
			want: 330,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Graph{
				Edges: tt.Edges,
				Paths: tt.Paths,
			}
			if got := g.FindGreatestHappiness(); got != tt.want {
				t.Errorf("Graph.FindGreatestHappiness() = %v, want %v", got, tt.want)
			}
		})
	}
}
