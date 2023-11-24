package main

import "testing"

func TestCheckWinner(t *testing.T) {
	tests := []struct {
		player   string
		opponent string
		expected string
	}{
		{"Human", "Ant", "you win!"},
		{"Elephant", "Human", "you win!"},
		{"Ant", "Elephant", "you win!"},
		{"Ant", "Human", "you lose!"},
		{"Human", "Elephant", "you lose!"},
		{"Elephant", "Ant", "you lose!"},
		{"Human", "Human", "tie!"},
		{"Elephant", "Elephant", "tie!"},
		{"Ant", "Ant", "tie!"},
	}

	for _, test := range tests {
		result := checkWinner(test.player, test.opponent)
		if result != test.expected {
			t.Errorf("expected: %s, got: %s\n", test.expected, result)
		}
	}
}
