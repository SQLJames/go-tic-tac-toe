package main

import "fmt"

func (b *board) validMove(moveInt int) bool {
	if b.playable_field[moveInt] == 0 && moveInt >= 0 && moveInt < len(b.playable_field) {
		return true
	}
	fmt.Println("invalid move, please choose again")
	return false
}

func (b *board) executeMove(player int) {
	moveInt := getMove()
	if !b.validMove(moveInt) {

		b.executeMove(player)
	}
	if player == 1 {
		b.playable_field[moveInt] = player_x
	} else {
		b.playable_field[moveInt] = player_O
	}

}
