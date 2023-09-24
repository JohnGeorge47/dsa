package dsalgo

/*
Apparently its a topological map i had no clue what it is
We first need to understand that this is a graph problem and any
course which depends on each other means there is a cycle
Also values within prerequisites array  will be courses from 0 to num-1

for eg numCourses = 2, prerequisites = [[1,0]]
Here we can draw the graph like 1->0

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
so here 1-->0 and 0-->1 so there is a cycle

The question now is how do we even represent this because a course
can be dependent on more than one other course so we have a course map
coursemap[course1]=[course2,course3....]
This map will have a course and an array of all the courses it depends on
also there maybe courses which depend on none

Now how do we find cycles we will need to do a dfs (recursion here)
so we will go through every key in coursemap
*/

/*
This is a key to keep track of every visit while doing our dfs note that
if we find a key then we can assume it was visited and hence a cycle exists
*/
var visited map[int]bool

var courseMap map[int][]int

func canFinish(numCourses int, prerequisites [][]int) bool {
	courseMap = make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		courseMap[i] = make([]int, 0)
	}
	//Lets fill the map of courses with dependent course arrays
	for _, prerequisite := range prerequisites {
		courseMap[prerequisite[0]] = append(courseMap[prerequisite[0]], prerequisite[1])
	}
	//Lets do a dfs on every course which needs to be taken even if one returns false we cant
	for k, _ := range courseMap {
		if !checkCourses(k) {
			return false
		}
	}
	return true
}

func checkCourses(courseval int) bool {
	//Incase its not dependent on anything
	if len(courseMap[courseval]) == 0 {
		return true
	}
	//We check if we visited that node/course before so cycle exists
	if _, ok := visited[courseval]; ok {
		return false
	}
	//Mark current node as visited
	visited[courseval] = true

	//Now we visit all the dependent nodes if any are false theres a cycle
	for _, course := range courseMap[courseval] {
		if !checkCourses(course) {
			return false
		}
	}
	//Now we reset visited
	delete(visited, courseval)
	//Incase its a path with no cycles we can mark it as empty since we know this
	//path has no cycles so we need not visit again
	courseMap[courseval] = []int{}
	return true
}
