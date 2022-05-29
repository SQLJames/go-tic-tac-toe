package main

type board struct {
	columns        int
	rows           int
	playable_field []int
}

func newBoard(columns, rows int) (b *board) {
	b = &board{
		columns:        columns,
		rows:           rows,
		playable_field: make([]int, rows*columns),
	}
	return b
}
