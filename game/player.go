package game

import (
	"fmt"
	"math/rand"
)

type Player interface {
	GetMove(board *Board) (int, int)
  IsHuman() bool
}

type HumanPlayer struct{}

var _ Player = (*HumanPlayer)(nil)

func (p *HumanPlayer) IsHuman() bool {
  return true
}

func (p *HumanPlayer) GetMove(board *Board) (int, int) {
	// Get move from user
	var i, j int
	for {
		fmt.Scan(&i, &j)
		if !board.CheckGoodMove(i, j) {
			fmt.Println("Bad move -- Try again")
			continue
		}
		break
	}
	return i, j
}

type RandomPlayer struct{}

var _ Player = (*RandomPlayer)(nil)

func (p *RandomPlayer) IsHuman() bool {
  return false
}

func (p *RandomPlayer) GetMove(board *Board) (int, int) {
	moves := board.ListMoves()

	move := moves[rand.Intn(len(moves))]
	return move[0], move[1]
}
