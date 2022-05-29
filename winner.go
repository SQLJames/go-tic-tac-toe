package main

func (b *board) checkforWin() (won bool) {
	sums := make([]int, 8)
	//0 | 1 | 2
	//---------
	//3 | 4 | 5
	//---------
	//6 | 7 | 8
	//---------
	//get column totals
	sums[0] = b.playable_field[0] + b.playable_field[3] + b.playable_field[6]
	sums[1] = b.playable_field[1] + b.playable_field[4] + b.playable_field[7]
	sums[2] = b.playable_field[2] + b.playable_field[5] + b.playable_field[8]
	//get row totals
	sums[3] = b.playable_field[0] + b.playable_field[1] + b.playable_field[2]
	sums[4] = b.playable_field[3] + b.playable_field[4] + b.playable_field[5]
	sums[5] = b.playable_field[6] + b.playable_field[7] + b.playable_field[8]
	//get diag totals
	sums[6] = b.playable_field[0] + b.playable_field[4] + b.playable_field[8]
	sums[7] = b.playable_field[2] + b.playable_field[4] + b.playable_field[6]

	for _, v := range sums {
		if v == player_O*3 {
			return true
		}
		if v == player_x*3 {
			return true
		}

	}
	return false
}

func (b *board) checkforTie(turn int) bool {
	return turn == 9
}
