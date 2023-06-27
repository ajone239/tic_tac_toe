package bot

import (
	"fmt"

  "github.com/ajone239/tic_tac_toe/game"
)

type boardString string

// Max val constants
const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

/*
 * BotPlayer
 *  - Implements the Player interface
 *  - Uses a game tree to determine the best move
 *  - Uses a minimax algorithm to determine the best move
 */

var _ game.Player = (*BotPlayer)(nil)

type BotPlayer struct {
	game_tree *gameTree
	noc       game.Square
  debug     bool
}

func (p *BotPlayer) IsHuman() bool {
  return false
}

func NewBotPlayer(player int, debug bool) *BotPlayer {
	var noc game.Square
	if player == 1 {
		noc = game.Cross
	} else {
		noc = game.Nought
	}
	board := game.NewBoard()
	tree := newGameTree(board, debug)

	return &BotPlayer{game_tree: tree, noc: noc, debug: debug}
}

/*
 * GetMove
 *  - Implements the Player interface
 *  - Uses a game tree to determine the best move
 *  - Uses a minimax algorithm to determine the best move
 */
func (p *BotPlayer) GetMove(board *game.Board) (int, int) {
  // Check if the board is in the tree and get the node
  node, ok := p.game_tree.nodeMap[boardString(board.String())]
  if !ok {
    panic("Board not in tree")
  }

  if p.debug {
    fmt.Println()
    fmt.Println("**********")
    fmt.Println()
  }

  // Maximize if cross, minimize if nought
	best_move, best_eval := node.getMinimaxMove(p.noc == game.Cross, p.debug)

  if p.debug {
    fmt.Println("Best move:", best_move, "Eval:", best_eval)
  }

	if !board.CheckGoodMove(best_move.i, best_move.j) {
		panic("Bad move")
	}

  if p.debug {
    fmt.Println()
    fmt.Println("**********")
    fmt.Println()
  }

	return best_move.i, best_move.j
}

