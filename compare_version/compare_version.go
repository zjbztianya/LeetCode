package compare_version

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	n1, n2 := len(v1), len(v2)
	n := n1

	if n1 < n2 {
		for i := 0; i < n2-n1; i++ {
			v1 = append(v1, "0")
		}
		n = n2
	} else if n1 > n2 {
		for i := 0; i < n1-n2; i++ {
			v2 = append(v2, "0")
		}
	}

	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(v1[i])
		b, _ := strconv.Atoi(v2[i])
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
	}
	return 0
}
