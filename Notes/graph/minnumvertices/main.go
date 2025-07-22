package main

import "fmt"

/*
Given a directed acyclic graph with n nodes labeled from 0 to n-1, determine the smallest number of initial nodes
such that you can access all the nodes by traversing edges. Return these nodes.

Examples
Example 1:
Input:
n = 6
edges = [[0,1],[0,2],[2,5],[3,4],[4,2]]
These are just edges and not nodes - so example - 0 goes to 1, 0 goes to 2, 2 goes to 5, etc.
Expected Output: [0,3]

We need to find the nodes that have no incoming edges which means those are teh ones that we start with.
*/

func minVertices(l [][]int, n int) []int {
	inDegree := make([]int, n)
	result := []int{}
	for _, e := range l {
		_, v := e[0], e[1] // the second num in inner slice is a node that has incoming
		inDegree[v]++      // we want to find the nodes with 0 incoming so we add 1 to the nodes that have incoming so we can search for the ones with 0
	}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			result = append(result, i)
		}
	}

	return result

}

func main() {
	edgeList := [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}
	fmt.Println(minVertices(edgeList, 6))
}
