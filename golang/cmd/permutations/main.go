package permutations

import (
	"sort"
)

func PermutationsFn(s string) []string {
	result := make(map[string]bool)
	var permute func([]rune, int)

	permute = func(runes []rune, l int) {
		if l == len(runes) {
			result[string(runes)] = true
			return
		}
		for i := l; i < len(runes); i++ {
			runes[l], runes[i] = runes[i], runes[l]
			permute(runes, l+1)
			runes[l], runes[i] = runes[i], runes[l]
		}
	}
	permute([]rune(s), 0)

	// Convert the map keys to a sorted slice
	var uniquePermutations []string
	for perm := range result {
		uniquePermutations = append(uniquePermutations, perm)
	}
	sort.Strings(uniquePermutations)
	return uniquePermutations
}
