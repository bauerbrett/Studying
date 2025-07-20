package main

import (
	"fmt"
	"sort"
)

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	visited := make([]int, n) // 0 = unvisited, 1 = visiting, 2 = safe
	result := []int{}

	var dfs func(int) bool
	dfs = func(node int) bool {
		if visited[node] == 2 {
			return true
		}
		if visited[node] == 1 {
			return false
		}
		visited[node] = 1
		for _, neighbor := range graph[node] {
			if !dfs(neighbor) {
				return false
			}
		}
		visited[node] = 2
		return true
	}
	for i := 0; i < len(graph); i++ {
		if dfs(i) {
			result = append(result, i)
		}
	}
	/*
		could also have this if you just ran dfs(i) in loop. But probably dont want to double loop
		for i, val := range visited {
			if val == 2 {
				result = append(result, i)
			}
		}
	*/
	sort.Ints(result)
	return result
}

func main() {
	g := [][]int{
		{1, 2},
		{2, 3},
		{5},
		{0},
		{5},
		{},
		{},
	}
	fmt.Println(eventualSafeNodes(g)) // Output: [2 4 5 6]
}
