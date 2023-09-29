package dsalgo

//https://leetcode.com/problems/rotting-oranges/

/*
I tried dfs first but it turned out too be too complicated since in dfs we start
from one rotten and travel everywhere  but this leads to issues since we can have multiple
rotten oranges disconnected from each other and the time which it will start at is same for all
oranges regardless
*/

func orangesRotting(grid [][]int) int {
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var rotten [][2]int
	fresh := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				rotten = append(rotten, [2]int{i, j})
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	timeElapsed := 0
	for len(rotten) != 0 {
		for _, oran := range rotten {
			for _, direction := range directions {
				x, y := direction[0]+oran[0], direction[1]+oran[1]
				if x >= 0 || x < len(grid) || y >= 0 || y <= len(grid[0]) || grid[x][y] == 1 {
					grid[x][y] = 2
					rotten = append(rotten, [2]int{x, y})
					fresh--
				}
			}
			rotten = rotten[1:]
		}
		timeElapsed++
	}
	if fresh == 0 {
		return timeElapsed
	}
	return -1
}
