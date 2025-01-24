package graph

import (
	"fmt"
)

type GraphMatrix struct {
	vertices map[rune]int // maps a vertex to its index
	indexMap []rune       // maps an index to a vertex
	matrix   [][]int      // adjacency matrix
	size     int          // number of vertices
}

func NewGraphMatrix() *GraphMatrix {
	return &GraphMatrix{
		vertices: make(map[rune]int),
		indexMap: []rune{},
		matrix:   [][]int{},
		size:     0,
	}
}

func (g *GraphMatrix) AddVertex(vertex rune) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = g.size
		g.indexMap = append(g.indexMap, vertex)
		g.size++
		// Expand matrix
		for i := range g.matrix {
			g.matrix[i] = append(g.matrix[i], 0)
		}
		g.matrix = append(g.matrix, make([]int, g.size))
	}
}

func (g *GraphMatrix) AddEdge(from, to rune) {
	g.AddVertex(from)
	g.AddVertex(to)
	i, j := g.vertices[from], g.vertices[to]
	g.matrix[i][j] = 1
	g.matrix[j][i] = 1
}

func (g *GraphMatrix) PrintMatrix() {
	fmt.Print("  ")
	for _, label := range g.indexMap {
		fmt.Printf("%c ", label)
	}
	fmt.Println()
	for i, row := range g.matrix {
		fmt.Printf("%c ", g.indexMap[i])
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func (g *GraphMatrix) PrintGraph() {
	fmt.Printf("vertices: %v\n", g.vertices)
	fmt.Printf("indexMap: %v\n", g.indexMap)
	fmt.Printf("matrix:   %v\n", g.matrix)
	fmt.Printf("size:     %v\n", g.size)
}

func (g *GraphMatrix) getDegree(vertex int) int {
	degree := 0
	for i := 0; i < g.size; i++ {
		degree += g.matrix[vertex][i]
	}
	return degree
}

func (g *GraphMatrix) getUnusedEdges(used [][]bool) map[int][]int {
	unused := make(map[int][]int)
	for i := 0; i < g.size; i++ {
		unused[i] = make([]int, 0)
		for j := 0; j < g.size; j++ {
			if g.matrix[i][j] == 1 && !used[i][j] {
				unused[i] = append(unused[i], j)
			}
		}
	}
	return unused
}

func (g *GraphMatrix) PrintEulerCircuit() {
	vertexList, err := g.EulerCircuit()
	if err != nil {
		fmt.Println(err)
	}

	for i, vertex := range vertexList {
		fmt.Printf("%c", vertex)
		if i < len(vertexList)-1 {
			fmt.Print(" - ")
		}
	}

}

/*
BEGIN
  IF graph infeasible THEN END
  start ← suitable node
  tour ← {start}
  REPEAT
	current = start ← node in tour with
				  unvisited edge
	subtour ← {start}
	DO
	  {current, u} ← take unvisited edge
	  subtour ← subtour ∪ {u}
	  current ← u
	WHILE start ≠ current
	Integrate subtour in tour
  UNTIL tour is Eulerian path/cycle
END
*/

func (g *GraphMatrix) EulerCircuit() ([]rune, error) {
	// Check if graph is feasible
	// at least one edge
	// connected graph
	// all vertices or exactly two vertices have even degree
	if g.size <= 1 {
		return nil, fmt.Errorf("graph is not feasible: at least two vertices are required")
	}

	numberOddDegree := 0
	for i := 0; i < g.size; i++ {
		if g.getDegree(i)%2 != 0 {
			numberOddDegree++
		}
	}
	if numberOddDegree != 0 && numberOddDegree != 2 {
		return nil, fmt.Errorf("graph is not feasible: all vertices or exactly two vertices need to have an even degree")
	}

	// Track used edges
	used := make([][]bool, g.size)
	for i := range used {
		used[i] = make([]bool, g.size)
	}

	// Start circuit from first vertex
	circuit := []int{0}
	current := 0

	for {
		unused := g.getUnusedEdges(used)

		if len(unused[current]) > 0 {
			// Follow unused edge
			next := unused[current][0]
			used[current][next] = true
			used[next][current] = true
			current = next
			circuit = append(circuit, current)
		} else {
			// Look for vertex with unused edges
			found := false
			for i, vertex := range circuit {
				if len(unused[vertex]) > 0 {
					// Start new subcircuit
					subStart := vertex
					subCurrent := vertex
					subCircuit := []int{subStart}

					for len(unused[subCurrent]) > 0 {
						next := unused[subCurrent][0]
						used[subCurrent][next] = true
						used[next][subCurrent] = true
						subCurrent = next
						subCircuit = append(subCircuit, subCurrent)
					}

					// Splice subcircuit into main circuit
					newCircuit := make([]int, 0)
					newCircuit = append(newCircuit, circuit[:i]...)
					newCircuit = append(newCircuit, subCircuit...)
					newCircuit = append(newCircuit, circuit[i+1:]...)
					circuit = newCircuit

					found = true
					break
				}
			}

			if !found {
				break // Circuit complete
			}
		}
	}

	// Convert vertex indices back to runes
	result := make([]rune, len(circuit))
	for i, idx := range circuit {
		result[i] = g.indexMap[idx]
	}

	return result, nil
}
