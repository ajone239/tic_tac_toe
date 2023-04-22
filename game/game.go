package game

import "fmt"

type Game struct {
	crossPlayer  Player
	nougthPlayer Player
	board        *Board
	whosTurn     square
}

func NewGame(crossPlayer, nougthPlayer Player) *Game {
	return &Game{
		crossPlayer:  crossPlayer,
		nougthPlayer: nougthPlayer,
		board:        NewBoard(),
		whosTurn:     cross,
	}
}

func (g *Game) Loop() {
	var i, j int
	for {
		// Display the board to the user
		fmt.Println(g.board)
		fmt.Println(g.whosTurnStr(), "'s turn to go (i j):")

		// Get input
		var player Player
		if g.isPlayer1() {
			player = g.crossPlayer
		} else {
			player = g.nougthPlayer
		}

		i, j = player.GetMove(g.board)

		// set the move
		g.board.MakeMove(i, j, g.whosTurn)

		// check for winning
		winner := g.board.CheckForWin()

		if !winner.IsBlank() {
			fmt.Println(g.whosTurnStr(), "has won!")
			fmt.Println(g.board)
			return
		} else if g.board.IsFull() {
			fmt.Println("Draw!")
			fmt.Println(g.board)
			return
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
	switch g.whosTurn {
	case cross:
		g.whosTurn = nought
	case nought:
		g.whosTurn = cross
	default:
		panic("Invalid player")
	}
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
