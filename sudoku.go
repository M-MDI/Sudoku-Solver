package main

import (
	"fmt"
	"os"
)

func isValid_input(args []string) bool {
	if len(args) != 10 {
		return false
	}

	for i := 1; i <= 9; i++ {
		if len(args[i]) != 9 {
			return false
		}
	}

	return true
}

func isValid_char(char byte) bool {
	return char == '.' || (char >= '1' && char <= '9')
}

func isValid_boardFormat(board []string) bool {
	for _, row := range board {
		for _, char := range row {
			if !isValid_char(byte(char)) {
				return false
			}
		}
	}
	return true
}

func isValid_nbr(char byte) bool {
	return char >= '1' && char <= '9'
}

func isBoardValid(board [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 && !isValid_nbr(byte(board[i][j]+'0')) {
				return false
			}
		}
	}
	return true
}

func isValid_inputFormat(board []string) bool {
	return isValid_boardFormat(board) && isBoardValid(check_Board(board))
}

func isValid(board [][]int, row, col, nbr int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == nbr || board[i][col] == nbr || board[3*(row/3)+i/3][3*(col/3)+i%3] == nbr {
			return false
		}
	}
	return true
}

func solve_board(board [][]int) bool {
	empty := empty_Cell(board)

	if empty == nil {
		return true
	}

	row, col := empty[0], empty[1]

	for nbr := 1; nbr <= 9; nbr++ {
		if isValid(board, row, col, nbr) {
			board[row][col] = nbr
			if solve_board(board) {
				return true
			}
			// backtrack => If placing the current number doesn't lead to a solution
			board[row][col] = 0
		}
	}
	// backtrack => No valid number was found for the current cell
	return false
}

func empty_Cell(board [][]int) []int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return []int{i, j}
			}
		}
	}
	return nil
}

func print_Board(board [][]int) {
	for _, row := range board {
		for _, nbr := range row {
			fmt.Printf("%d ", nbr)
		}
		fmt.Println()
	}
}

func check_Board(board []string) [][]int {
	result := make([][]int, 9)
	for i, row := range board {
		result[i] = make([]int, 9)
		for j, char := range row {
			if char != '.' {
				result[i][j] = int(char - '0')
			}
		}
	}
	return result
}

func validation(grid [][]int) bool {
	for _, row := range grid {
		res := 0
		for _, rowNum := range row {
			res += rowNum
		}
		if res != 45 {
			return false
		}

	}
	return true
}

func main() {
	if !isValid_input(os.Args) {
		fmt.Println("Error")
		return
	}

	sudokuBoard := check_Board(os.Args[1:])
	if !isValid_inputFormat(os.Args[1:]) {
		fmt.Println("Error")
		return
	}

	if solve_board(sudokuBoard) {
		print_Board(sudokuBoard)
	} else {
		fmt.Println("Error")
	}
}
