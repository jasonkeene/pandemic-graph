package main

import (
	"fmt"

	"github.com/jasonkeene/pandemic-graph"
)

func main() {
	shortest, _ := pandemic.FloydWarshall(
		pandemic.Vertices,
		pandemic.AdjacencyList,
	)

	// output general information
	fmt.Print("General Information:\n")
	maxDistance := shortest.Max()
	fmt.Printf("    Average distance between all nodes: %f\n", shortest.Average())
	fmt.Printf("    Largest distance between any two nodes: %d\n", maxDistance)

	maxPairs := shortest.SortedPairs()[maxDistance]
	for _, pair := range maxPairs {
		fmt.Printf(
			"        Pair: %s -> %s\n",
			pandemic.Vertices[pair.A],
			pandemic.Vertices[pair.B],
		)
	}
}
