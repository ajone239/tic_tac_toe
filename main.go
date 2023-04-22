package main

import (
  "github.com/ajone239/tic_tac_toe/game"
	"flag"
)

func main() {
	// Get player types from command line arguments
	player1Type := flag.String("p1", "random", "Player 1 type")
	player2Type := flag.String("p2", "random", "Player 2 type")

	flag.Parse()

	player1 := getPlayerType(*player1Type, 1)
	player2 := getPlayerType(*player2Type, 2)

	// Create new game
	game := game.NewGame(player1, player2)

	// Run the game
	game.Loop()
}

// Get player type from string argument
func getPlayerType(playerType string, playerNumber int) game.Player {
	switch playerType {
	case "random", "r":
		return game.RandomPlayer{}
	case "human", "h":
		return game.HumanPlayer{}
	case "bot", "b":
		bot := game.NewBotPlayer(playerNumber)
		return bot
	default:
		return game.RandomPlayer{}
	}
}
