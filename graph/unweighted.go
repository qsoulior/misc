// Package graph implements graph data structures and algorithms.
// It provides unweighted and weighted graph implementations, BFS, DFS and Dijkstra's algorithm.
package graph

import (
	"github.com/qsoulior/misc/queue"
	"github.com/qsoulior/misc/set"
)

// UnweightedGraph is graph without weights represented as adjacency map.
type UnweightedGraph[T comparable] map[T][]T

// BFS represents breadth-first search with complexity O(n+m),
// where n is number of vertices and m is number of edges.
// BFS starts from vertex start and uses cmp to compare each vertex with target.
// It returns found vertex or default value of type T and false as second value.
func (g UnweightedGraph[T]) BFS(start T, cmp func(value T) bool) (T, bool) {
	deque := queue.NewListDeque[T]()
	enqueued := make(set.HashSet[T])

	if _, ok := g[start]; ok {
		deque.PushBack(start)
		enqueued.Add(start)
	}

	for deque.Len() > 0 {
		value, _ := deque.PopFront()
		if cmp(value) {
			return value, true
		}

		for _, adjacent := range g[value] {
			if !enqueued.Contains(adjacent) {
				deque.PushBack(adjacent)
				enqueued.Add(adjacent)
			}
		}
	}

	var value T
	return value, false
}

// DFS represents depth-first search with complexity O(n+m),
// where n is number of vertices and m is number of edges.
// DFS starts from vertex start and uses cmp to compare each vertex with target.
// It returns found vertex or default value of type T and false as second value.
func (g UnweightedGraph[T]) DFS(start T, cmp func(value T) bool) (T, bool) {
	deque := queue.NewListDeque[T]()
	enqueued := make(set.HashSet[T])

	if _, ok := g[start]; ok {
		deque.PushBack(start)
		enqueued.Add(start)
	}

	for deque.Len() > 0 {
		value, _ := deque.PopBack()
		if cmp(value) {
			return value, true
		}

		adjacents := g[value]
		for i := len(adjacents) - 1; i >= 0; i-- {
			if adjacent := adjacents[i]; !enqueued.Contains(adjacent) {
				deque.PushBack(adjacent)
				enqueued.Add(adjacent)
			}
		}
	}

	var value T
	return value, false
}
