package main

var (
	size int = 3
)

func main() {
	gameOver := false
	turnNumber := 1
	board := newBoard(size)
	for !gameOver {
		board.printBoard()
		player := getPlayer(turnNumber)
		selectedSquare := board.getMove()
		board.executeMove(selectedSquare, player)
		if turnNumber >= 5 {
			gameOver = board.getGameStatus(turnNumber, player)
		}
		turnNumber++
	}

}
