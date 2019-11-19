package isomorphic_strings

func check(s string, t string) bool {
	ms := make(map[byte]byte)
	for i := range s {
		if v, ok := ms[s[i]]; ok && v != t[i] {
			return false
		} else {
			ms[s[i]] = t[i]
		}
	}
	return true
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return check(s, t) && check(t, s)
}
