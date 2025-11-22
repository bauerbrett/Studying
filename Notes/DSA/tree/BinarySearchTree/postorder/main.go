package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct{}

/*
postOrder
1. get root
2. traverse left
3. traverse right
4. visit node and get value
*/

func (s Solution) postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	s.postOrder(root.Left)
	s.postOrder(root.Right)
	fmt.Println(root.Val)
}

func main() {
	// Create a sample BST:
	//         8
	//       /   \
	//      3     10
	//     / \      \
	//    1   6      14
	//       /
	//      4
	root := &TreeNode{Val: 8}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 10}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 14}

	pt := Solution{}
	fmt.Println("Postorder Traversal (Root Last):")
	pt.postOrder(root) // Expected output: 1 4 6 3 14 10 8
	fmt.Println()
}
