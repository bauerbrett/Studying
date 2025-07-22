package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode //we can fit any letter in the alphabet in the slice
	endWord  bool
}
type Solution struct {
	root *TrieNode
}

func (s *Solution) Insert(word string) {
	node := s.root
	for _, c := range word {
		index := c - 'a' // this will basically find index of the letter in the alphabet. We know c is the 3rd letter so the index would 2 for (0, 1, 2)
		if node.children[index] == nil {
			node.children[index] = &TrieNode{}
		}
		node = node.children[index] // move to the next character in the word
	}
	node.endWord = true // if we make it out of for loop that means we are at teh end of word and need to mark the node as end
}

func (s *Solution) Search(word string) bool {
	node := s.root
	for _, c := range word {
		index := c - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.endWord
}

func (s *Solution) StartsWith(word string) bool {
	node := s.root
	for _, c := range word {
		index := c - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true
}

func NewSolution() *Solution {
	return &Solution{root: &TrieNode{}}
}

func main() {

	t := NewSolution()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))
	fmt.Println(t.StartsWith("app"))
	fmt.Println(t.StartsWith("abb"))
	fmt.Println(t.Search("apples"))

}
