package problem0020

type stack []byte

func (s stack) Push(c byte) stack {
	return append(s, c)
}

func (s stack) Pop() (stack, byte) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Len() int {
	return len(s)
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	var st stack
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			st = st.Push(s[i])
		} else {
			if st.Len() == 0 {
				return false
			}
			var top byte
			st, top = st.Pop()
			if s[i] == ')' && top != '(' {
				return false
			}
			if s[i] == ']' && top != '[' {
				return false
			}
			if s[i] == '}' && top != '{' {
				return false
			}
		}
	}
	return st.Len() == 0
}
