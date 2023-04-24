package game

import (
	"fmt"
)

type boardString string

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

  // best_move, best_eval := node.getBestEvalMove(p.noc == cross)
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
    board: board,
    move_children_map: make(map[playerMove]*treeNode),
    eval: 0,
  }

	nodeMap := make(map[boardString]*treeNode)
	nodeMap[boardString(board.String())] = root

	g := gameTree{root: root, nodeMap: nodeMap}

	g.expandTree(root, cross)
  // g.PrintTree()

	return &g
}

// BFS print the tree
func (g *gameTree) PrintTree() {
	queue := []*treeNode{g.root}
	for len(queue) > 0 {
		node := queue[0]

    fmt.Println(node)

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
	board       *Board
	move_square square
	eval        int
  move_children_map map[playerMove]*treeNode
	expanded    bool
}

type playerMove struct {
  i, j int
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

		// Get the board if it is in the map
		if n, ok := g.nodeMap[boardString(new_board.String())]; ok {
			node.move_children_map[playerMove{i, j}] = n
      continue
    }
    eval, is_leaf := node.checkForWinOrDraw()

    // Build the new node and add it to the tree
    new_node := &treeNode{
      board: new_board,
      move_square: square_to_play,
      move_children_map: make(map[playerMove]*treeNode),
      eval: eval,
      expanded: is_leaf,
    }
    node.move_children_map[playerMove{i, j}] = new_node
    // Add the node to the map
    g.nodeMap[boardString(new_board.String())] = new_node
  }
}

// Check node for win or Draw
func (n *treeNode) checkForWinOrDraw() (int, bool) {
  // Check for win
  if winner := n.board.CheckForWin(); winner != blank {
    // Winning node
    switch winner {
    case cross:
      return 1, true
    case nought:
      return -1, true
    }
  } else if n.board.IsFull() {
    // Draw node
    return 0, true
  }
  return 0, false
}

// Get the best move for the node
func (n *treeNode) getBestEvalMove(max_or_min bool) (playerMove, int) {
  var best_eval int
  if max_or_min {
    best_eval = MinInt
  } else {
    best_eval = MaxInt
  }
  var best_move playerMove = playerMove{-1, -1}
  for move, child := range n.move_children_map {
    fmt.Println("Move:", move, "Eval:", child.eval)
    if (best_move == playerMove{-1, -1}) ||
      (max_or_min && child.eval > best_eval) ||
      (!max_or_min && child.eval < best_eval) {
        best_move = move
        best_eval = child.eval
    }
  }

  return best_move, best_eval
}

// Minimax for node
func (n *treeNode) getMinimaxMove(max_or_min bool) (playerMove, int) {
  var best_eval int
  if max_or_min {
    best_eval = MinInt
  } else {
    best_eval = MaxInt
  }
  var best_move playerMove = playerMove{-1, -1}
  for move, child := range n.move_children_map {
    // MiniMax the tree
    eval := child.minimax(!max_or_min)
    fmt.Println("Move:", move, "Eval:", eval, "BestMove:", best_move, "BestEval:", best_eval)

    // Update accordingly
    if (best_move == playerMove{-1, -1}) ||
       (max_or_min && eval > best_eval) ||
       (!max_or_min && eval < best_eval) {
      best_eval = eval
      best_move = move
     }
  }
  return best_move, best_eval
}

func (n *treeNode) minimax(max_or_min bool) int {
  if len(n.move_children_map) == 0 {
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
