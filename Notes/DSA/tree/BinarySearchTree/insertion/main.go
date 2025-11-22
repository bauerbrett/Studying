package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertNode(node *TreeNode, val int) *TreeNode {
	if node == nil { // you need to check if the node they give you is nil. Because at some point we will hit nil and we need to create the node
		return &TreeNode{Val: val}
	}

	if node.Val < val {
		node.Right = insertNode(node.Right, val)
	} else if node.Val > val {
		node.Left = insertNode(node.Left, val)
	} else {
		fmt.Println("Cannot have duplicate node values!")
	}
	return node // you need this because it is used to rebuild the tree on the way back up after the base case nil is hit. Because those other node did not hit nil
	// which means they need to return something to the previous node.left or node.right =. If you didn't have this it would not connect back up
}
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Println(root.Val)
	inorder(root.Right)
}

func main() {
	var root *TreeNode

	root = insertNode(root, 8)
	insertNode(root, 5)
	insertNode(root, 3)
	insertNode(root, 7)
	insertNode(root, 2)
	insertNode(root, 4)
	insertNode(root, 6)
	insertNode(root, 8) // Attempting to insertNode duplicate 8

	fmt.Println("BST Insertion Completed.")
	fmt.Print("Inorder Traversal (Sorted Order): ")
	inorder(root)
	fmt.Println()

}
