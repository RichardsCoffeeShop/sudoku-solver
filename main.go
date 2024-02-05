package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] // Get's all lines of sudoku from command line

	if !isInputValid(args) {
		fmt.Println("Error")
		return
	}

	board := parseSudoku(args)

	if !isBoardValid(board) {
		fmt.Println("Error")
		return
	}

	solveAndPrintSudoku(board)
}

func solveAndPrintSudoku(board [9][9]int) {
	resolveSudoku(board)
}

func resolveSudoku(board [9][9]int) bool {
	// If there are no empty cells, the sudoku is solved
	if !hasEmptyCell(board) {
		printSudoku(board)
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if resolveSudoku(board) {
							return true
						}
					}
					board[i][j] = 0
				}
				return false
			}
		}
	}

	return false
}

func hasEmptyCell(board [9][9]int) bool {
	for _, row := range board {
		for _, value := range row {
			if value == 0 {
				return true
			}
		}
	}

	return false
}

func printSudoku(board [9][9]int) {
	for _, row := range board {
		for _, value := range row {
			fmt.Printf("%v ", value)
		}
		fmt.Println()
	}
}

func parseSudoku(args []string) [9][9]int {
	board := [9][9]int{}

	if len(args) == 1 {
		args = strings.Split(args[0], " ")
	}

	if len(args) == 9 {
		for i, arg := range args {
			for j, char := range arg {
				if char == '.' {
					board[i][j] = 0
				} else {
					board[i][j] = int(char - '0')
				}
			}
		}
	}

	return board
}

func isBoardValid(board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		if hasDuplicates(getRow(board, i)) {
			return false
		}
		if hasDuplicates(getColumn(board, i)) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if hasDuplicates(getSquare(board, i, j)) {
				return false
			}
		}
	}

	return true
}

func getRow(board [9][9]int, row int) [9]int {
	return board[row]
}

func getColumn(board [9][9]int, column int) [9]int {
	var result [9]int

	for i := 0; i < 9; i++ {
		result[i] = board[i][column]
	}

	return result
}

func getSquare(board [9][9]int, row int, column int) [9]int {
	var result [9]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i*3+j] = board[row+i][column+j]
		}
	}

	return result
}

func hasDuplicates(line [9]int) bool {
	seen := make(map[int]bool)

	for _, value := range line {
		if value != 0 {
			if seen[value] {
				return true
			}
			seen[value] = true
		}
	}

	return false
}

func isInputValid(args []string) bool {
	if len(args) == 1 {
		args = strings.Split(args[0], " ")
	}

	if len(args) != 9 {
		return false
	}

	for _, arg := range args {
		if len(arg) != 9 {
			return false
		}
	}

	return true
}
