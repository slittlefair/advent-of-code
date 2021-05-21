package main

import (
	"reflect"
	"testing"
)

func TestGraph_addNode(t *testing.T) {
	tests := []struct {
		name  string
		nodes []string
		n     string
		want  []string
	}{
		{
			name: "doesn't add a node if it's already in the graph",
			nodes: []string{
				"Australia",
				"Belgium",
				"Cyprus",
			},
			n: "Belgium",
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
			n: "Denmark",
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
			g.addNode(tt.n)
			if !reflect.DeepEqual(g.Nodes, tt.want) {
				t.Errorf("Graph.addNode() = %v, want %v", g.Nodes, tt.want)
			}
		})
	}
}

func TestGraph_addEdge(t *testing.T) {
	// cyprus := string{name: "Cyprus"}
	// denmark := string{name: "Denmark"}
	type fields struct {
		Edges []Edge
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
				Edges: []Edge{
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
				Edges: []Edge{
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
				Edges: []Edge{
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
			name: "adds just an edge to a graph (parent and child already exist)",
			args: args{
				parent: "Australia",
				child:  "Belgium",
				cost:   44,
			},
			fields: fields{
				Nodes: []string{"Australia", "Belgium"},
			},
			want: &Graph{
				Edges: []Edge{
					{
						Parent: "Australia",
						Child:  "Belgium",
						Cost:   44,
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
			g.addEdge(tt.args.parent, tt.args.child, tt.args.cost)
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("Graph.addEdge() = %v, want %v", g, tt.want)
			}
		})
	}
}

func TestGraph_parseInput(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    *Graph
		wantErr bool
	}{
		{
			name: "returns an error if an input line cost can't be converted correctly",
			input: []string{
				"London to Dublin = 464",
				"London to Belfast = around 518",
				"Dublin to Belfast = 141",
			},
			want: &Graph{
				Edges: []Edge{
					{
						Parent: "London",
						Child:  "Dublin",
						Cost:   464,
					},
				},
				Nodes: []string{"London", "Dublin"},
			},
			wantErr: true,
		},
		{
			name: "populates graph from input",
			input: []string{
				"London to Dublin = 464",
				"London to Belfast = 518",
				"Dublin to Belfast = 141",
			},
			want: &Graph{
				Edges: []Edge{
					{
						Parent: "London",
						Child:  "Dublin",
						Cost:   464,
					},
					{
						Parent: "London",
						Child:  "Belfast",
						Cost:   518,
					},
					{
						Parent: "Dublin",
						Child:  "Belfast",
						Cost:   141,
					},
				},
				Nodes: []string{
					"London",
					"Dublin",
					"Belfast",
				},
				Paths: [][]string{
					{"London", "Dublin", "Belfast"},
					{"Dublin", "London", "Belfast"},
					{"Belfast", "Dublin", "London"},
					{"Dublin", "Belfast", "London"},
					{"Belfast", "London", "Dublin"},
					{"London", "Belfast", "Dublin"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{}
			if err := g.parseInput(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("Graph.parseInput() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("Graph.parseInput() = %+v, want %+v", g, tt.want)
			}
		})
	}
}

func TestGraph_getDistanceOfPath(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want int
	}{
		{
			name: "returns path Dublin -> London -> Belfast = 982",
			arg:  []string{"Dublin", "London", "Belfast"},
			want: 982,
		},
		{
			name: "returns path London -> Dublin -> Belfast = 605",
			arg:  []string{"London", "Dublin", "Belfast"},
			want: 605,
		},
		{
			name: "returns path London -> Belfast -> Dublin = 659",
			arg:  []string{"London", "Belfast", "Dublin"},
			want: 659,
		},
		{
			name: "returns path Dublin -> Belfast -> London = 659",
			arg:  []string{"Dublin", "Belfast", "London"},
			want: 659,
		},
		{
			name: "returns path Belfast -> Dublin -> London = 605",
			arg:  []string{"Belfast", "Dublin", "London"},
			want: 605,
		},
		{
			name: "returns path Belfast -> London -> Dublin = 982",
			arg:  []string{"Belfast", "London", "Dublin"},
			want: 982,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Edges: []Edge{{
					Parent: "London",
					Child:  "Dublin",
					Cost:   464,
				},
					{
						Parent: "London",
						Child:  "Belfast",
						Cost:   518,
					},
					{
						Parent: "Dublin",
						Child:  "Belfast",
						Cost:   141,
					}},
			}
			if got := g.getDistanceOfPath(tt.arg); got != tt.want {
				t.Errorf("Graph.getDistanceOfPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_findMinimumAndMaximumPaths(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			name:  "returns the correct minimum and maximum paths",
			want:  605,
			want1: 982,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Edges: []Edge{
					{
						Parent: "London",
						Child:  "Dublin",
						Cost:   464,
					},
					{
						Parent: "London",
						Child:  "Belfast",
						Cost:   518,
					},
					{
						Parent: "Dublin",
						Child:  "Belfast",
						Cost:   141,
					},
				},
				Nodes: []string{
					"London",
					"Dublin",
					"Belfast",
				},
				Paths: [][]string{
					{"London", "Dublin", "Belfast"},
					{"Dublin", "London", "Belfast"},
					{"Belfast", "Dublin", "London"},
					{"Dublin", "Belfast", "London"},
					{"Belfast", "London", "Dublin"},
					{"London", "Belfast", "Dublin"},
				},
			}
			got, got1 := g.findMinimumAndMaximumPaths()
			if got != tt.want {
				t.Errorf("Graph.findMinimumAndMaximumPaths() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Graph.findMinimumAndMaximumPaths() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
