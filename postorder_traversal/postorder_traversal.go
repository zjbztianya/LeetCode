package postorder_traversal

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, ans)
	dfs(root.Right, ans)
	*ans = append(*ans, root.Val)
}

func postorderTraversal(root *TreeNode) (res []int) {
	dfs(root, &res)
	return res
}
