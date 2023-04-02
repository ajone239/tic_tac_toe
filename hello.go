package main

import "fmt"

func main() {
	board := newBoard()
	var whos_move square = cross

	var i, j int
	for {
		// Display the board to the user
		fmt.Println(board)
		fmt.Println(squareToWord(whos_move), "'s turn to go (i j):")

		// Get input
		for {
			fmt.Scan(&i, &j)
			if !board.checkGoodMove(i, j) {
        fmt.Println("Bad move -- Try again")
				continue
			}
			break
		}

		// set the move
		board.board[i][j] = whos_move

		// switch
		switch whos_move {
		case nought:
			whos_move = cross
		case cross:
			whos_move = nought
		default:
			fmt.Println("unreachable")
			return
		}

		// check for winning
		winner := board.checkForWin()

		if winner != blank {
			fmt.Println(squareToWord(whos_move), "has won!")
			fmt.Println(board)
			return
		}
	}

}
