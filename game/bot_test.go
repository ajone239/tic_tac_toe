package game

import (
  "testing"
  "fmt"
)

func TestCheckforWinOrDraw(t *testing.T) {
  node := &treeNode{
    board:             NewBoard(),
    move_square:       cross,
    move_children_map: make(map[playerMove]*treeNode),
    eval:              0,
    expanded:          false,
  }
  node.board.MakeMove(0, 0, cross)
  node.board.MakeMove(0, 1, cross)
  node.board.MakeMove(0, 2, cross)
  if eval, is_leaf := checkForWinOrDraw(node.board); eval != 1 || !is_leaf {
    t.Log("Eval", eval,"IsLeaf", is_leaf)
    t.Error("Check for win failed 1")
  }

  // Clear the board
  node.board = NewBoard()
  node.board.MakeMove(0, 0, cross)
  node.board.MakeMove(0, 1, cross)
  node.board.MakeMove(0, 2, nought)
  if eval, is_leaf := checkForWinOrDraw(node.board); eval != 0 || is_leaf {

    t.Error("Check for draw failed 2")
  }
}

func TestMom(t *testing.T) {
  a := []int{1, 2, 3, 4, 5}
  max := mom(a, true)
  min := mom(a, false)
  if max != 5 || min != 1 {
    t.Error("Mom failed")
  }
}

type MiniMaxTest struct {
  board *Board
  player square
  expected_move playerMove
}

var MiniMaxTests = []MiniMaxTest {
  {
    &Board{
      [3][3]square{
        {cross, nought, cross},
        {blank, nought, blank},
        {nought, cross, cross},
      },
    },
    nought, playerMove{1, 2},
  },
  {
    &Board{
      [3][3]square{
        {cross, nought, cross},
        {blank, nought, blank},
        {cross, cross, nought},
      },
    },
    nought, playerMove{1, 0},
  },
  {
    &Board{
      [3][3]square{
        {cross, blank, nought},
        {nought, nought, cross},
        {cross, blank, cross},
      },
    },
    nought, playerMove{2, 1},
  },
  {
    &Board{
      [3][3]square{
        {cross, blank, cross},
        {nought, nought, cross},
        {cross, blank, nought},
      },
    },
    nought, playerMove{0, 1},
  },
  {
    &Board{
      [3][3]square{
        {cross, blank, blank},
        {blank, nought, blank},
        {cross, nought, cross},
      },
    },
    nought, playerMove{0, 1},
  },
}

func TestMiniMax(t *testing.T) {
  fmt.Println("********************************************")
  // Create a game tree
  g := newGameTree(NewBoard())

  // Run through the tests
  for _, test := range MiniMaxTests {
    node := g.nodeMap[boardString(test.board.String())]
    if !node.board.CheckGoodMove(test.expected_move.i, test.expected_move.j) {
      t.Error("Bad test")
    }
    if move, _ := node.getMinimaxMove(test.player == cross); move != test.expected_move {
      fmt.Println("********************************************")
      g.PrintTree(node)
      fmt.Println("********************************************")
      t.Error("MiniMax failed, got", move, "expected", test.expected_move)
    }
  }

}
