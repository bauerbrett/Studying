package main

import (
	"fmt"
)

/*
If two nodes from the given trees share the same position, their values should be summed up in the resulting tree.
If a node exists in one tree but not in the other, the resulting tree should have a node at the same position with the value from the existing node.

Examples

Example 1:

Trees:

Tree 1:      1            Tree 2:       1
           /   \                      /   \
          3     2                    2     3

Merged:       2
            /   \
           5     5
Justification: The root nodes of both trees have the value 1, so the merged tree's root has a value of 1 + 1 = 2. For the left child, 3 + 2 = 5 and
for the right child, 2 + 3 = 5.
*/

//Lets create insert func. We can use this to build the tree off the new values returned from visiting the nodes.

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else if val > root.Val {
		root.Right = insert(root.Right, val)
	} else {
		fmt.Println("Value already exist!", root.Val)
	}
	return root
}

func mergeTrees(root1, root2 *TreeNode) *TreeNode {
	var list1 []*TreeNode
	var merge func(*TreeNode, *TreeNode)

	merge = func(root1, root2 *TreeNode) {
		if root1 == nil && root2 == nil {
			return
		}
		val := 0
		if root1 == nil {
			val = root2.Val
		} else if root2 == nil {
			val = root1.Val
		} else {
			val = root1.Val + root2.Val
		}
		list1 = append(list1, &TreeNode{Val: val})
		merge(root1.Left, root2.Left)
		merge(root1.Right, root2.Right)
	}
	merge(root1, root2)
	for _, value := range list1 {
		fmt.Println(value.Val)
		//insert(list1[0], value.Val)
	}
	return list1[0]
}
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Println(root.Val)
	inorder(root.Right)

}

func correctWay(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil { // if either is nil just return the other
		return t2
	}
	if t2 == nil {
		return t1
	}

	merged := &TreeNode{Val: t1.Val + t2.Val}     // if htey both exist, build the new node because we need a new node with the new value
	merged.Left = correctWay(t1.Left, t2.Left)    //go left for both
	merged.Right = correctWay(t1.Right, t2.Right) // go right for both

	return merged //return the merged one. This will build the tree because it will start at first node, and then as the recursion goes deepeer it will build the sub
	// trees and return them up and attach to first
}

func main() {
	// Tree 1:      1
	//             / \
	//            3   2
	tree1 := &TreeNode{Val: 1}
	tree1.Left = &TreeNode{Val: 3}
	tree1.Right = &TreeNode{Val: 2}

	// Tree 2:      1
	//
	//	 / \
	//	2   3
	tree2 := &TreeNode{Val: 1}
	tree2.Left = &TreeNode{Val: 2}
	tree2.Right = &TreeNode{Val: 3}

	root := mergeTrees(tree1, tree2)
	//inorder(root)
	fmt.Println(root, "no")

	rootCorrect := correctWay(tree1, tree2)
	inorder(rootCorrect)

}
