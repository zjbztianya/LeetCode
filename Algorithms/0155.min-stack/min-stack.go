package problem0155

type MinStack struct {
	s1 []int
	s2 []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.s1 = append(this.s1, x)
	if len(this.s2) == 0 || this.s2[len(this.s2)-1] >= x {
		this.s2 = append(this.s2, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.s1) == 0 {
		return
	}
	s1, s2 := this.s1, this.s2
	top := s1[len(s1)-1]
	s1 = s1[:len(s1)-1]

	if top == s2[len(s2)-1] {
		s2 = s2[:len(s2)-1]
	}
	this.s1, this.s2 = s1, s2
}

func (this *MinStack) Top() int {
	return this.s1[len(this.s1)-1]
}

func (this *MinStack) GetMin() int {
	return this.s2[len(this.s2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
