package game

import "fmt"

type GameResult int

const (
	Cross  GameResult = iota
	Draw   GameResult = iota
	Nought GameResult = iota
)

type Game struct {
	crossPlayer  *Player
	noughtPlayer *Player
	board        *Board
	whosTurn     square
}

func NewGame(crossPlayer, noughtPlayer *Player) *Game {
	return &Game{
		crossPlayer:  crossPlayer,
		noughtPlayer: noughtPlayer,
		board:        NewBoard(),
		whosTurn:     cross,
	}
}

func (g *Game) Loop() GameResult {
	var i, j int
	for {
		// Get input
		var player *Player
		if g.isPlayer1() {
			player = g.crossPlayer
		} else {
			player = g.noughtPlayer
		}

    if (*player).IsHuman() {
      // Display the board to the user
      fmt.Println(g.board)
      fmt.Println(g.whosTurnStr(), "'s turn to go (i j):")
    }


		// TODO(austin): Does this have performance implications?
		i, j = (*player).GetMove(g.board)

		// set the move
		g.board.MakeMove(i, j, g.whosTurn)

    // Print a non-human's move
    if !(*player).IsHuman() {
      fmt.Println(g.whosTurnStr(), "moved to", i, j)
    }

		// check for winning
		winner := g.board.CheckForWin()

		if !winner.IsBlank() {
			fmt.Println(g.whosTurnStr(), "has won!")
			fmt.Println(g.board)
			return winner.ToGameResult()
		} else if g.board.IsFull() {
			fmt.Println("Draw!")
			fmt.Println(g.board)
			return Draw
		}
		// switch
		g.SwitchPlayer()
	}
}

// Is it player 1's turn?
func (g *Game) isPlayer1() bool {
	return g.whosTurn == cross
}

// Switch the player
func (g *Game) SwitchPlayer() {
	g.whosTurn = g.whosTurn.Switch()
}

// Print whose turn it is long
func (g *Game) whosTurnStr() string {
	switch g.whosTurn {
	case cross:
		return "Cross"
	case nought:
		return "Nought"
	default:
		panic("Invalid player")
	}
}
