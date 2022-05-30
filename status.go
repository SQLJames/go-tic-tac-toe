package main

import "fmt"

func (b *board) getGameStatus(turnNumber, player int) (gameOver bool) {
	if b.checkforWin() {
		fmt.Printf("Player %d has Won!\n", player)
		gameOver = true
	}
	if b.checkforTie(turnNumber) {
		fmt.Println("Game has tied")
		gameOver = true
	}
	return gameOver
}
