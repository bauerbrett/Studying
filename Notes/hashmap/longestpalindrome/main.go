package main

import "fmt"

// count letters in string and put in map. Add all evens up add odds but subtrack -1 from them. Add all up and this is longest.
// if the total length is odd than add one char for a middle character
func longestPalindrome(p string) int {
	charCount := map[byte]int{}
	for i := 0; i < len(p); i++ {
		charCount[p[i]]++
	}
	count := 0
	oddLength := false
	for _, value := range charCount {
		if value%2 == 0 {
			count += value
		}
		if value%2 != 0 {
			count += value - 1
			oddLength = true
		}
	}
	if oddLength {
		count++
	}
	return count
}

func main() {
	pali := "bananas"
	fmt.Println(longestPalindrome(pali))

}
