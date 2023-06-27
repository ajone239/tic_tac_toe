package game

import (
	"strconv"
)

const BOARD_SIDE_LENGTH = 3

/*
 * square
 */

type Square int

// Square values
const (
	Blank  Square = iota
	Nought Square = iota
	Cross  Square = iota
)

func (s Square) ToGameResult() GameResult {
	switch s {
	case Nought:
		return NoughtWin
	case Cross:
		return CrossWin
	default:
		return Draw
	}
}

// square to string
func (s Square) String() string {
	switch s {
	case Blank:
		return "_"
	case Nought:
		return "O"
	case Cross:
		return "X"
	default:
		panic("unreachable")
	}
}

func squareToWord(s Square) string {
	switch s {
	case Nought:
		return "Nought"
	case Cross:
		return "Cross"
	default:
		return "Blank"
	}
}

func (s Square) IsBlank() bool {
	return s == Blank
}

func (s Square) Switch() Square {
	switch s {
	case Nought:
		return Cross
	case Cross:
		return Nought
	default:
		return Blank
	}
}

/*
 * Board
 */

type Board struct {
	board [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]Square
}

func NewBoard() *Board {
	board := Board{
		board: [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]Square{
			{Blank, Blank, Blank},
			{Blank, Blank, Blank},
			{Blank, Blank, Blank},
		},
	}
	return &board
}

func NewBoardFromArray(a [BOARD_SIDE_LENGTH][BOARD_SIDE_LENGTH]Square) *Board {
  board := Board{
    board: a,
  }
  return &board
}

// Copy a board
func (board *Board) Copy() *Board {
	new_board := NewBoard()
	for i := 0; i < BOARD_SIDE_LENGTH; i++ {
		for j := 0; j < BOARD_SIDE_LENGTH; j++ {
			new_board.board[i][j] = board.board[i][j]
		}
	}
	return new_board
}

func (board *Board) IsFull() bool {
	for i := 0; i < BOARD_SIDE_LENGTH; i++ {
		for j := 0; j < BOARD_SIDE_LENGTH; j++ {
			if board.board[i][j] == Blank {
				return false
			}
		}
	}
	return true
}

func (board *Board) CheckGoodMove(i, j int) bool {
	if i >= BOARD_SIDE_LENGTH || i < 0 ||
		j >= BOARD_SIDE_LENGTH || j < 0 {
		return false
	}

	return board.board[i][j] == Blank
}

func (board *Board) MakeMove(i, j int, s Square) {
	if s.IsBlank() {
		panic("Cannot make a Blank move")
	}
	board.board[i][j] = s
}

// Evaluate a board by checking for a win and mapping the value to a score
func (board *Board) Evaluate() int {
	winner := board.CheckForWin()
	switch winner {
	case Cross:
		return 1
	case Nought:
		return -1
	default:
		return 0
	}
}

func (board *Board) CheckForWin() Square {
	row_winner := board.checkRows()
	col_winner := board.checkColumns()
	diag_winner := board.checkDiagonols()

	winners := []Square{row_winner, col_winner, diag_winner}

	for _, w := range winners {
		if w != Blank {
			return w
		}
	}

	return Blank
}

func (board *Board) ListMoves() [][2]int {
	rv := make([][2]int, 0)

	for i := 0; i < BOARD_SIDE_LENGTH; i++ {
		for j := 0; j < BOARD_SIDE_LENGTH; j++ {
			if board.board[i][j] == Blank {
				rv = append(rv, [2]int{i, j})
			}
		}
	}
	return rv
}

func (board *Board) checkRows() Square {
	for _, row := range board.board {
		winner := checkSquaresForWin(row)
		if winner != Blank {
			return winner
		}
	}
	return Blank
}

func (board *Board) checkColumns() Square {
	for i := 0; i < BOARD_SIDE_LENGTH; i++ {
		column := [BOARD_SIDE_LENGTH]Square{
			board.board[0][i],
			board.board[1][i],
			board.board[2][i],
		}
		winner := checkSquaresForWin(column)
		if winner != Blank {
			return winner
		}
	}
	return Blank
}

func (board *Board) checkDiagonols() Square {
	diag := [BOARD_SIDE_LENGTH]Square{
		board.board[0][0],
		board.board[1][1],
		board.board[2][2],
	}
	winner := checkSquaresForWin(diag)
	if winner != Blank {
		return winner
	}

	diag = [BOARD_SIDE_LENGTH]Square{
		board.board[2][0],
		board.board[1][1],
		board.board[0][2],
	}

	winner = checkSquaresForWin(diag)
	if winner != Blank {
		return winner
	}

	return Blank
}

func checkSquaresForWin(s [BOARD_SIDE_LENGTH]Square) Square {
	if s[0] == Blank {
		return Blank
	}
	equal := true
	for _, v := range s[1:] {
		equal = equal && (v == s[0])
	}

	if equal {
		return s[0]
	} else {
		return Blank
	}
}

func (b Board) String() string {
	ret_string := ""
	ret_string += " 012 - j\n"
	for i := 0; i < BOARD_SIDE_LENGTH; i++ {
		ret_string += strconv.Itoa(i)
		for j := 0; j < BOARD_SIDE_LENGTH; j++ {
			ret_string += b.board[i][j].String()
		}
		ret_string += "\n"
	}
	ret_string += "i\n"

	return ret_string
}
