package utils

import (
	"os"
)

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

func ReadInput() string {
	return string(Unwrap(os.ReadFile("input.txt")))
}
