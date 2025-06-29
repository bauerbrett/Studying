package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
This is a depth first search used to find the max height of bst

Don't try to unfold the recursion. Think about the recursion on a single level. If it works on any individual level then it works overall.

*/

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := maxHeight(root.Left)
	maxRight := maxHeight(root.Right)

	if maxLeft > maxRight {
		return maxLeft + 1
	}
	return maxRight + 1
}

/*
	1
	   \
	    2
	     \
	      3

So finding min is different then max
1. The issue here is if there isnt a subtree on right or left. It needs to be empty on both sides not just one.
So if we do something similar to max we will always return the wrong number because one node migth only have one subtree but since the other side
is empty it would return that value since it is smaller.

To fix this we need to check if one side exist and if both side exist. If both side exist we use the same type of recursion as max and do the left and right min,
but if there is one side nil, then we recursively search that side because we need to find the end where right and left == nil to be a leaf.

2. So check that one side is not nil, if it is search that side to find the bottom which would be base 0
3. If they both are not nil mind min of both sides recursively and return the smaller height.

4. These will build back up until we find the smallest. We don't need variables for the first root.left and root.rigth == nil if statements because the funcs
themselelves will return the ints back.

Note: Don't try to unfold the recursion. Think about the recursion on a single level. If it works on any individual level then it works overall.
*/
func minHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minHeight(root.Right) + 1
	}
	if root.Right == nil {
		return minHeight(root.Left) + 1
	}
	leftMin := minHeight(root.Left)
	rightMin := minHeight(root.Right)

	if leftMin < rightMin {
		return leftMin + 1
	}
	return rightMin + 1

}
func main() {

	// Example 1
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Left = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 5}
	fmt.Println(maxHeight(root1)) // Expected output: 3

	// Example 2
	root2 := &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 2}
	root2.Right.Right = &TreeNode{Val: 3}
	fmt.Println(maxHeight(root2)) // Expected output: 3

	// Example 3
	root3 := &TreeNode{Val: 1}
	root3.Left = &TreeNode{Val: 2}
	root3.Right = &TreeNode{Val: 3}
	root3.Left.Left = &TreeNode{Val: 4}
	root3.Left.Right = &TreeNode{Val: 7}
	root3.Left.Right.Right = &TreeNode{Val: 9}
	fmt.Println(maxHeight(root3)) // Expected output: 4
	fmt.Println(minHeight(root3))
}
