package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// insert
func insert(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}

	if node.Val < val {
		node.Right = insert(node.Right, val)
	} else if node.Val > val {
		node.Left = insert(node.Left, val)
	} else {
		fmt.Println("Key already exist!", val)
	}
	return node
}

// delethelper
func helpDelete(node *TreeNode) *TreeNode {
	current := node
	for current != nil {
		current = node.Left
	}
	return current
}

// delete
func delete(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	if val > node.Val {
		node.Right = delete(node.Right, val)
	} else if val < node.Val {
		node.Left = delete(node.Left, val)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			succession := helpDelete(node.Right)
			node.Val = succession.Val
			node.Right = delete(node.Right, succession.Val)
		}
	}
	return node

}

// search
func search(node *TreeNode, key int) *TreeNode {
	if node == nil || node.Val == key {
		return node
	}
	if key < node.Val {
		return search(node.Left, key)
	}
	return search(node.Right, key)

}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Println(root.Val)
	inOrder(root.Right)
}

// main
func main() {
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 12}

	fmt.Println("First inorder:")
	inOrder(root)

	fmt.Println("After insert 18 and 15")
	root = insert(root, 18)
	root = insert(root, 15)
	inOrder(root)
	fmt.Println("After delete 6")
	root = delete(root, 6)
	inOrder(root)

	fmt.Println("Search for 8")
	//root = delete(root, 8)
	if search(root, 8) == nil {
		fmt.Println("Could not find value in BST: ", 8)
	} else {
		fmt.Println("Value found in BST: ", 8)
	}

}
