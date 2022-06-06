package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type Sudoku struct {
	Id       int
	Mission  string
	Solution string
}

func TestHasDublicate(t *testing.T) {
	tt := []struct {
		name string
		arr  *Row
		res  bool
	}{
		{name: "all zero", arr: &Row{0, 0, 0, 0, 0, 0, 0, 0, 0}, res: false},
		{name: "some zero", arr: &Row{1, 1, 0, 1, 1, 0, 1, 0, 1}, res: true},
		{name: "one zero", arr: &Row{1, 2, 3, 4, 5, 6, 7, 0, 8}, res: false},
		{name: "all dublicate", arr: &Row{1, 1, 1, 1, 1, 1, 1, 1, 1}, res: true},
		{name: "some dublicate", arr: &Row{1, 1, 5, 7, 6, 9, 8, 1, 2}, res: true},
		{name: "one dublicate", arr: &Row{9, 8, 2, 3, 4, 2, 7, 6, 1}, res: true},
		{name: "good order", arr: &Row{1, 2, 3, 4, 5, 6, 7, 8, 9}, res: false},
		{name: "good no order", arr: &Row{6, 4, 8, 9, 1, 2, 5, 3, 7}, res: false},
		{name: "good reverse order", arr: &Row{9, 8, 7, 6, 5, 4, 3, 2, 1}, res: false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.res != HasDublicate(tc.arr) {
				t.Errorf("Got %t, should be %t", HasDublicate(tc.arr), tc.res)
			}
		})
	}
}

func TestBoardCheck(t *testing.T) {
	tt := []struct {
		name  string
		board Board
		res   bool
	}{
		{name: "all zero", board: Board{
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, res: true},
		{name: "one row good", board: Board{
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{2, 1, 3, 4, 5, 6, 7, 8, 9},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, res: true},
		{name: "all dublicate", board: Board{
			Row{1, 1, 1, 1, 1, 1, 1, 1, 1},
			Row{1, 1, 1, 1, 1, 1, 1, 1, 1},
			Row{1, 1, 1, 1, 1, 1, 1, 1, 1},
			Row{1, 1, 1, 1, 1, 1, 1, 1, 1},
			Row{1, 1, 1, 1, 1, 1, 1, 1, 1},
			Row{1, 1, 1, 1, 1, 1, 1, 1, 1},
			Row{1, 1, 1, 1, 1, 1, 1, 1, 1},
			Row{1, 1, 1, 1, 1, 1, 1, 1, 1},
			Row{1, 1, 1, 1, 1, 1, 1, 1, 1},
		}, res: false},
		{name: "some dublicate", board: Board{
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{1, 1, 1, 1, 1, 1, 1, 1, 1},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, res: false},
		{name: "more dublicate", board: Board{
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{2, 1, 3, 4, 5, 3, 7, 8, 9},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, res: false},
		{name: "dublicate in y", board: Board{
			Row{0, 0, 0, 0, 1, 0, 0, 0, 0},
			Row{2, 1, 3, 4, 5, 6, 7, 8, 9},
			Row{0, 0, 0, 0, 2, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 3, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 4, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 5, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 6, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 7, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 8, 0, 0, 0, 0},
		}, res: false},
		{name: "more dublicate in y", board: Board{
			Row{0, 6, 0, 0, 1, 0, 0, 0, 0},
			Row{2, 1, 3, 4, 5, 6, 7, 8, 9},
			Row{0, 7, 0, 0, 2, 0, 0, 0, 0},
			Row{0, 7, 0, 0, 3, 0, 0, 0, 0},
			Row{0, 2, 0, 0, 4, 0, 0, 0, 0},
			Row{0, 3, 0, 0, 5, 0, 0, 0, 0},
			Row{0, 4, 0, 0, 6, 0, 0, 0, 0},
			Row{0, 5, 0, 0, 7, 0, 0, 0, 0},
			Row{0, 6, 0, 0, 8, 0, 0, 0, 0},
		}, res: false},
		{name: "dublicate in square", board: Board{
			Row{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{9, 8, 2, 7, 6, 5, 6, 3, 1},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, res: false},
		{name: "dublicate in other square", board: Board{
			Row{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{9, 6, 4, 7, 6, 5, 8, 3, 1},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, res: false},
		{name: "no dublicate square", board: Board{
			Row{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{9, 6, 4, 7, 2, 8, 5, 3, 1},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, res: true},
		{name: "good part board", board: Board{
			Row{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Row{4, 5, 6, 7, 8, 9, 1, 2, 3},
			Row{9, 8, 7, 1, 2, 3, 4, 5, 6},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
			Row{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}, res: true},
		{name: "good board", board: Board{
			Row{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Row{4, 5, 6, 7, 8, 9, 1, 2, 3},
			Row{9, 8, 7, 1, 0, 3, 4, 5, 6},
			Row{2, 3, 4, 5, 6, 7, 8, 9, 1},
			Row{5, 6, 8, 9, 1, 2, 3, 4, 7},
			Row{7, 9, 1, 0, 4, 8, 0, 6, 5},
			Row{0, 0, 5, 6, 7, 4, 9, 1, 8},
			Row{8, 0, 0, 0, 3, 5, 6, 7, 4},
			Row{6, 0, 0, 8, 9, 1, 5, 3, 2},
		}, res: true},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.res != CheckBoard(&tc.board) {
				t.Errorf("Got %t, should be %t", CheckBoard(&tc.board), tc.res)
			}
		})
	}
}

func boardToString(board *Board) string {
	s := ""
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s += fmt.Sprint(board[i][j])
		}
	}
	return s
}

func TestExamples(t *testing.T) {
	var examples []Sudoku
	data, err := ioutil.ReadFile("examples.json")
	if err != nil {
		t.Fatal("Cannot read example.json file")
	}
	json.Unmarshal(data, &examples)
	for _, sudoku := range examples {
		t.Run(fmt.Sprint(sudoku.Id), func(t *testing.T) {
			solution := Run(stringToBoard(sudoku.Mission))
			if boardToString(solution) != sudoku.Solution {
				t.Errorf("Wrong Solution. Got %s, should be %s", boardToString(solution), sudoku.Solution)
			}
			solution = nil
		})
	}
}
