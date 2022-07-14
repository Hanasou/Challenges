package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return dfs(p, q)
}

// Perform dfs on both of the trees to see if they're the same
func dfs(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return dfs(p.Left, q.Left) && dfs(p.Right, q.Right)

}
