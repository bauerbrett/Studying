package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// This finds the very last deepest value in a tree. This helper func is used because if the
// node has a left and right child and we delete the node we need to replace it with a value that greater than the left but smaller than the right
// So we go right on the node we are deleting and then go all the way left on its tree. This is value we want to replace the middle node with.
func helpdeleteNode(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

/*
Explain the delete
1. see if the node is nil - return nil if it is
2. if the node val is smaller than the key val then we recursively delete right
3. if the node val is greater than the key val then we recursively delete left
4. like insert we need to rebuild tree so we need to make the nodes = when we do the recursive delete. Whereas search we just return the value and not rebuild tree
5. once match is found we check if left or right exist. if one of the other does we return the other because that will fill the deleted node
6. if both exist we find the smallest node value on the right subtree to replace the node. This will make sure it is smaller than right and larger than left
7. delete the node, we do this by sinply making teh successor val = to node val and then we recursively delete the old successor.val from the right subtree where we found it
8. return node because once if it gets to this part at the end we just changed the val of node to the sucessor val. So it will return with its new value.
*/

func delete(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Val > val {
		node.Left = delete(node.Left, val)
	} else if node.Val < val {
		node.Right = delete(node.Right, val)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			successor := helpdeleteNode(node.Right)
			node.Val = successor.Val
			node.Right = delete(node.Right, successor.Val)
		}
	}
	return node

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

	root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 8}

	fmt.Println("Original BST (Inorder Traversal):")
	inorder(root)
	fmt.Println()

	keyToDelete := 3
	fmt.Printf("\nDeleting node %d...\n", keyToDelete)
	root = delete(root, keyToDelete)

	fmt.Println("BST after deletion (Inorder Traversal):")
	inorder(root)
	fmt.Println()
}
