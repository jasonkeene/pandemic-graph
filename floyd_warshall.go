package pandemic

import "math"

// FloydWarshall traverses the graph and computes the shortest path lengths
// between all pairs. It also computes predicates so you can rebuild the
// shortest path between two pairs.
func FloydWarshall(vertices []string, adjacencyList []IntSet) (Matrix, Matrix) {
	// init shortest
	shortest := make(Matrix, len(vertices))
	for u := range shortest {
		shortest[u] = make([]int, len(vertices))
		for v := range shortest[u] {
			shortest[u][v] = math.MaxInt32 - 2
		}
	}

	// init pred
	pred := make(Matrix, len(vertices))
	for u := range pred {
		pred[u] = make([]int, len(vertices))
		for v := range pred[u] {
			pred[u][v] = -1
		}
	}

	// populate existing values for shortest and pred based on edges
	for u := range vertices {
		shortest[u][u] = 0
		for e := range adjacencyList[u] {
			shortest[u][e] = 1
			pred[u][e] = u
		}
	}

	// consider every path of u->v, if it can be improved by visiting x do so
	for x := range vertices {
		for u := range vertices {
			for v := range vertices {
				// see if going from u->v via x is shorter
				shortestViaX := shortest[u][x] + shortest[x][v]
				if shortestViaX < shortest[u][v] {
					// update the shortest val
					shortest[u][v] = shortestViaX
					// pred should now be pred between x->v
					pred[u][v] = pred[x][v]
				}
			}
		}
	}

	return shortest, pred
}
