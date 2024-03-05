package dijkstra

import (
	"fmt"
	"math"
)

// no heap implementation for now
func dijkstra(num_vertices int, edges [][]int, start_vertex int, end_vertex int) int {
	if start_vertex < 0 || end_vertex < 0 ||
		start_vertex >= num_vertices || end_vertex >= num_vertices {
		return math.MaxInt
	}

	if len(edges) != num_vertices {
		return math.MaxInt
	}
	for i := range edges {
		if len(edges[i]) != num_vertices {
			return math.MaxInt
		}
	}

	d := make([]int, num_vertices) // estimated distances
	for i := range d {
		d[i] = math.MaxInt
	}
	d[start_vertex] = 0

	pi := make([]int, num_vertices) // pi is reference to parents
	for i := range pi {
		pi[i] = -1
	}

	unvisited := map[int]struct{}{} // unvisited
	for i := 0; i < num_vertices; i++ {
		unvisited[i] = struct{}{}
	}

	// main algo
	for len(unvisited) > 0 {
		// find unvisited vertex with minimum distance
		mindistvertex := -1
		for k := range unvisited {
			if mindistvertex == -1 {
				mindistvertex = k
			} else {
				if d[k] < d[mindistvertex] {
					mindistvertex = k
				}
			}
		}

		// for each neighbor v of this min dist vertex
		for v, neighboredgeweight := range edges[mindistvertex] {
			if neighboredgeweight < 0 { // edge doesnt exist
				continue
			}

			if d[v] > d[mindistvertex] + neighboredgeweight {
				d[v] = d[mindistvertex] + neighboredgeweight
				pi[v] = mindistvertex
				fmt.Println(v, d[v])
			}
		}

		// remove from unvisited
		delete(unvisited, mindistvertex)
	}

	return d[end_vertex]
}
