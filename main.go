package main

import "fmt"

var (
	size int = 3
)

func main() {
	var gameOver bool
	turnNumber := 1
	board := newBoard(size)
	for !gameOver {
		fmt.Print(board.printBoard())
		player := board.getPlayer(turnNumber)
		selectedSquare := board.getMove()
		board.executeMove(selectedSquare, player)
		gameOver = board.getGameStatus(turnNumber, player)
		turnNumber++
	}
	fmt.Print(board.printBoard())
}
