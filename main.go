/*
 * @author Austin Jones
 * @version 1.0
 */
package main

import (
	"flag"
  "fmt"
  "sync"
	tic_tac_toe "github.com/ajone239/tic_tac_toe/game"
)

func main() {
	// Get player types from command line arguments
	player1Type := flag.String("p1", "human", "Player 1 type")
	player2Type := flag.String("p2", "bot", "Player 2 type")
  gameType := flag.String("g", "tic_tac_toe", "Game type")
  botVsBot := flag.Bool("bvb", false, "Run bot vs bot games")
  botVsBotGames := flag.Int("bvb-games", 100, "Number of bot vs bot games to run")
  debug := flag.Bool("debug", false, "Debug mode")

	flag.Parse()


  if *botVsBot {
    botVsBotTest(*botVsBotGames)
    return
  }

  // Get player getters
  getPlayerType := getPlayerGetter(*gameType)

	player1 := getPlayerType(*player1Type, 1, *debug)
	player2 := getPlayerType(*player2Type, 2, *debug)

	// Create new game
	game := tic_tac_toe.NewGame(&player1, &player2)

	// Run the game
	game.Loop()
}

// Run bot vs bot games in separate goroutines and make sure they draw
func botVsBotTest(games int) {
  gameResults := make(chan tic_tac_toe.GameResult)

  wg := new(sync.WaitGroup)

  for i := 0; i < games; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      playerType := "bot"
      getPlayerType := getPlayerGetter("tic_tac_toe")
      player1 := getPlayerType(playerType, 1, false)
      player2 := getPlayerType(playerType, 2, false)
      game := tic_tac_toe.NewGame(&player1, &player2)
      gameResults <- game.Loop()
    }()
  }

  go func() {
    wg.Wait()
    close(gameResults)
  }()

  for result := range gameResults {
    if result != tic_tac_toe.Draw {
      panic("Bot vs Bot game did not draw")
    }
  }
  fmt.Println("All bot vs bot games drew")
}

// Get player getter from game type
func getPlayerGetter(gameType string) func(string, int, bool) tic_tac_toe.Player {
  switch gameType {
  case "tic_tac_toe":
    return getPlayerTypeTicTacToe
  default:
    return getPlayerTypeTicTacToe
  }
}

// Get player type from string argument
func getPlayerTypeTicTacToe(playerType string, playerNumber int, debug bool) tic_tac_toe.Player {
	switch playerType {
	case "random", "r":
		return new(tic_tac_toe.RandomPlayer)
	case "human", "h":
		return new(tic_tac_toe.HumanPlayer)
	case "bot", "b":
		bot := tic_tac_toe.NewBotPlayer(playerNumber, debug)
		return bot
	default:
		return new(tic_tac_toe.RandomPlayer)
	}
}
