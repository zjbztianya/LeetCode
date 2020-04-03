package problem0179

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	a := make([]string, len(nums))
	for i, v := range nums {
		a[i] = strconv.Itoa(v)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i]+a[j] > a[j]+a[i]
	})

	s := strings.Join(a, "")
	if len(s) > 1 && s[0] == '0' {
		return "0"
	}
	return s
}
