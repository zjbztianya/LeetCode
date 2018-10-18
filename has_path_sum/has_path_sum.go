package has_path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return (sum - root.Val) == 0
	}

	l := hasPathSum(root.Left, sum-root.Val)
	r := hasPathSum(root.Right, sum-root.Val)
	return l || r
}
