package main

import (
	"fmt"
	"sort"
)

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Solution struct {
	root *TrieNode
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

func (s *Solution) IndexPairs(text string, words []string) [][]int {
	for _, word := range words {
		s.Insert(word)
	}

	res := [][]int{}
	for i := 0; i < len(text); i++ { //first loop gest teh first char of word
		node := s.root
		for j := i; j < len(text); j++ { //second loop starts at first char and extends as long as trie ahs children or we reach end of loop
			index := text[j] - 'a'
			if node.children[index] == nil {
				break
			}
			node = node.children[index]
			if node.isEnd {
				res = append(res, []int{i, j})
			}
		}
	}

	// Sort by start index, then end index
	sort.Slice(res, func(i, j int) bool {
		if res[i][0] == res[j][0] {
			return res[i][1] < res[j][1]
		}
		return res[i][0] < res[j][0]
	})

	return res
}
func Constructor() *Solution {
	return &Solution{root: &TrieNode{}}
}

func main() {
	s := Constructor()
	s.Insert("programmingisfun")
	str := "programmingisfun"
	//fmt.Println(s.root.children[1])
	l := []string{"pro", "is", "fun", "gram"}
	fmt.Println(s.IndexPairs(str, l))

}
