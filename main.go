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
	g.AddEdge('C', 'A')
	g.PrintMatrix()
	g.PrintGraph()
	result, err := g.EulerCircuit()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	fmt.Printf("\n")

	h := graph.NewGraphMatrix()
	h.AddEdge('A', 'B')
	h.AddEdge('A', 'C')
	h.AddEdge('B', 'C')
	h.AddEdge('B', 'D')
	h.AddEdge('B', 'E')
	h.AddEdge('C', 'D')
	h.AddEdge('C', 'E')
	h.AddEdge('D', 'E')
	h.PrintMatrix()
	h.PrintGraph()

	result, err = h.EulerCircuit()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	fmt.Printf("\n")

	i := graph.NewGraphMatrix()
	i.AddEdge('A', 'B')
	i.AddEdge('A', 'F')
	i.AddEdge('F', 'E')
	i.AddEdge('E', 'B')
	i.AddEdge('B', 'D')
	i.AddEdge('B', 'C')
	i.AddEdge('D', 'C')
	i.PrintMatrix()
	i.PrintGraph()

	result, err = i.EulerCircuit()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
