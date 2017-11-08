package maximum_length_of_pair_chain

import "sort"

type Seg struct {
	l int
	r int
}

type segList []Seg

func (s segList) Len() int           { return len(s) }
func  Less(i, j int) bool { return s[i].r < s[j].r }
func (s segList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func findLongestChain(pairs [][]int) int {
	var p segList
	n := len(pairs)
	for i := 0; i < n; i++ {
		p = append(p, Seg{pairs[i][0], pairs[i][1]})
	}
	sort.Sort(p)
	r := p[0].r
	ans := 1
	for i := 1; i < n; i++ {
		if p[i].l > r {
			r = p[i].r
			ans++
		}
	}
	return ans
}
