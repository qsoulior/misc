package graph

import (
	"github.com/qsoulior/alg/queue"
	"github.com/qsoulior/alg/set"
)

type Graph[T comparable] map[T]map[T]int

func (g Graph[T]) BFS(start T, cmp func(value T) bool) (T, bool) {
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

		for adjacent := range g[value] {
			if !enqueued.Contains(adjacent) {
				deque.PushBack(adjacent)
				enqueued.Add(adjacent)
			}
		}
	}

	var value T
	return value, false
}
