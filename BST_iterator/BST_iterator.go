package BST_iterator

//173. Binary Search Tree Iterator

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	values []*TreeNode
}

func (s *stack) Push(c *TreeNode) {
	s.values = append(s.values, c)
}

func (s *stack) Pop() *TreeNode {
	value := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return value
}

func (s *stack) Len() int {
	return len(s.values)
}

type BSTIterator struct {
	values []int
}

func Constructor(root *TreeNode) BSTIterator {
	s := new(stack)
	iter := new(BSTIterator)
	p := root
	for p != nil || s.Len() > 0 {
		for p != nil {
			s.Push(p)
			p = p.Left
		}
		if s.Len() > 0 {
			top := s.Pop()
			iter.values = append(iter.values, top.Val)
			p = top.Right
		}
	}
	return *iter
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	value := this.values[0]
	this.values = this.values[1:len(this.values)]
	return value
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.values) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
