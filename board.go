package main

type board struct {
	size            int
	playable_field  []int
	player_O        int
	player_x        int
	player_O_marker string
	player_x_marker string
	winningTotalO   int
	winningTotalX   int
}

func newBoard(s int) (b *board) {
	x := 111
	o := 1000
	b = &board{
		size:            s,
		playable_field:  make([]int, s*s),
		player_O:        x,
		player_x:        o,
		player_O_marker: "X",
		player_x_marker: "O",
		winningTotalO:   o * s,
		winningTotalX:   x * s,
	}
	return b
}
