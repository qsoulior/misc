package graph

import (
	"reflect"
	"slices"
	"testing"
)

func emptyWeightedGraph() WeightedGraph[string] { return make(WeightedGraph[string]) }

func simpleWeightedGraph() WeightedGraph[string] {
	return WeightedGraph[string]{
		"book":   map[string]int{"record": 5, "poster": 0},
		"record": map[string]int{"guitar": 15, "drum": 20},
		"poster": map[string]int{"guitar": 30, "drum": 35},
		"guitar": map[string]int{"piano": 20},
		"drum":   map[string]int{"piano": 10},
	}
}

func TestWeightedGraph_Unweighted(t *testing.T) {
	want := UnweightedGraph[string]{
		"book":   []string{"poster", "record"},
		"record": []string{"drum", "guitar"},
		"poster": []string{"drum", "guitar"},
		"guitar": []string{"piano"},
		"drum":   []string{"piano"},
	}

	tests := []struct {
		name string
		g    WeightedGraph[string]
		want UnweightedGraph[string]
	}{
		{"EmptyGraph", emptyWeightedGraph(), emptyUnweightedGraph()},
		{"SimpleGraph", simpleWeightedGraph(), want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.g.Unweighted()
			for _, nodes := range got {
				slices.Sort(nodes)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeightedGraph.Unweighted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeightedGraph_Dijkstra(t *testing.T) {
	want := map[string]int{"book": 0, "drum": 25, "guitar": 20, "piano": 35, "poster": 0, "record": 5}
	want1 := map[string]string{"drum": "record", "guitar": "record", "piano": "drum", "poster": "book", "record": "book"}

	type args struct {
		start string
	}
	tests := []struct {
		name  string
		g     WeightedGraph[string]
		args  args
		want  map[string]int
		want1 map[string]string
	}{
		{"EmptyGraph", emptyWeightedGraph(), args{"book"}, nil, nil},
		{"SimpleGraph", simpleWeightedGraph(), args{"book"}, want, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.g.Dijkstra(tt.args.start)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeightedGraph.Dijkstra() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("WeightedGraph.Dijkstra() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWeightedGraph_QuickDijkstra(t *testing.T) {
	want := map[string]int{"book": 0, "drum": 25, "guitar": 20, "piano": 35, "poster": 0, "record": 5}
	want1 := map[string]string{"drum": "record", "guitar": "record", "piano": "drum", "poster": "book", "record": "book"}

	type args struct {
		start string
	}
	tests := []struct {
		name  string
		g     WeightedGraph[string]
		args  args
		want  map[string]int
		want1 map[string]string
	}{
		{"EmptyGraph", emptyWeightedGraph(), args{"book"}, nil, nil},
		{"SimpleGraph", simpleWeightedGraph(), args{"book"}, want, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.g.QuickDijkstra(tt.args.start)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeightedGraph.QuickDijkstra() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("WeightedGraph.QuickDijkstra() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
