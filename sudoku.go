package main

import (
	"fmt"
	"os"
)

func solve(board *[9][9]int) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == 0 {
				for n := 1; n <= 9; n++ {
					if isValid(board, r, c, n) {
						board[r][c] = n
						if solve(board) {
							return true
						}
						board[r][c] = 0
					}
				}
				return false
			}
		}
	}
	return true

}
func isValid(board *[9][9]int, r, c, n int) bool {
	for i := 0; i < 9; i++ {
		if board[r][i] == n || board[i][c] == n || board[(r/3)*3+i/3][(c/3)*3+i%3] == n {
			return false
		}
	}
	return true
}
func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	var board [9][9]int
	for i, s := range os.Args[1:] {
		for j, ch := range s {
			if ch == '.' {
				continue
			} else if ch >= '1' && ch <= '9' {
				board[i][j] = int(ch - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}
	if solve(&board) {
		for _, row := range board {
			for _, n := range row {
				fmt.Printf("%d ", n)
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Error")
	}
}
