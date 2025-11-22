package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given a root node of the Binary Search Tree (BST) and integer 'k'. Return the Kth smallest element among all node values of the binary tree.

Examples:

Example 1:
Input:

	   8
	  / \
	 3   10
	/ \    \

1   6    14

	 /  \  /
	4   7  13

k = 4
Expected Output: 6
Justification: The in-order traversal of the tree is [1, 3, 4, 6, 7, 8, 10, 13, 14]. The 4th element in this sequence is 6.

I want to just do inorder traversal and append to slice and let select slice[k-1]

The other way is to keep a count var and increment by 1 every time we visit a node. Inorder travseral visit is second, so left, visit, right.
So we need to find base case which is when we hit 0 or we hit our count. After we start traversing and visit the nodes. Once we visit we put our logic in
a add 1 to count and check if count == k. If it is we know the visited node is the result so we make our result == k.
*/
func kthSmallest1(node *TreeNode, k int) int {
	var inorder func(*TreeNode)
	var s []int
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		s = append(s, node.Val)
		inorder(node.Right)
	}
	inorder(node)
	return s[k-1]
}

var count int
var result int

func kthSmallest(node *TreeNode, k int) {
	if node == nil || count >= k {
		return
	}
	kthSmallest(node.Left, k)

	count++

	if count == k {
		result = node.Val
	}
	kthSmallest(node.Right, k)

}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 6}
	fmt.Println(kthSmallest1(root, 3))

	kthSmallest(root, 3)
	fmt.Println(result)

}
