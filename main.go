package main

import (
	"example/game"
	"fmt"
)

func main() {
	board := game.NewBoard()

	var i, j int
	for {
		// Display the board to the user
		fmt.Println(board)
		fmt.Println(board.WhosMove(), "'s turn to go (i j):")

		// Get input
		for {
			fmt.Scan(&i, &j)
			if !board.CheckGoodMove(i, j) {
        fmt.Println("Bad move -- Try again")
				continue
			}
			break
		}

		// set the move
		board.MakeMove(i,j)

		// check for winning
		winner := board.CheckForWin()

		if !winner.IsBlank() {
			fmt.Println(board.WhosMove(), "has won!")
			fmt.Println(board)
			return
		}
		// switch
    board.SwitchPlayer()
	}

}
