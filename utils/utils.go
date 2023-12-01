package utils

import (
	"os"
	"strings"
)

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

func FindIndex[T comparable](slice []T, predicate func(T) bool) int {
	for idx, value := range slice {
		if predicate(value) {
			return idx
		}
	}

	return -1
}

func ReadInput() string {
	return strings.Trim(string(Unwrap(os.ReadFile("input.txt"))), "\n")
}
