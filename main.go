package main

import (
	"erik-schuetze/eulercircuit/graph"
	"fmt"
)

func main() {
	fmt.Println("fully connected graph: ")
	g := graph.NewGraphMatrix()
	g.AddEdge('A', 'B')
	g.AddEdge('B', 'C')
	g.AddEdge('C', 'D')
	g.AddEdge('D', 'A')
	g.PrintMatrix()
	g.PrintGraph()

	fmt.Printf("\n")

	fmt.Println("graph missing one connection: ")
	h := graph.NewGraphMatrix()
	h.AddEdge('A', 'B')
	h.AddEdge('B', 'C')
	h.AddEdge('C', 'D')
	h.PrintMatrix()
	h.PrintGraph()
}
