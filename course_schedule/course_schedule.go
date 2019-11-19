package course_schedule

import "container/list"

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	deg := make([]int, numCourses)
	for _, edge := range prerequisites {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		deg[edge[1]]++
	}

	que := list.New()
	for i, v := range deg {
		if v == 0 {
			que.PushBack(i)
		}
	}

	for que.Len() > 0 {
		front := que.Remove(que.Front()).(int)
		for _, v := range graph[front] {
			deg[v]--
			if deg[v] == 0 {
				que.PushBack(v)
			}
		}
	}

	for _, v := range deg {
		if v > 0 {
			return false
		}
	}
	return true
}
