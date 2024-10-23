package tree

import "sandbox/problems"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeDfs(p, q)
}

// Perform dfs on both of the trees to see if they're the same
func isSameTreeDfs(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTreeDfs(p.Left, q.Left) && isSameTreeDfs(p.Right, q.Right)

}

func isValidBST(root *TreeNode) bool {
	// Take in two additional arguments
	// We want to make sure all the elements in the left subtree are less than this element
	// Make sure all the elements in the right subtree are greater than this element
	return isValidBstDfs(root, root.Val, root.Val)
}

func isValidBstDfs(root *TreeNode, left int, right int) bool {
	if root == nil {
		return true
	}
	// Check values in the left subtree are less than the value of root
	if root.Left != nil && (root.Left.Val < left || root.Left.Val > root.Val) {
		return false
	}
	if root.Right != nil && (root.Right.Val > right || root.Right.Val < root.Val) {
		return false
	}
	// Check values in the right subtree are greater than the value of root
	return isValidBstDfs(root.Left, left, right) && isValidBstDfs(root.Right, left, right)
}

func isValidBstDfsSolution(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val > max && root.Val < min {
		return false
	}
	// What's the logic behind setting the boundaries this way?
	// When we search the left subtree, we'll be comparing it to the max
	// When we search the right subtree, we'll be comparing it to the min
	return isValidBstDfsSolution(root.Left, min, root.Val) && isValidBstDfsSolution(root.Right, root.Val, max)
}

// Given two integer arrays preorder and inorder
// where each array contains the respective order of a tree
// Construct and return the binary tree they both represent
func buildTree(preorder []int, inorder []int) *TreeNode {
	// If preorder and inorder are empty, just return nil
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// The root is always going to be the first element of preorder
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	// Find the index of root node in the inorder array
	mid := problems.IndexOfArray(inorder, root.Val)
	// Recursively build the trees.
	// We know that all elements left of mid are in the left subtree
	// We know that all elements right of mid are in the right subtree

	// Build left subtree
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	// Build right subtree
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}
