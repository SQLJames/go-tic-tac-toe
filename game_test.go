package main

import "testing"

func TestGetPlayer(t *testing.T) {
	t.Run("Player Turn", func(t *testing.T) {
		board := newBoard(3)
		cases := []struct {
			desc   string
			Turn   int
			Player int
		}{
			{"Player 1", 1, 1},
			{"Player 2", 2, 2},
			{"Player 1", 3, 1},
			{"Player 2", 4, 2},
			{"Player 1", 5, 1},
			{"Player 2", 6, 2},
			{"Player 1", 7, 1},
			{"Player 2", 8, 2},
		}
		for _, tc := range cases {
			p := board.getPlayer(tc.Turn)
			if p != tc.Player {
				t.Fatalf("TurnNumber %d: Player Expected: %d Player Got: %d",
					tc.Turn, tc.Player, p)
			}
		}
	})
}
