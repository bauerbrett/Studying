package main

import "fmt"

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

func (s *Solution) IndexPairs(str string, l []string) [][]int {
	result := [][]int{}
	for _, word := range l {
		s.Insert(word)
	}
	for i := 0; i < len(str); i++ { //first loop finds start of word
		node := s.root                  //Every time we go up we still want to start at beginning of trie
		for j := i; j < len(str); j++ { // This will go to teh end of the word.
			index := str[j] - 'a'
			if node.children[index] == nil {
				break
			}
			node = node.children[index] //since root = null in trie we always check the children. Even with 0 we would check child to get the first index.
			if node.isEnd {             //same thing when we get to end, we check child and then return the index we are on. this is why we check child and not current node for isEnd.
				pair := []int{i, j}
				result = append(result, pair)
			}
		}
	}
	return result
}

func main() {
	s := Constructor()
	str := "bluebirdskyscraper"
	l := []string{"blue", "bird", "sky"}
	fmt.Println(s.IndexPairs(str, l))
}
