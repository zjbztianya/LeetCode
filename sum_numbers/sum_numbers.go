package sum_numbers

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func dfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		ans += sum
	}
	dfs(root.Left, sum)
	dfs(root.Right, sum)
}

func sumNumbers(root *TreeNode) int {
	ans = 0
	dfs(root, 0)
	return ans
}
