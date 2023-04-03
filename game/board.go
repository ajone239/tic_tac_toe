package game

import (
	"fmt"
	"strconv"
)

const BOARD_SIZE = 3

type square int

const (
	blank  = iota
	nought = iota
	cross  = iota
)

type Board struct {
	board [BOARD_SIZE][BOARD_SIZE]square
  whos_move square
}

func squareToWord(s square) string {
	switch s {
	case nought:
		return "nought"
	case cross:
		return "cross"
	default:
		return "blank"
	}
}

func (s square) IsBlank() bool {
  return s == blank
}

/*
 * Helpers
 */

func NewBoard() *Board {
	board := Board{
		board: [BOARD_SIZE][BOARD_SIZE]square{
			{blank, blank, blank},
			{blank, blank, blank},
			{blank, blank, blank},
		},
    whos_move: cross,
	}
	return &board
}

func (board *Board) SwitchPlayer() {
    switch board.whos_move {
		case nought:
			board.whos_move = cross
		case cross:
			board.whos_move = nought
		default:
      panic("unreachable")
		}
}

func (board *Board) WhosMove() string{
  return squareToWord(board.whos_move)
}

func (board *Board) CheckGoodMove(i, j int) bool {
	return !(i >= BOARD_SIZE || i < 0 ||
		j >= BOARD_SIZE || j < 0 ||
		board.board[i][j] != blank)
}

func (board *Board) MakeMove(i, j int) {
  		board.board[i][j] = board.whos_move
}

func (board *Board) CheckForWin() square {
	row_winner := board.checkRows()
	col_winner := board.checkColumns()
	diag_winner := board.checkDiagonols()

	winners := []square{row_winner, col_winner, diag_winner}

	for i, w := range winners {
		if w != blank {
      fmt.Println("this won: ", i)
			return w
		}
	}

	return blank
}

func (board *Board) listMoves() [][2]int {
	rv := make([][2]int, 0)

	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if board.board[i][j] == blank {
				rv = append(rv, [2]int{i, j})
			}
		}
	}
	return rv
}

func (board *Board) checkRows() square {
	for _, row := range board.board {
		winner := checkSquaresForWin(row)
		if winner != blank {
			return winner
		}
	}
	return blank
}

func (board *Board) checkColumns() square {
	for i := 0; i < BOARD_SIZE; i++ {
		row := [3]square{
			board.board[0][i],
			board.board[1][i],
			board.board[2][i],
		}
		winner := checkSquaresForWin(row)
    fmt.Println(row)
		if winner != blank {
			return winner
		}
	}
	return blank
}

func (board *Board) checkDiagonols() square {
	diag := [3]square{
		board.board[0][0],
		board.board[1][1],
		board.board[2][2],
	}
	winner := checkSquaresForWin(diag)
	if winner != blank {
		return winner
	}

	diag = [3]square{
		board.board[2][0],
		board.board[1][1],
		board.board[0][2],
	}

	winner = checkSquaresForWin(diag)
	if winner != blank {
		return winner
	}

	return blank
}

func checkSquaresForWin(s [3]square) square {
	if s[0] == blank {
		return blank
	}
	equal := true
	for _, v := range s {
		equal = equal && (v == s[0])
	}

	if equal {
		return s[0]
	} else {
		return blank
	}
}

func (b Board) String() string {
	ret_string := ""
	ret_string += " 012 - i\n"
	for i := 0; i < BOARD_SIZE; i++ {
		ret_string += strconv.Itoa(i)
		for j := 0; j < BOARD_SIZE; j++ {
			switch b.board[j][i] {
			case blank:
				ret_string += "_"
			case nought:
				ret_string += "O"
			case cross:
				ret_string += "X"
			}
		}
		ret_string += "\n"
	}
	ret_string += "j\n"

	return ret_string
}
