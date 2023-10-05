package dsalgo

//https://leetcode.com/problems/spiral-matrix/

func spiralOrder(matrix [][]int) []int {
	resultArr := make([]int, 0)
	spiralDfs(matrix, &resultArr, 0, 0, "right")
	return resultArr
}

// Brute force yay!!!
/*
Saw a one line python solution but it was too much for poor me
*/
func spiralDfs(matrix [][]int, result *[]int, i, j int, traversal string) {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) || matrix[i][j] == 10000 {
		return
	}
	if traversal == "right" {
		for j < len(matrix[i]) && matrix[i][j] != 10000 {
			*result = append(*result, matrix[i][j])
			matrix[i][j] = 10000
			j = j + 1
		}
		spiralDfs(matrix, result, i+1, j-1, "down")
	}
	if traversal == "down" {
		for i < len(matrix) && matrix[i][j] != 10000 {
			*result = append(*result, matrix[i][j])
			matrix[i][j] = 10000
			i++
		}
		spiralDfs(matrix, result, i-1, j-1, "left")
	}
	if traversal == "left" {
		for j >= 0 && matrix[i][j] != 10000 {
			*result = append(*result, matrix[i][j])
			matrix[i][j] = 10000
			j--
		}
		spiralDfs(matrix, result, i-1, j+1, "up")
	}
	if traversal == "up" {
		for i >= 0 && matrix[i][j] != 10000 {
			*result = append(*result, matrix[i][j])
			matrix[i][j] = 10000
			i--
		}
		spiralDfs(matrix, result, i+1, j+1, "right")
	}
}
