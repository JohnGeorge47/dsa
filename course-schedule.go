package dsalgo

var visited map[int]bool
var courseMap map[int][]int

func canFinish(numCourses int, prerequisites [][]int) bool {
	courseMap = make(map[int][]int)
	for _, prerequisite := range prerequisites {
		courseMap[prerequisite[1]] = append(courseMap[prerequisite[1]], prerequisite[0])
	}
	for k, _ := range courseMap {
		if !checkCourses(k) {
			return false
		}
	}
	return true
}

func checkCourses(courseval int) bool {
	if len(courseMap[courseval]) == 0 {
		return true
	}
	if _, ok := visited[courseval]; ok {
		return false
	}
	visited[courseval] = true
	for _, course := range courseMap[courseval] {
		if !checkCourses(course) {
			return false
		}
	}
	delete(visited, courseval)
	courseMap[courseval] = []int{}
	return true
}
