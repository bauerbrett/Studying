package main

import "fmt"

func jewelStones(j, s string) int {
	jewelSet := map[byte]bool{}
	for i := 0; i < len(j); i++ {
		jewelSet[j[i]] = true
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if jewelSet[s[i]] {
			count++
		}
	}
	return count

}
func main() {
	jewels := "aA"
	stones := "aAaZz"
	fmt.Println(jewelStones(jewels, stones))
}
