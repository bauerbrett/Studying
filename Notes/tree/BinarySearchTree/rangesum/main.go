package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSum(node *TreeNode, a, b int) int {
	var total int
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		postorder(node.Left)
		postorder(node.Right)

		if node.Val >= a && node.Val <= b {
			total += node.Val
		}
	}
	postorder(node)
	return total

}

func rangeSum2(node *TreeNode, a, b int) int {
	if node == nil {
		return 0
	}

	l := rangeSum2(node.Left, a, b)
	r := rangeSum2(node.Right, a, b)

	if node.Val >= a && node.Val <= b {
		return l + r + node.Val
	}
	return l + r
}

func main() {
	root := &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 12}
	root.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 3}
	fmt.Println(rangeSum(root, 5, 10))  // should return 23
	fmt.Println(rangeSum2(root, 5, 10)) // should return 23

}
