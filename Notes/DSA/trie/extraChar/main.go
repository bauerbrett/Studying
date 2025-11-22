package main

import "fmt"

/*
Given a string s and an array of words words. Break string s into multiple non-overlapping substrings such that
each substring should be part of the words. There are some characters left which are not part of any substring.

Return the minimum number of remaining characters in s, which are not part of any substring after string break-up.

Examples
Example 1:

Input: s = "amazingracecar", words = ["race", "car"]
Expected Output: 7
Justification: The string s can be rearranged to form "racecar", leaving 'a', 'm', 'a', 'z', 'i', 'n', 'g' as extra.

Plan: Return the index of words in the trie. Once we find index

*/

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Solution struct {
	root *TrieNode
}

func Constructor() *Solution {
	return &Solution{root: &TrieNode{}}
}

func (s *Solution) Insert(word string) {
	node := s.root
	for _, c := range word {
		index := c - 'a'
		if node.children[index] == nil {
			node.children[index] = &TrieNode{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

func (s *Solution) extraChar(str string, l []string) int {
	count := 0
	result := [][]int{}
	for _, word := range l {
		s.Insert(word)
	}

	for i := 0; i < len(str); i++ {
		node := s.root
		for j := i; j < len(str); j++ {
			index := str[j] - 'a'
			if node.children[index] == nil {
				break
			}
			node = node.children[index]
			if node.isEnd {
				result = append(result, []int{i, j})
			}
		}
	}

	for _, l := range result {
		for i := l[0]; i <= l[1]; i++ {
			count++
		}
	}
	return len(str) - count
}

func main() {
	s := Constructor()

	str := "thedoggybarksatnight"
	l := []string{"dog", "bark", "night"}
	fmt.Println(s.extraChar(str, l))

}
