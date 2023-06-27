package bot

import (
  "fmt"
  "github.com/ajone239/tic_tac_toe/game"
)

/*
 * Game Tree
 */

type gameTree struct {
	// The tree for calculating the best move
	root *treeNode
	// A map of boards to nodes for quick lookup
	nodeMap map[boardString]*treeNode
  debug   bool
}

func newGameTree(board *game.Board, debug bool) *gameTree {
	root := &treeNode{
		board:             board,
		move_children_map: make(map[playerMove]*treeNode),
		eval:              0,
	}

	nodeMap := make(map[boardString]*treeNode)
	nodeMap[boardString(board.String())] = root

	g := gameTree{root: root, nodeMap: nodeMap, debug: debug}

	g.expandTree(root, game.Cross)

  if debug {
     g.PrintTree(root)
  }

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
		if len(node.move_children_map) != 0 {
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

func (g *gameTree) expandTree(node *treeNode, square_to_play game.Square) {
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
}

