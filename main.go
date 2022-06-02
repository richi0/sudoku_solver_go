package main

import (
	"fmt"
)

type Row [9]int
type Board [9]Row

var solution *Board

func HasDublicate(row *Row) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if i != j && row[j] != 0 && row[i] == row[j] {
				return true
			}
		}
	}
	return false
}

func getYRow(board *Board, index int) *Row {
	row := Row{}
	for i, r := range board {
		row[i] = r[index]
	}
	return &row
}

func getSquare(board *Board, index int) *Row {
	row := Row{}
	keys := [9][2]int{
		{0, 0}, {0, 3}, {0, 6},
		{3, 0}, {3, 3}, {3, 6},
		{6, 0}, {6, 3}, {6, 6},
	}
	key := keys[index]
	counter := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			row[counter] = board[key[0]+i][key[1]+j]
			counter++
		}
	}
	return &row
}

func CheckBoard(board *Board) bool {
	for i, row := range board {
		if HasDublicate(&row) ||
			HasDublicate(getYRow(board, i)) ||
			HasDublicate(getSquare(board, i)) {
			return false
		}
	}
	return true
}

func getNextZero(board *Board) (int, int) {
	for i, row := range board {
		for j, n := range row {
			if n == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func solve(board *Board) {
	if solution != nil {
		return
	}
	i, j := getNextZero(board)
	if i == -1 && CheckBoard(board) {
		solution = board
		return
	}
	for k := 1; k <= 9; k++ {
		board[i][j] = k
		if CheckBoard(board) {
			solve(board)
			if solution != nil {
				return
			}
		}
	}
	board[i][j] = 0
}

func Run(board *Board) *Board {
	solution = nil
	solve(board)
	return solution
}

func main() {
	b := Board{
		Row{0, 0, 0, 0, 0, 0, 0, 0, 4},
		Row{5, 0, 0, 0, 8, 0, 0, 1, 7},
		Row{0, 2, 0, 5, 0, 0, 0, 0, 0},
		Row{0, 0, 6, 0, 0, 9, 0, 0, 0},
		Row{2, 0, 0, 0, 7, 0, 0, 8, 5},
		Row{0, 0, 0, 0, 0, 0, 3, 0, 0},
		Row{0, 0, 0, 0, 3, 0, 2, 0, 0},
		Row{1, 0, 0, 0, 0, 0, 4, 0, 0},
		Row{0, 9, 0, 0, 0, 7, 0, 3, 1},
	}
	solve(&b)
	if solution != nil {
		fmt.Println("Found solution:")
		for _, row := range solution {
			fmt.Println(row)
		}
	} else {
		fmt.Println("No soultion found")
	}
}
