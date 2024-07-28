package graph

import (
	"github.com/qsoulior/alg/queue"
	"github.com/qsoulior/alg/set"
)

type WeightedGraph[T comparable] map[T]map[T]int

func (g WeightedGraph[T]) Unweighted() UnweightedGraph[T] {
	graph := make(UnweightedGraph[T], len(g))
	for value, adjacents := range g {
		graph[value] = make([]T, 0, len(adjacents))
		for adjacent := range adjacents {
			graph[value] = append(graph[value], adjacent)
		}
	}

	return graph
}

// Dijkstra's algorithm, O(n^2).
func (g WeightedGraph[T]) Dijkstra(start T) (map[T]int, map[T]T) {
	if _, ok := g[start]; !ok {
		return nil, nil
	}

	dists := map[T]int{start: 0}
	parents := make(map[T]T)
	processed := make(set.HashSet[T])

	minNode, minDist := start, 0
	for minDist >= 0 {
		// Update minimum distances to neighbors.
		for neighbor, weight := range g[minNode] {
			newDist := minDist + weight
			oldDist, ok := dists[neighbor]
			if !ok || newDist < oldDist {
				dists[neighbor] = newDist
				parents[neighbor] = minNode
			}
		}

		processed.Add(minNode)

		// Find unprocessed node with minimum distance, O(n).
		minDist = -1
		for node, dist := range dists {
			if !processed.Contains(node) && (minDist == -1 || dist < minDist) {
				minDist = dist
				minNode = node
			}
		}
	}

	return dists, parents
}

// Dijkstra's algorithm, O(m*log(m)).
func (g WeightedGraph[T]) QuickDijkstra(start T) (map[T]int, map[T]T) {
	if _, ok := g[start]; !ok {
		return nil, nil
	}

	dists := map[T]int{start: 0}
	parents := make(map[T]T)

	// Since PopFront returns value with maximum priority,
	// distances are stored as negative priorities.
	queue := queue.NewPriorityQueue[T]()
	queue.Push(start, 0)

	for queue.Len() > 0 {
		minNode, minDist, _ := queue.PopFront()

		minDist = -minDist // minimum distance is maximum negative priority
		if minDist > dists[minNode] {
			continue
		}

		// Update minimum distances to neighbors.
		for neighbor, weight := range g[minNode] {
			newDist := minDist + weight
			if dist, ok := dists[neighbor]; !ok || newDist < dist {
				dists[neighbor] = newDist
				parents[neighbor] = minNode
				queue.Push(neighbor, -newDist) // save distance as negative priority
			}
		}
	}

	return dists, parents
}
