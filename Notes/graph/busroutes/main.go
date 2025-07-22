package main

import (
	"fmt"
)

/*
You are given an array routes where routes{i} is the list of bus stops that the ithe bus travels in a cyclic manner. For example,
if routes{0} = {2, 3, 7}, it means that bus 0 travels through the stops 2 -> 3 -> 7 -> 2 -> 3 -> 7 ... and then repeats this sequence indefinitely.

You start at a bus stop called source and wish to travel to a bus stop called target using the bus routes.
You can switch buses at any bus stop that is common to the routes of two buses.

Return the minimum number of buses you need to take to travel from source to target. If it is not possible to reach the target, return -1.

Examples
Example 1

Input: routes = {{2, 3, 4}, {5, 6, 7, 8}, {4, 5, 9, 10}, {10, 12}}, source = 3, target = 12
Expected Output: 3
Justification: Start at stop 3, take bus 0 to stop 4, switch to bus 2 to reach stop 10, and then take bus 3 to reach to 12. You need 3 buses.
*/

func busRoutes(r [][]int, source, target int) int {
	if source == target {
		return 0
	}
	g := make(map[int][]int)
	for i, bus := range r {
		for _, route := range bus {
			g[route] = append(g[route], i)
		}
	}
	queue := []int{source}
	visitedStops := make(map[int]bool)
	visitedStops[source] = true
	visitedBuses := make(map[int]bool)
	busTaken := 0

	for len(queue) > 0 {
		n := len(queue)
		busTaken++

		for i := 0; i < n; i++ {
			stop := queue[i]
			for _, bus := range g[stop] {
				if visitedBuses[bus] {
					continue
				}
				visitedBuses[bus] = true
				for _, busStop := range r[bus] {
					if busStop == target {
						return busTaken
					}
					if !visitedStops[busStop] {
						visitedStops[busStop] = true
						queue = append(queue, busStop)
					}

				}
			}

		}
		queue = queue[n:]
	}

	return -1
}

func main() {
	r := [][]int{{2, 3, 4}, {5, 6, 7, 8}, {4, 5, 9, 10}, {10, 12}}

	fmt.Println(busRoutes(r, 3, 12))
}
