package game

import (
	"fmt"
	"math/rand"
)

type boardString string

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
	nodeMap map[boardString]*treeNode
	// Count the number of nodes with no moves
	no_moves_count int
}

// BFS print the tree
func (g *gameTree) PrintTree() {
	queue := []*treeNode{g.root}
	for len(queue) > 0 {
		node := queue[0]

		// Print the node
		fmt.Printf("Board:\n" +
			node.board.String() +
			"Eval:" +
			fmt.Sprintln(node.eval) +
			"Move:" +
			fmt.Sprint(node.i) +
			"," +
			fmt.Sprintln(node.j))
		fmt.Println()

		queue = queue[1:]
		for _, child := range node.children {
			queue = append(queue, child)
		}
	}
}

func newGameTree(board *Board) *gameTree {
	root := &treeNode{board: board, eval: 0, i: -1, j: -1}

	nodeMap := make(map[boardString]*treeNode)
	nodeMap[boardString(board.String())] = root

	g := gameTree{root: root, nodeMap: nodeMap, no_moves_count: 0}

	g.expandTree(root, cross)

	return &g
}

func (g *gameTree) expandTree(node *treeNode, square_to_play square) {

	// Expand the node
	g.expandNode(node, square_to_play)

	// Switch the square to play
	square_to_play = square_to_play.Switch()

	// Expand the children who have not been expanded
	for _, child := range node.children {
		if child.expanded {
			continue
		}
		g.expandTree(child, square_to_play)
		child.expanded = true
	}
}

func (g *gameTree) evalTree(node *treeNode, square_to_play square) {

}

/*
 * Tree Node
 */

type treeNode struct {
	board       *Board
	move_square square
	i, j        int
	eval        int
	children    []*treeNode
	expanded    bool
}

func contains(s [][2]int, e [2]int) bool {
	for _, a := range s {
		if a[0] == e[0] && a[1] == e[1] {
			return true
		}
	}
	return false
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
			node.children = append(node.children, n)
		} else {
			// Build the new node and add it to the tree
			new_node := &treeNode{board: new_board, move_square: square_to_play, eval: 0, i: i, j: j, expanded: false}
			node.children = append(node.children, new_node)
			// Add the node to the map
			g.nodeMap[boardString(new_board.String())] = new_node
		}
	}
}
