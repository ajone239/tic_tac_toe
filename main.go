package main

import (
	"flag"
	"github.com/ajone239/tic_tac_toe/game"
)

func main() {
	// Get player types from command line arguments
	player1Type := flag.String("p1", "human", "Player 1 type")
	player2Type := flag.String("p2", "bot", "Player 2 type")

	flag.Parse()

	player1 := getPlayerType(*player1Type, 1)
	player2 := getPlayerType(*player2Type, 2)

	// Create new game
	game := game.NewGame(&player1, &player2)

	// Run the game
	game.Loop()
}

// Get player type from string argument
func getPlayerType(playerType string, playerNumber int) game.Player {
	switch playerType {
	case "random", "r":
		return new(game.RandomPlayer)
	case "human", "h":
		return new(game.HumanPlayer)
	case "bot", "b":
		bot := game.NewBotPlayer(playerNumber)
		return bot
	default:
		return new(game.RandomPlayer)
	}
}
