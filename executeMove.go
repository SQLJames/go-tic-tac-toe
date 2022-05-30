package main

import "fmt"

func (b *board) validMove(moveInt int) bool {
	if moveInt < 0 || moveInt > len(b.playable_field) {
		fmt.Println("invalid move, please choose again")
		return false
	}
	if b.playable_field[moveInt] != 0 {
		fmt.Println("That space is already occupied")
		return false
	}

	return true
}

func (b *board) executeMove(moveInt, player int) {
	if player == 1 {
		b.playable_field[moveInt] = b.player_x
	} else {
		b.playable_field[moveInt] = b.player_O
	}

}
