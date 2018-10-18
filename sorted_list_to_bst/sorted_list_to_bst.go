package sorted_list_to_bst

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var myhead *ListNode

func buildTree(l int, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	node := new(TreeNode)
	node.Left = buildTree(l, mid-1)
	node.Val = myhead.Val
	myhead = myhead.Next
	node.Right = buildTree(mid+1, r)
	return node

}

func sortedListToBST(head *ListNode) *TreeNode {
	var n int
	for p := head; p != nil; p = p.Next {
		n++
	}
	myhead = head
	return buildTree(0, n-1)
}
