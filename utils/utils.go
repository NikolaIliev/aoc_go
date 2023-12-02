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

func Every[T any](slice []T, predicate func(T) bool) bool {
	for _, value := range slice {
		if !predicate(value) {
			return false
		}
	}

	return true
}

func Map[T any, R any](slice []T, mapper func(T) R) []R {
	result := []R{}

	for _, value := range slice {
		result = append(result, mapper(value))
	}

	return result
}

func Reduce[T any, R any](slice []T, reducer func(R, T) R, initializer R) R {
	acc := initializer

	for _, v := range slice {
		acc = reducer(acc, v)
	}

	return acc
}

func ReadInput() string {
	return strings.Trim(string(Unwrap(os.ReadFile("input.txt"))), "\n")
}
