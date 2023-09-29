package dsalgo

//https://leetcode.com/problems/number-of-islands/description/

/*
Suprisingly easy problem if you realize its sort of a flood fill type of algo
Main thing to realize is its a type of dfs algo
What we need to realize is when you encouter a '1' we need to find all connected islands on
four sides which is grid(i+1),grid(i-1),grid(j+1),grid(j-1) once you visit each if its a 1 you mark
it as visited. The break point for dfs is when its not '1' or you pass array bounds

You do this each time a unique '1' is noticed and you fill the other connected '1's and marking
as visited along dfs
*/

func numIslands(grid [][]byte) int {
	noIslands := 0

	rows := len(grid)
	col := len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				mark(i, j, rows, col, grid)
				noIslands++
			}
		}
	}
	return noIslands
}

func mark(i, j, row, col int, grid [][]byte) {
	if i < 0 || i >= row || j < 0 || j >= col || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '*'
	mark(i+1, j, row, col, grid)
	mark(i-1, j, row, col, grid)
	mark(i, j+1, row, col, grid)
	mark(i, j-1, row, col, grid)
}
