package problem0210

import "container/list"

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	deg := make([]int, numCourses)
	var ans []int
	for _, edge := range prerequisites {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		deg[edge[0]]++
	}

	que := list.New()
	for i, v := range deg {
		if v == 0 {
			que.PushBack(i)
			ans = append(ans, i)
		}
	}

	for que.Len() > 0 {
		front := que.Remove(que.Front()).(int)
		for _, v := range graph[front] {
			deg[v]--
			if deg[v] == 0 {
				que.PushBack(v)
				ans = append(ans, v)
			}
		}
	}

	if len(ans) != numCourses {
		return []int{}
	}
	return ans
}
