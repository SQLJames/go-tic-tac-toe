package main

func (b *board) checkforWin() (won bool) {

	row := b.calcStraightWinner()
	diag := b.calcDiagWinner()

	if row || diag {
		return true
	}
	return false
}

func (b *board) checkforTie(turn int) bool {
	return turn == b.size*b.size
}

func (b *board) calcStraightWinner() (winner bool) {
	var totalRow, totalColumn int
	for rowNumber := 0; rowNumber < b.size; rowNumber++ {

		for columnNumber, _ := range make([]int, b.size) {
			totalRow += b.playable_field[columnNumber+(rowNumber*b.size)]
			totalColumn += b.playable_field[rowNumber+(columnNumber*b.size)]
		}
		if totalRow == b.winningTotalO || totalRow == b.winningTotalX || totalColumn == b.winningTotalO || totalColumn == b.winningTotalX {
			return true
		}
		totalRow, totalColumn = 0, 0

	}
	return false
}

func (b *board) calcDiagWinner() (winner bool) {
	var totalDownToRight, totalDownToLeft []int
	index := 0
	for row := 0; row < b.size; row++ {
		for column := 0; column < b.size; column++ {
			//fmt.Printf("Column: %d Row: %d Index: %d \n", column, row, index)
			if (row*b.size)+row == index {
				totalDownToRight = append(totalDownToRight, b.playable_field[index])
				//fmt.Println("adding to totalDownToRight")
			}
			if (row+1)*b.size-1*(row+1) == index {
				//fmt.Println("adding to totalDownToLeft")
				totalDownToLeft = append(totalDownToLeft, b.playable_field[index])
			}
			index++
		}
	}

	//fmt.Printf("totalDownToRight: %d totalDownToLeft: %d \n", totalDownToRight, totalDownToLeft)
	if sumSlice(totalDownToLeft) == b.winningTotalO || sumSlice(totalDownToLeft) == b.winningTotalX {
		return true
	}
	if sumSlice(totalDownToRight) == b.winningTotalO || sumSlice(totalDownToRight) == b.winningTotalX {
		return true
	}
	return false
}

func sumSlice(slice []int) (result int) {
	for _, v := range slice {
		result += v
	}
	return result
}
