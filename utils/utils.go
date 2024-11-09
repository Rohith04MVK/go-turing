package utils

import (
	"strings"

	"github.com/Rohith04MVK/turing-machine/config"
)

func ListToString(toStringify []string) string {
	return strings.Join(toStringify, "")
}

func Pipeify(s string) string {
	chars := make([]string, len(s))
	for i, c := range s {
		chars[i] = string(c)
	}
	return strings.Join(chars, "|")
}

func NextIndex(index int, direction string) int {
	if config.IsAllowedMovement(direction) {
		return index + config.TapeMovementFor(direction)
	}
	return index
}

func RemoveEmptyCharacter(dirtyList []string) []string {
	result := make([]string, 0)
	for _, char := range dirtyList {
		if char != config.EmptyCharacter() {
			result = append(result, char)
		}
	}
	return result
}

func CountOccurrences(inList []string) map[string]int {
	cleanList := RemoveEmptyCharacter(inList)
	occurrences := make(map[string]int)

	for _, char := range cleanList {
		occurrences[char]++
	}
	return occurrences
}
