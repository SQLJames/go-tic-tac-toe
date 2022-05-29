package main

import (
	"fmt"
	"strings"

	color "github.com/TwiN/go-color"
)

func (b *board) printBoard() {
	rowdivider := " | "
	for i, v := range b.playable_field {
		if v == 0 {
			fmt.Printf("%d", i)
		} else if v == b.player_x {
			fmt.Print(color.Ize(color.Green, "X"))
		} else if v == b.player_O {
			fmt.Print(color.Ize(color.Yellow, "O"))
		}
		// Print the values in columns and rows
		if i > 0 && (i+1)%b.size == 0 {
			fmt.Printf("\n")
			fmt.Println(strings.Repeat("-", len(rowdivider)*b.size))
		} else {
			fmt.Print(rowdivider)
		}
	}

}
