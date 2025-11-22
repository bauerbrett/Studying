package main

import "fmt"

type Graph struct {
	vert int
	list [][]int
}

func Constructor(v int) *Graph {
	g := &Graph{vert: v, list: make([][]int, v)}
	return g
}

func (g *Graph) addEdges(a, b int) {
	g.list[a] = append(g.list[a], b)
	g.list[b] = append(g.list[b], a)
}

func pathExist(g *Graph, start, end int) bool {
	visited := make([]bool, g.vert)

	var dfs func(*Graph, int, int) bool
	dfs = func(g *Graph, a, b int) bool {
		visited[a] = true
		if a == b {
			return true
		}
		path := false
		for _, neighbor := range g.list[a] {
			if !visited[neighbor] {
				path = path || dfs(g, neighbor, end)
				/* same as
				if dfs(g, neighbor, end) {
					path = true
				}
				*/
			}
		}
		return path

	}
	return dfs(g, start, end)
}

func main() {
	g := Constructor(4)
	g.addEdges(0, 1)
	g.addEdges(1, 2)
	g.addEdges(2, 3)
	fmt.Println(pathExist(g, 0, 3))
}
