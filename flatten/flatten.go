package flatten

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre *TreeNode

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	right := root.Right
	if pre != nil {
		pre.Left = nil
		pre.Right = root
	}
	pre = root
	dfs(root.Left)
	dfs(right)
}

func flatten(root *TreeNode) {
	pre = nil
	dfs(root)
}
