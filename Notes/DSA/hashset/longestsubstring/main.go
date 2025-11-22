package main

import "fmt"

/*
Given a string, identify the length of its longest segment that contains distinct characters. In other words, find the maximum length of a substring that has no repeating characters.

Examples:

Example 1:

Input: "abcdaef"
Expected Output: 6
Justification: The longest segment with distinct characters is "bcdaef", which has a length of 6.

brute force is double for loop
optimize is sliding window that takes the window from start end. Only loops once.
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// "abcdaef"
func slidingWindow(str string) int {
	set := map[byte]bool{}
	start, end, maxLength := 0, 0, 0
	for end < len(str) {
		if !set[str[end]] {
			set[str[end]] = true
			maxLength = max(maxLength, end-start+1)
			end++
		} else {
			delete(set, str[start])
			start++
		}
	}
	return maxLength
}

func longestSubstring(str string) int {
	longest := 0
	for i := 0; i < len(str); i++ {
		set := map[byte]bool{}
		for j := i; j < len(str); j++ {
			if set[str[j]] {
				break
			}
			set[str[j]] = true
		}
		if len(set) > longest {
			longest = len(set)
		}
	}
	return longest
}

func main() {
	str := "abrkaabcdefghijjxxx"
	str1 := "abcdaef"
	fmt.Println(longestSubstring(str), longestSubstring(str1), slidingWindow(str), slidingWindow(str1))
}
