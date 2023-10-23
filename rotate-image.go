package dsalgo

import "fmt"

// Saw an amazing solution to this
// We reverse vertically(horizontally in case of anticlockwise) then we do swap of i,j values
func rotate(matrix [][]int) {
	for j := 0; j < len(matrix); j++ {
		start := 0
		end := len(matrix) - 1
		for start < end {
			fmt.Println(matrix[start][j], matrix[end][j])
			matrix[start][j], matrix[end][j] = matrix[end][j], matrix[start][j]
			start++
			end--
		}
	}
	for i := 0; i < len(matrix); i++ {
		//i+1 because we dont want to repeat swaps again
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func reverseArray(row []int) []int {
	start := 0
	end := len(row) - 1
	for start < end {
		row[start], row[end] = row[end], row[start]
		start++
		end--
	}
	return row
}
