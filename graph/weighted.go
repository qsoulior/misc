package graph

import (
	"github.com/qsoulior/alg/queue"
	"github.com/qsoulior/alg/set"
)

// WeightedGraph is graph with weights represented as adjacency map.
type WeightedGraph[T comparable] map[T]map[T]int

// Unweighted creates unweighted graph from g and returns int.
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

// Dijkstra represents Dijkstra's algorithm with complexity O(n^2),
// where n is number of vertices.
// Algorithm starts from vertex start and returns distance map and parent map.
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

// QuickDijkstra represents Dijkstra's algorithm with complexity O(m*log(m)),
// where m is number of edges.
// Algorithm starts from vertex start and returns distance map and parent map.
func (g WeightedGraph[T]) QuickDijkstra(start T) (map[T]int, map[T]T) {
	if _, ok := g[start]; !ok {
		return nil, nil
	}

	dists := map[T]int{start: 0}
	parents := make(map[T]T)

	// Minimum distance has the highest priority.
	queue := queue.NewMinPriorityQueue[T]()
	queue.Push(start, 0)

	for queue.Len() > 0 {
		// PopFront returns node with minimum distance, O(log(m)).
		minNode, minDist, _ := queue.PopFront()
		if minDist > dists[minNode] {
			continue
		}

		// Update minimum distances to neighbors.
		for neighbor, weight := range g[minNode] {
			newDist := minDist + weight
			if dist, ok := dists[neighbor]; !ok || newDist < dist {
				dists[neighbor] = newDist
				parents[neighbor] = minNode
				queue.Push(neighbor, newDist)
			}
		}
	}

	return dists, parents
}
