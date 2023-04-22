package game

import "math/rand"

var _ Player = (*BotPlayer)(nil)

type BotPlayer struct {
  game_tree *gameTree
  noc square
}

func (p BotPlayer)GetMove(board *Board) (int, int) {
  moves := board.listMoves()

  // Get random move
  random_move := moves[rand.Intn(len(moves))]
  i, j := random_move[0], random_move[1]
  return i, j
}


func newBotPlayer(noc square) *BotPlayer {
  board := NewBoard()
  game_tree := newGameTree(board)
  return &BotPlayer{game_tree: game_tree, noc: noc}
}

type gameTree struct {
  root *treeNode
  nodeMap map[*Board]*treeNode
}

func newGameTree(board *Board) *gameTree {
  root := &treeNode{board: board, eval: 0, i: -1, j: -1}
  nodeMap := make(map[*Board]*treeNode)
  nodeMap[board] = root

  return &gameTree{root: root, nodeMap: nodeMap}
}

type treeNode struct {
  board *Board
  i, j int
  eval int
  children []*treeNode
}

func (g *gameTree)expandNode(node *treeNode) {
  moves := node.board.listMoves()

  if len(moves) == 0 {
    return
  }

  for _, move := range moves {
    i, j := move[0], move[1]
    new_board := node.board
    // fixme
    // new_board.MakeMove(i, j,)
    new_node := &treeNode{board: new_board, eval: 0, i: i, j: j}
    node.children = append(node.children, new_node)
    g.nodeMap[new_board] = new_node
  }
}


