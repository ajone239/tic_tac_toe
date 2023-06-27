package bot

import (
  "fmt"
  "github.com/ajone239/tic_tac_toe/game"
)

/*
 * Tree Node
 */

type treeNode struct {
	board             *game.Board
	move_square       game.Square
	eval              int
	move_children_map map[playerMove]*treeNode
	expanded          bool
  is_leaf           bool
}

type playerMove struct {
	i, j int
}

func nullMove() playerMove {
  return playerMove{-1, -1}
}

// treeNode to string
func (n *treeNode) String() string {
	ret_string := ""
	ret_string += "Board:\n"
	ret_string += n.board.String()
	ret_string += "Eval:"
	ret_string += fmt.Sprintln(n.eval)
	ret_string += "Move:"

	return ret_string
}

func (g *gameTree) expandNode(node *treeNode, square_to_play game.Square) {

  // Get all the moves
	moves := node.board.ListMoves()

	for _, move := range moves {
		i, j := move[0], move[1]
		// Copy the board and make the move
		new_board := node.board.Copy()
		new_board.MakeMove(i, j, square_to_play)

    // Check if the board has already been added to the tree
		if n, ok := g.nodeMap[boardString(new_board.String())]; ok {
			node.move_children_map[playerMove{i, j}] = n
			continue
		}

		eval, is_leaf := checkForWinOrDraw(new_board)

		// Build the new node and add it to the tree
		new_node := &treeNode{
			board:             new_board,
			move_square:       square_to_play,
			move_children_map: make(map[playerMove]*treeNode),
			eval:              eval,
			expanded:          is_leaf,
      is_leaf:           is_leaf,
		}
		node.move_children_map[playerMove{i, j}] = new_node
		// Add the node to the map
		g.nodeMap[boardString(new_board.String())] = new_node
	}
}

// Check node for win or Draw
func checkForWinOrDraw(board *game.Board) (int, bool) {
  // Check for win
  eval := board.Evaluate()
  // Is a leaf if there is a win or the board is full
  is_leaf := eval != 0 || board.IsFull()

  return eval, is_leaf
}

// Minimax for node
func (n *treeNode) getMinimaxMove(max_or_min bool, debug bool) (playerMove, int) {
  // Sentinel null move
	var best_move playerMove = nullMove()
  // Sentinel null move allows for us to not init best_eval
  var best_eval int = 0

	for move, child := range n.move_children_map {
		// MiniMax the tree
		eval:= child.minimax(!max_or_min)

		// Update accordingly
		if (best_move == nullMove()) ||
			(max_or_min && eval > best_eval) ||
			(!max_or_min && eval < best_eval) {
			best_eval = eval
			best_move = move
    }
    // Print the move and eval
    if debug {
      fmt.Println(">>>  Move:", move, "Eval:", eval)
    }
	}

  if debug {
    fmt.Println()
    fmt.Println("Best Move:", best_move, "Best Eval:", best_eval)
  }

	return best_move, best_eval
}

func (n *treeNode) minimax(max_or_min bool) int {
	if n.is_leaf {
		return n.eval
	}

	child_evals := make([]int, 0)
	for _, child := range n.move_children_map {
		eval := child.minimax(!max_or_min)
		child_evals = append(child_evals, eval)
	}
	best_eval := mom(child_evals, max_or_min)

	return best_eval
}

func mom(a []int, max_or_min bool) int {
	if len(a) == 0 {
		return 0
	}
	best := a[0]
	for _, v := range a {
		if (max_or_min && v > best) || (!max_or_min && v < best) {
			best = v
		}
	}
	return best
}
