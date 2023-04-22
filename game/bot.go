package game

import (
	"fmt"
	"math/rand"
)

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

// Print the game tree
func (p BotPlayer) PrintTree() {
  p.game_tree.PrintTree()
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

func (p BotPlayer) GetMove(board *Board) (int, int) {
	// For now, just return a random move
	moves := board.listMoves()

	// Get random move
	random_move := moves[rand.Intn(len(moves))]
	i, j := random_move[0], random_move[1]
	return i, j
}

/*
 * Game Tree
 */

type gameTree struct {
	// The tree for calculating the best move
	root *treeNode
	// A map of boards to nodes for quick lookup
	nodeMap map[*Board]*treeNode
}

// BFS print the tree
func (g *gameTree) PrintTree() {
	queue := []*treeNode{g.root}
	for len(queue) > 0 {
		node := queue[0]

		// Print the node
    fmt.Printf("Board:" +
                node.board.String() +
                " Eval:" +
                fmt.Sprintln(node.eval) +
                " Move:" +
                fmt.Sprint(node.i) +
                "," +
                fmt.Sprintln(node.j))

		queue = queue[1:]
		for _, child := range node.children {
			queue = append(queue, child)
		}
	}
}

func newGameTree(board *Board) *gameTree {
	root := &treeNode{board: board, eval: 0, i: -1, j: -1}
	nodeMap := make(map[*Board]*treeNode)
	nodeMap[board] = root

	return &gameTree{root: root, nodeMap: nodeMap}
}

/*
 * Tree Node
 */

type treeNode struct {
	board    *Board
	i, j     int
	eval     int
	children []*treeNode
}

func (g *gameTree) expandNode(node *treeNode, square_to_play square) {
	moves := node.board.listMoves()

	if len(moves) == 0 {
		return
	}

	for _, move := range moves {
		i, j := move[0], move[1]
		new_board := node.board
		new_board.MakeMove(i, j, square_to_play)
		new_node := &treeNode{board: new_board, eval: 0, i: i, j: j}
		node.children = append(node.children, new_node)
		g.nodeMap[new_board] = new_node
	}
}
