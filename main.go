package main

var (
	columns, rows int = 3, 3
	player_x          = 10
	player_O          = 100
)

func main() {
	gameOver := false
	turnNumber := 1

	board := newBoard(columns, rows)
	for !gameOver {
		board.printBoard()
		player := getPlayer(turnNumber)
		board.executeMove(player)
		gameOver = board.getGameStatus(turnNumber, player)
		turnNumber++

	}

}
