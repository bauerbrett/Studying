package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func difference(a, b int) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= 1
}

func checkBalance(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	leftDepth, L := checkBalance(node.Left)
	rightDepth, R := checkBalance(node.Right)

	if !L || !R {
		return 0, false
	}

	/*
		This the cleaner way but I had it rigth originally also. Basically we need to find the depth and the balance.
		We use DFS like we did for max depth problem, to do this we recusively search the left and right side and add 1 on each run of the call.
		So on the bubble up it will add the depth up.

		We then search for the balance also. At every level of the calls we search if left and right are balanced. As it comes back up and the call stack
		it will add the additional depths of the other trees and compare each time their differences. If it ever finds the difference to be more than one
		it will return a false bool and this false bool will be checked by the if !L or !R check. If it ever sees the false it will keep returning false all
		the way up the call stack to the top.

		depth := max(leftDepth, rightDepth) + 1
		isBalanced := difference(leftDepth, rightDepth)
	*/

	if leftDepth > rightDepth {
		return leftDepth + 1, difference(leftDepth, rightDepth)
	}
	if leftDepth < rightDepth {
		return rightDepth + 1, difference(leftDepth, rightDepth)
	}
	return leftDepth + 1, difference(leftDepth, rightDepth)

}

func main() {
	root := &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 20}
	root.Right.Right.Right = &TreeNode{Val: 22}
	root.Right.Right.Right.Right = &TreeNode{Val: 24}
	root.Right.Right.Right.Right.Right = &TreeNode{Val: 26}
	root.Right.Left = &TreeNode{Val: 13}

	root.Left = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 8}
	root.Left.Right.Right = &TreeNode{Val: 9}

	_, check := checkBalance(root)
	if !check {
		fmt.Println("Unbalanced Tree")
	} else {
		fmt.Println("Balanced Tree")
	}
}
