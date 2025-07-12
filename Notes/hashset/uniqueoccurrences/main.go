package main

import "fmt"

func numOccurrences(arr []int) bool {
	countMap := map[int]int{}
	set := map[int]bool{}
	for _, val := range arr {
		countMap[val]++
	}
	unique := true
	for _, v := range countMap {
		if set[v] {
			unique = false
		}
		set[v] = true
	}
	return unique
}
func main() {
	arr := []int{11, 12, 13, 14, 14, 15, 15, 15}
	fmt.Println(numOccurrences(arr))
}
