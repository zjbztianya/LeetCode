package build_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val, n := preorder[0], 0
	for inorder[n] != val {
		n++
	}
	node := &TreeNode{val, nil, nil}
	node.Left = makeTree(preorder[1:n+1], inorder[:n])
	node.Right = makeTree(preorder[n+1:], inorder[n+1:])
	return node
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return makeTree(preorder, inorder)
}
