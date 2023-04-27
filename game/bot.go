package game

import (
	"fmt"
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

var _ Player = (*BotPlayer)(nil)

type BotPlayer struct {
	game_tree *gameTree
	noc       square
}

func NewBotPlayer(player int) *BotPlayer {
	var noc square
	if player == 1 {
		noc = cross
	} else {
		noc = nought
	}
	board := NewBoard()
	tree := newGameTree(board)

	return &BotPlayer{game_tree: tree, noc: noc}
}

func (p *BotPlayer) GetMove(board *Board) (int, int) {
	// Get the node for the board
	node := p.game_tree.nodeMap[boardString(board.String())]

	fmt.Println()
	fmt.Println("**********")
	fmt.Println()

  // Maximize if cross, minimize if nought
	best_move, best_eval := node.getMinimaxMove(p.noc == cross)

	fmt.Println("Best move:", best_move, "Eval:", best_eval)

	if !board.CheckGoodMove(best_move.i, best_move.j) {
		panic("Bad move")
	}

	fmt.Println()
	fmt.Println("**********")
	fmt.Println()

	return best_move.i, best_move.j
}

/*
 * Game Tree
 */

type gameTree struct {
	// The tree for calculating the best move
	root *treeNode
	// A map of boards to nodes for quick lookup
	nodeMap map[boardString]*treeNode
}

func newGameTree(board *Board) *gameTree {
	root := &treeNode{
		board:             board,
		move_children_map: make(map[playerMove]*treeNode),
		eval:              0,
	}

	nodeMap := make(map[boardString]*treeNode)
	nodeMap[boardString(board.String())] = root

	g := gameTree{root: root, nodeMap: nodeMap}

	g.expandTree(root, cross)
	// g.PrintTree()

	return &g
}

// BFS print the tree
func (g *gameTree) PrintTree(root *treeNode) {
	queue := []*treeNode{root}
	for len(queue) > 0 {
		node := queue[0]

		fmt.Println(node)
    fmt.Println()

		queue = queue[1:]
		for _, child := range node.move_children_map {
			queue = append(queue, child)
		}
	}

	// Count wins and draws and losses
	wins := 0
	draws := 0
	losses := 0

	for _, node := range g.nodeMap {
		if len(node.move_children_map) == 0 {
			continue
		}
		if node.eval == 1 {
			wins++
		} else if node.eval == 0 {
			draws++
		} else if node.eval == -1 {
			losses++
		}
	}

	fmt.Println("Tree size:", len(g.nodeMap))
	fmt.Println("CrossWins:", wins)
	fmt.Println("NoughtWins:", losses)
	fmt.Println("Draws:", draws)
}

func (g *gameTree) expandTree(node *treeNode, square_to_play square) {
	// Expand the node
	g.expandNode(node, square_to_play)

	if len(node.move_children_map) == 0 {
		return
	}

	// Switch the square to play
	square_for_children_to_play := square_to_play.Switch()

	// Expand the children who have not been expanded
	for _, child := range node.move_children_map {
		if child.expanded {
			continue
		}
		g.expandTree(child, square_for_children_to_play)
		child.expanded = true
	}

	// for _, child := range node.move_children_map {
	//   node.eval += child.eval
	// }
}

/*
 * Tree Node
 */

type treeNode struct {
	board             *Board
	move_square       square
	eval              int
	move_children_map map[playerMove]*treeNode
	expanded          bool
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

func (g *gameTree) expandNode(node *treeNode, square_to_play square) {
	moves := node.board.listMoves()

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

		eval, is_leaf := node.checkForWinOrDraw()

		// Build the new node and add it to the tree
		new_node := &treeNode{
			board:             new_board,
			move_square:       square_to_play,
			move_children_map: make(map[playerMove]*treeNode),
			eval:              eval,
			expanded:          is_leaf,
		}
		node.move_children_map[playerMove{i, j}] = new_node
		// Add the node to the map
		g.nodeMap[boardString(new_board.String())] = new_node

    bad_board_string := "Board:\n 012 - i\n0XO\n1OO\n2OX\nj\n"
    bad_board_string2 := "Board:\n 012 - i\n0XOX\n1OOX\n2OXX\nj\n"

    if new_board.String() == bad_board_string || new_board.String() == bad_board_string2 {
      // Print all information about the board
      fmt.Println("Bad board")
      fmt.Println("Move:", i, j)
      fmt.Println("Square:", square_to_play)
      fmt.Println("Eval:", eval)
      fmt.Println("Is leaf:", is_leaf)
      fmt.Println("Node:", new_node)
      fmt.Println("Parent:", node)
    }
	}
}

// Check node for win or Draw
func (n *treeNode) checkForWinOrDraw() (int, bool) {
  // Check for win
  eval := n.board.Evaluate()
  draw := true
  if eval == 0 {
    draw = n.board.IsFull()
  }
  return eval, draw
}

// Minimax for node
func (n *treeNode) getMinimaxMove(max_or_min bool) (playerMove, int) {
  // init best_eval
	var best_eval int
	if max_or_min {
		best_eval = MinInt
	} else {
		best_eval = MaxInt
	}

  // No draws
  best_draw_count := 0

  // Sentinel null move
	var best_move playerMove = nullMove()

	for move, child := range n.move_children_map {
		// MiniMax the tree
		eval, draw_count := child.minimax(!max_or_min)

		// Update accordingly
		if (best_move == nullMove()) ||
			(max_or_min && eval > best_eval) ||
			(!max_or_min && eval < best_eval)  ||
      (eval == best_eval && draw_count > best_draw_count) {
			best_eval = eval
			best_move = move
      best_draw_count = draw_count
      fmt.Println("Hit")
    }

    // Print best move and eval
    fmt.Println("Best Move:", best_move, "Eval:", best_eval)
    // Print the move and eval
    fmt.Println(">>>  Move:", move, "Eval:", eval, "Draws:", draw_count)
	}

  fmt.Println()
  fmt.Println("Best Move:", best_move, "Best Eval:", best_eval, "Draws:", best_draw_count)

	return best_move, best_eval
}

func (n *treeNode) minimax(max_or_min bool) (int, int) {
	if len(n.move_children_map) == 0 {
		return n.eval, 0
	}

  draw_count := 0
	child_evals := make([]int, 0)
	for _, child := range n.move_children_map {
		eval, _ := child.minimax(!max_or_min)
    if eval == 0 { draw_count++ }
		child_evals = append(child_evals, eval)
	}
	best_eval := mom(child_evals, max_or_min)

	return best_eval, draw_count
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
