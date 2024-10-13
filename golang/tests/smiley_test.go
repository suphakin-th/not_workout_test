package tests

import (
	"not_workout/cmd/smiley"
	"testing"
)

func TestCountSmileys(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{":)", ";(", ";}", ":-D"}, 2},
		{[]string{";D", ":-(", ":-)", ";~)"}, 3},
		{[]string{";]", ":[", ";*", ":$", ";-D"}, 1},
		{[]string{}, 0},                         // Edge case: empty array
		{[]string{":", ";", "-", "~"}, 0},       // No valid smileys
		{[]string{":D", ";-D", ":~)", ";)"}, 4}, // Multiple valid smileys
	}

	for _, test := range tests {
		result := smiley.CountSmileys(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
