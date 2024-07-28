package graph

import "testing"

func emptyUnweightedGraph() UnweightedGraph[string] { return make(UnweightedGraph[string]) }

func simpleUnweightedGraph() UnweightedGraph[string] {
	return UnweightedGraph[string]{
		"you":    []string{"bob", "claire", "jane"},
		"bob":    []string{"anuj", "peggy"},
		"jane":   []string{"peggy"},
		"claire": []string{"jonny"},
		"anuj":   []string{},
		"peggy":  []string{},
		"jonny":  []string{},
	}
}

func unweightedCmp(value string) bool { return len(value) > 0 && value[0] == 'j' }

func TestUnweightedGraph_BFS(t *testing.T) {
	type args struct {
		start string
		cmp   func(value string) bool
	}
	tests := []struct {
		name  string
		g     UnweightedGraph[string]
		args  args
		want  string
		want1 bool
	}{
		{"EmptyGraph", emptyUnweightedGraph(), args{"you", unweightedCmp}, "", false},
		{"SimpleGraph", simpleUnweightedGraph(), args{"you", unweightedCmp}, "jane", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.g.BFS(tt.args.start, tt.args.cmp)
			if got != tt.want {
				t.Errorf("UnweightedGraph.BFS() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UnweightedGraph.BFS() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUnweightedGraph_DFS(t *testing.T) {
	type args struct {
		start string
		cmp   func(value string) bool
	}
	tests := []struct {
		name  string
		g     UnweightedGraph[string]
		args  args
		want  string
		want1 bool
	}{
		{"EmptyGraph", emptyUnweightedGraph(), args{"you", unweightedCmp}, "", false},
		{"SimpleGraph", simpleUnweightedGraph(), args{"you", unweightedCmp}, "jonny", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.g.DFS(tt.args.start, tt.args.cmp)
			if got != tt.want {
				t.Errorf("UnweightedGraph.DFS() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UnweightedGraph.DFS() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
