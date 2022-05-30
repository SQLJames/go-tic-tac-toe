package main

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	t.Run("game size", func(t *testing.T) {
		cases := []struct {
			desc     string
			size     int
			expected int
		}{
			{"3x3", 3, 9},
			{"4x4", 4, 16},
			{"5x5", 5, 25},
			{"6x6", 6, 36},
			{"7x7", 7, 49},
			{"8x8", 8, 64},
		}
		for _, tc := range cases {
			actual := newBoard(tc.size)
			if len(actual.playable_field) != tc.expected {
				
				t.Fatalf("desc %s: expected: %d got: %d sizeProvided %d",
					tc.desc, tc.expected, actual.playable_field, tc.size)
			}
		}
	})

}
