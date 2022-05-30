package main

import (
	"testing"
)

func TestValidMove(t *testing.T) {
	t.Run("Space already occupied", func(t *testing.T) {
		board := newBoard(3)
		for chosenSpace := range board.playable_field {
			b := newBoard(3)
			b.playable_field[chosenSpace] = b.player_O
			move := b.validMove(chosenSpace)

			if move {
				t.Errorf("move valid: %t on Space %d", move, chosenSpace)
			}
		}
	})
	t.Run("Out of Range of playable field", func(t *testing.T) {
		board := newBoard(3)
		cases := []struct {
			desc      string
			testSpace int
			expected  bool
		}{
			{"negative number", -1, false},
			{"larger than field", 10, false},
		}
		for _, tc := range cases {
			actual := board.validMove(tc.testSpace)
			if actual != tc.expected {
				t.Fatalf("desc %s: expected: %t got: %t SpaceSelected %d",
					tc.desc, tc.expected, actual, tc.testSpace)
			}
		}
	})
	t.Run("In range not occupied", func(t *testing.T) {
		board := newBoard(3)
		for chosenSpace := range board.playable_field {
			move := board.validMove(chosenSpace)
			if !move {
				t.Errorf("move valid: %t on Space %d", move, chosenSpace)
			}
		}
	})
}

func TestExecuteMove(t *testing.T) {
	t.Run("Move is marked", func(t *testing.T) {
		board := newBoard(3)
		cases := []struct {
			desc         string
			MoveInt      int
			Player       int
			PlayerNumber int
		}{
			{"Player 1", 1, 1, board.player_x},
			{"Player 2", 2, 2, board.player_O},
		}
		for _, tc := range cases {
			board.executeMove(tc.MoveInt, tc.Player)
			if board.playable_field[tc.MoveInt] != tc.PlayerNumber {
				t.Fatalf("desc %s: expected: %d got: %d SpaceSelected %d",
					tc.desc, tc.PlayerNumber, board.playable_field[tc.MoveInt], tc.MoveInt)
			}
		}
	})
}
