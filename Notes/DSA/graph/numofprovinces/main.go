package main

import "fmt"

/*
There are n cities. Some of them are connected in a network. If City A is directly connected to City B, and City B is directly connected to City C, city A is indirectly connected to City C.

If a group of cities are connected directly or indirectly, they form a province.

Given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise, determine the total number of provinces.

Examples
Example 1:
Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
Image
Expected Output: 2

Justification: Here, city 1 and 2 form a single provenance, and city 3 is one provinces itself.

Plan: Use DFS and visit each graph[i] which is the node. This counts as a provinence. In the node visit its neighbors,
if the neighbors but dont add 1 for the neighbors. Just add them to visited list. When you get to new node that is unvisited outside of neighbors, that is new
province add 1

Solution: Basically loop through each nested list. In the list recursively visit the connected cities. If you hit places inside the recursion that arent
visited and == 1 (means they are connected), dfs it. The dfs will then go and search that city and do the same thing. What this will do is inside the dfs
the visits to the cities are not count++ because they are technically still in the loop cycle. But once the loop goes to the next cycle it will either hit
a city that was already visited in one of the dfs calls, or it will not have been visited and will count++ and then go and recursively visit that one.
*/

func numProvinces(l [][]int) int {
	n := len(l)
	//fmt.Println(l[0][3])
	visited := make([]bool, len(l))
	count := 0

	var dfs func(int)
	dfs = func(node int) {
		visited[node] = true
		for neighbor := 0; neighbor < n; neighbor++ {
			if l[node][neighbor] == 1 && !visited[neighbor] {
				dfs(neighbor)
			}

		}
	}
	for city := 0; city < len(l); city++ {
		if !visited[city] {
			count++
			dfs(city)
		}

	}
	return count
}

func main() {
	l := [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}}
	fmt.Println(numProvinces(l))
}
