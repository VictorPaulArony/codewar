package main

import (
	"fmt"
	"os"
	"strings"
)

// Define constants for player symbols
const (
	PlayerX = "X"
	PlayerO = "O"
	Empty   = " "
)

// Define the Tic-Tac-Toe board
var board [3][3]string

// func main() {
// 	// Initialize the board
// 	clearBoard()

// 	// Game loop
// 	for {
// 		printBoard()

// 		// Get player input
// 		var row, col int
// 		fmt.Printf("Player %s's turn (row, col): ", currentPlayer())
// 		fmt.Scanln(&row, &col)

// 		// Check if the move is valid
// 		if isValidMove(row, col) {
// 			board[row][col] = currentPlayer()

// 			// Check for win or draw
// 			if checkWin() {
// 				printBoard()
// 				fmt.Printf("Player %s wins!\n", currentPlayer())
// 				break
// 			} else if checkDraw() {
// 				printBoard()
// 				fmt.Println("It's a draw!")
// 				break
// 			}

// 			// Switch players
// 			switchPlayer()
// 		} else {
// 			fmt.Println("Invalid move. Please try again.")
// 		}
// 	}
// }

// clearBoard initializes the board with empty spaces
func clearBoard() {
	for i := range board {
		for j := range board[i] {
			board[i][j] = Empty
		}
	}
}

// printBoard prints the current state of the board with borders
func printBoard() {
	fmt.Println("  0 1 2")
	for i, row := range board {
		fmt.Printf("%d ", i)
		fmt.Println(strings.Join(row[:], "|"))
		if i < len(board)-1 {
			fmt.Println("  -----")
		}
	}
}

// currentPlayer returns the symbol of the current player
func currentPlayer() string {
	countX, countO := countSymbols()
	if countX <= countO {
		return PlayerX
	}
	return PlayerO
}

// isValidMove checks if the given move is valid
func isValidMove(row, col int) bool {
	return row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == Empty
}

// checkWin checks if the current player has won the game
func checkWin() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer() && board[i][1] == currentPlayer() && board[i][2] == currentPlayer() {
			return true // Horizontal win
		}
		if board[0][i] == currentPlayer() && board[1][i] == currentPlayer() && board[2][i] == currentPlayer() {
			return true // Vertical win
		}
	}
	if board[0][0] == currentPlayer() && board[1][1] == currentPlayer() && board[2][2] == currentPlayer() {
		return true // Diagonal win (top-left to bottom-right)
	}
	if board[0][2] == currentPlayer() && board[1][1] == currentPlayer() && board[2][0] == currentPlayer() {
		return true // Diagonal win (top-right to bottom-left)
	}
	return false
}

// checkDraw checks if the game has ended in a draw
func checkDraw() bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == Empty {
				return false
			}
		}
	}
	return true
}

// countSymbols returns the count of X's and O's on the board
func countSymbols() (int, int) {
	var countX, countO int
	for _, row := range board {
		for _, cell := range row {
			if cell == PlayerX {
				countX++
			} else if cell == PlayerO {
				countO++
			}
		}
	}
	return countX, countO
}

// switchPlayer switches the current player
// switchPlayer switches the current player
func switchPlayer() {
	if currentPlayer() == PlayerX {
		// Check for draw
		if checkDraw() {
			fmt.Println("It's a draw!")
			printBoard()
			os.Exit(0)
		}
	}
}
