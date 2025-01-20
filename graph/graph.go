package graph

import "fmt"

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
