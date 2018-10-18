package min_depth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dl, dr := minDepth(root.Left), minDepth(root.Right)
	if dl == 0 {
		return dr + 1
	}
	if dr == 0 {
		return dl + 1
	}
	return min(dl, dr) + 1
}
