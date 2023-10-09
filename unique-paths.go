package dsalgo

// https://leetcode.com/problems/unique-paths/
func uniquePaths(m int, n int) int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	matrix[m][n] = 1
	for i := len(matrix) - 1; i > 0; i-- {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			if j == n-1 || i == m-1 {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = matrix[i+1][j] + matrix[i][j+1]
			}
		}
	}
	return matrix[0][0]
}

// Brute force solution dosent work for large btw
func path(i, j, m, n int) int {
	if i < 0 || j < 0 || i > m || j > m {
		return 0
	}
	if i == m && j == n {
		return 1
	}
	return path(i+1, j, m, n) + path(i, j+1, m, n)
}

//Dp bottom up is the solution kind of easy to visualize if you draw it
//Its like to reach a point we can either come from below or from right so
//dp[m][n] is 1 since to reach that point its one step now anything in last
// lines m and n is 1 since only one way to reach the point from right and below respectfully
// So for dp[i][j]=dp[i+1][j]+dp[i][j+1]
