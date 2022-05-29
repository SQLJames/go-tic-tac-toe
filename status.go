package main

import "fmt"

func (b *board) getGameStatus(turnNumber, player int) (gameOver bool) {
	if b.checkforTie(turnNumber) {
		fmt.Println("Game has tied")
		b.printBoard()
		return true
	}
	if b.checkforWin() {
		fmt.Printf("Player %d has Won!\n", player)
		b.printBoard()
		return true
	}
	return false
}
