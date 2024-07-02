package graph

import "testing"

func TestGraph_BFS(t *testing.T) {
	graph := Graph[string]{
		"you":    map[string]int{"alice": 1, "bob": 1, "claire": 1},
		"bob":    map[string]int{"anuj": 1, "peggy": 1},
		"alice":  map[string]int{"peggy": 1},
		"claire": map[string]int{"thom": 1, "jonny": 1},
		"anuj":   map[string]int{},
		"peggy":  map[string]int{},
		"thom":   map[string]int{},
		"jonny":  map[string]int{},
	}

	type args struct {
		start string
		cmp   func(value string) bool
	}
	tests := []struct {
		name  string
		g     Graph[string]
		args  args
		want  string
		want1 bool
	}{
		{"EmptyGraph", make(Graph[string]), args{"you", func(value string) bool { return value != "" }}, "", false},
		{"Graph", graph, args{"you", func(value string) bool { return value[len(value)-1] == 'm' }}, "thom", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.g.BFS(tt.args.start, tt.args.cmp)
			if got != tt.want {
				t.Errorf("Graph.BFS() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Graph.BFS() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
