package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func minDifference(node *TreeNode) int {
	queue := []*TreeNode{node}
	visited := map[*TreeNode]int{}
	small := math.MaxInt32

	for len(queue) > 0 {
		currentNode := queue[0]
		visited[currentNode] = currentNode.Val
		queue = append(queue, currentNode.Left)
		queue = append(queue, currentNode.Right)
		for _, value := range visited {
			diff := difference(currentNode.Val, value)
			if diff < small {
				small = diff
			}
		}
		queue = queue[1:]

	}
	return small
}
*/

func minDifference1(node *TreeNode) int {
	minDiff := math.MaxInt32
	var prev *TreeNode
	inorder(node, &minDiff, &prev)
	return minDiff
}
func difference(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b

}

func inorder(node *TreeNode, minDiff *int, prev **TreeNode) {
	if node == nil {
		return
	}

	inorder(node.Left, minDiff, prev)
	if *prev != nil {
		diff := difference((*prev).Val, node.Val)
		if diff < *minDiff {
			*minDiff = diff
		}
	}
	*prev = node
	inorder(node.Right, minDiff, prev)
}

func main() {
	// Example 1// This is a valid BST
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	fmt.Println(minDifference1(root))

}
