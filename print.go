package main

import (
	"bytes"
	"fmt"
	"strings"

	color "github.com/TwiN/go-color"
)

func (b *board) printBoard() (buffer *bytes.Buffer) {
	buffer = new(bytes.Buffer)
	rowdivider := " | "
	for i, v := range b.playable_field {
		if v == 0 {
			buffer.WriteString(fmt.Sprintf("%d", i))
		} else if v == b.player_x {
			buffer.WriteString(fmt.Sprint(color.Ize(color.Green, b.player_x_marker)))
		} else if v == b.player_O {
			buffer.WriteString(fmt.Sprint(color.Ize(color.Yellow, b.player_O_marker)))
		}
		// Print the values in columns and rows
		if i > 0 && (i+1)%b.size == 0 {

			buffer.WriteString("\n" + strings.Repeat("-", len(rowdivider)*b.size) + "\n")

		} else {
			buffer.WriteString(rowdivider)
		}
	}
	return buffer
}
