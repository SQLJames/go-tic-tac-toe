package main

import "fmt"

func getPlayer(turnNumber int) (playerNumber int) {
	if turnNumber%2 == 1 {
		fmt.Println("Player 1's turn")
		return 1
	} else {
		fmt.Println("Player 2's turn")
		return 2
	}
}

func getMove() (moveInt int) {
	fmt.Println("Enter the number where you would like to go.")
	fmt.Scan(&moveInt)

	return moveInt
}
