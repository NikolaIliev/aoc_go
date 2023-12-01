package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := strings.TrimSpace(`
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`)
	expected := strings.TrimSpace(`142`)

	result := part1(input)

	if result != expected {
		t.Errorf("\n\nPart 1:\n  Expected: %q\n  Received: %q\n\n", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := strings.TrimSpace(`
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`)
	expected := strings.TrimSpace(`281`)

	result := part2(input)

	if result != expected {
		t.Errorf("\n\nPart 2:\n  Expected: %q\n  Received: %q\n\n", expected, result)
	}
}
