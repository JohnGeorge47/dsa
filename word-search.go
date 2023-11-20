package dsalgo

//https://leetcode.com/problems/word-search/

/*
Backtracking wohoo
*/

type Visited struct {
	I int
	J int
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			exists := false
			traverseBoard(0, 0, "", word, board, &exists)
			if exists {
				return true
			}
		}
	}
	return false
}

func traverseBoard(i, j int, currentWord, word string, board [][]byte, exists *bool) {
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 {
		return
	}
	if board[i][j] == '*' {
		return
	}
	currentWord = currentWord + string(board[i][j])
	if currentWord == word {
		*exists = true
		return
	}
	char := board[i][j]
	board[i][j] = '*'
	traverseBoard(i+1, j, currentWord, word, board, exists)
	traverseBoard(i-1, j, currentWord, word, board, exists)
	traverseBoard(i, j-1, currentWord, word, board, exists)
	traverseBoard(i, j-1, currentWord, word, board, exists)
	board[i][j] = char
}
