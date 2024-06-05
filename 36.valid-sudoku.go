/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	// 0ならおける、1ならおけない
	boardFlag := make([][][]int, 9)
	for i := range boardFlag {
		boardFlag[i] = make([][]int, 9)
		for j := range boardFlag[i] {
			boardFlag[i][j] = make([]int, 9)
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '1'
			if boardFlag[i][j][num] == 1 {
				return false
			}
			boardFlag = updateBoarFlag(i, j, int(num), boardFlag)
		}
	}
	return true
}

func updateBoarFlag(i, j, num int, boardFlag [][][]int) [][][]int {
	for k := 0; k < 9; k++ {
		boardFlag[i][j][k] = 1
		boardFlag[k][j][num] = 1
		boardFlag[i][k][num] = 1
	}
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			boardFlag[(i/3)*3+k][(j/3)*3+l][num] = 1
		}
	}
	return boardFlag
}

// @lc code=end
