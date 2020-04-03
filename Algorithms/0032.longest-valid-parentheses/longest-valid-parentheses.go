package problem0032

type Stack []int

func (s Stack) Push(c int) Stack {
	return append(s, c)
}

func (s Stack) Pop() (Stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s Stack) Len() int {
	return len(s)
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestValidParentheses(s string) int {
	var stack Stack
	a := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = stack.Push(i)
		} else if stack.Len() > 0 {
			var top int
			stack, top = stack.Pop()
			a[i], a[top] = true, true
		}
	}
	var ans, cnt int
	for i := 0; i < len(s); i++ {
		if a[i] {
			cnt++
		} else {
			ans = Max(ans, cnt)
			cnt = 0
		}
	}
	ans = Max(ans, cnt)
	return ans
}
