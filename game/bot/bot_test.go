package bot

import (
  "testing"
  "fmt"

  "github.com/ajone239/tic_tac_toe/game"
)

func TestCheckforWinOrDraw(t *testing.T) {
  node := &treeNode{
    board:             game.NewBoard(),
    move_square:       game.Cross,
    move_children_map: make(map[playerMove]*treeNode),
    eval:              0,
    expanded:          false,
  }
  node.board.MakeMove(0, 0, game.Cross)
  node.board.MakeMove(0, 1, game.Cross)
  node.board.MakeMove(0, 2, game.Cross)
  if eval, is_leaf := checkForWinOrDraw(node.board); eval != 1 || !is_leaf {
    t.Log("Eval", eval,"IsLeaf", is_leaf)
    t.Error("Check for win failed 1")
  }

  // Clear the board
  node.board = game.NewBoard()
  node.board.MakeMove(0, 0, game.Cross)
  node.board.MakeMove(0, 1, game.Cross)
  node.board.MakeMove(0, 2, game.Nought)
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
  board *game.Board
  player game.Square
  expected_move playerMove
}

var MiniMaxTests = []MiniMaxTest {
  {
    game.NewBoardFromArray(
      [3][3]game.Square{
        {game.Cross, game.Nought, game.Cross},
        {game.Blank, game.Nought, game.Blank},
        {game.Nought, game.Cross, game.Cross},
      },
    ),
    game.Nought, playerMove{1, 0},
  },
  {
    game.NewBoardFromArray(
      [3][3]game.Square{
        {game.Cross, game.Blank, game.Nought},
        {game.Nought, game.Nought, game.Cross},
        {game.Cross, game.Blank, game.Cross},
      },
    ),
    game.Nought, playerMove{2, 1},
  },
  {
    game.NewBoardFromArray(
      [3][3]game.Square{
        {game.Cross, game.Blank, game.Cross},
        {game.Nought, game.Nought, game.Cross},
        {game.Cross, game.Blank, game.Nought},
      },
    ),
    game.Nought, playerMove{0, 1},
  },
  {
    game.NewBoardFromArray(
      [3][3]game.Square{
        {game.Cross, game.Blank, game.Blank},
        {game.Blank, game.Nought, game.Blank},
        {game.Cross, game.Nought, game.Cross},
      },
    ),
    game.Nought, playerMove{0, 1},
  },
}

func TestMiniMax(t *testing.T) {
  fmt.Println("********************************************")
  // Create a game tree
  g := newGameTree(game.NewBoard(), false)

  // Run through the tests
  for _, test := range MiniMaxTests {
    node := g.nodeMap[boardString(test.board.String())]
    if !node.board.CheckGoodMove(test.expected_move.i, test.expected_move.j) {
      t.Error("Bad test")
    }
    if move, _ := node.getMinimaxMove(test.player == game.Cross, false); move != test.expected_move {
      fmt.Println("********************************************")
      g.PrintTree(node)
      fmt.Println("********************************************")
      t.Error("MiniMax failed, got", move, "expected", test.expected_move)
    }
  }

}
