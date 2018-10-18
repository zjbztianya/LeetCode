package sorted_array_to_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	mid := (n - 1) >> 1
	node := &TreeNode{nums[mid],
		sortedArrayToBST(nums[:mid]),
		sortedArrayToBST(nums[mid+1:])}
	return node
}
