package smiley

import (
	"regexp"
)

func CountSmileys(arr []string) int {
	// Old school, old ages but cool!
	smileyRegex := regexp.MustCompile(`^(?:[:;])(?:[-~]?)(?:[)D])$`)
	count := 0

	for _, face := range arr {
		if smileyRegex.MatchString(face) {
			count++
		}
	}
	return count
}
