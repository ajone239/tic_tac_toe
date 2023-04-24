package game

import (
	"fmt"
	"math/rand"
)

type Player interface {
	GetMove(board *Board) (int, int)
}

type HumanPlayer struct{}

var _ Player = (*HumanPlayer)(nil)

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

func (p *RandomPlayer) GetMove(board *Board) (int, int) {
	moves := board.listMoves()

	move := moves[rand.Intn(len(moves))]
	return move[0], move[1]
}
