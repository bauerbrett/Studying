package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//This is diffent then insert because we dont need to rebuild the tree like in insert where we had to return a node or do node.right = recursiveFunc()
// For this as soon as either search the whole tree or find the key we return the node and this gets return back up to the first call

func search(node *TreeNode, key int) *TreeNode {
	if node == nil || node.Val == key {
		return node
	}
	if key > node.Val {
		return search(node.Right, key)
	}
	return search(node.Left, key)

}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 8}

	key := 68 // Searching for key 6
	result := search(root, key)
	if result != nil {
		fmt.Printf("Key %d found in BST.\n", key)
	} else {
		fmt.Printf("Key %d not found in BST.\n", key)
	}
}
