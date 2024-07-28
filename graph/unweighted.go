package graph

import (
	"github.com/qsoulior/alg/queue"
	"github.com/qsoulior/alg/set"
)

type UnweightedGraph[T comparable] map[T][]T

// Breadth-first search, O(n+m).
func (g UnweightedGraph[T]) BFS(start T, cmp func(value T) bool) (T, bool) {
	deque := queue.NewDeque[T]()
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

// Depth-first search, O(n+m).
func (g UnweightedGraph[T]) DFS(start T, cmp func(value T) bool) (T, bool) {
	deque := queue.NewDeque[T]()
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
