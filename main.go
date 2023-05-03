package main

import (
	"flag"
	tic_tac_toe "github.com/ajone239/tic_tac_toe/game"
)

func main() {
	// Get player types from command line arguments
	player1Type := flag.String("p1", "human", "Player 1 type")
	player2Type := flag.String("p2", "bot", "Player 2 type")
  gameType := flag.String("g", "tic_tac_toe", "Game type")

	flag.Parse()

  // Get player getters
  getPlayerType := getPlayerGetter(*gameType)

	player1 := getPlayerType(*player1Type, 1)
	player2 := getPlayerType(*player2Type, 2)

	// Create new game
	game := tic_tac_toe.NewGame(&player1, &player2)

	// Run the game
	game.Loop()
}

// Get player getter from game type
func getPlayerGetter(gameType string) func(string, int) tic_tac_toe.Player {
  switch gameType {
  case "tic_tac_toe":
    return getPlayerTypeTicTacToe
  default:
    return getPlayerTypeTicTacToe
  }
}

// Get player type from string argument
func getPlayerTypeTicTacToe(playerType string, playerNumber int) tic_tac_toe.Player {
	switch playerType {
	case "random", "r":
		return new(tic_tac_toe.RandomPlayer)
	case "human", "h":
		return new(tic_tac_toe.HumanPlayer)
	case "bot", "b":
		bot := tic_tac_toe.NewBotPlayer(playerNumber)
		return bot
	default:
		return new(tic_tac_toe.RandomPlayer)
	}
}
