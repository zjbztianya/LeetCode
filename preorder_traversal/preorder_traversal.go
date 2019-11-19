package preorder_traversal

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
	*ans = append(*ans, root.Val)
	dfs(root.Left, ans)
	dfs(root.Right, ans)
}

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	dfs(root, &ans)
	return ans
}
