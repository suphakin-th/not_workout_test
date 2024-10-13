package tests

import (
	"not_workout/cmd/permutations"
	"reflect"
	"testing"
)

func TestPermutationsFn(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}

	for _, test := range tests {
		result := permutations.PermutationsFn(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %s, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
