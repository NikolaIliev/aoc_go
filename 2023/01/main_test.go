package main

import "testing"

func TestPart1(t *testing.T) {
	input := `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`
	expected := `142`

	result := part1(input)

	if result != expected {
		t.Errorf("\n\nPart 1:\n  Expected: %q\n  Received: %q\n\n", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := ``
	expected := ``

	result := part2(input)

	if result != expected {
		t.Errorf("\n\nPart 2:\n  Expected: %q\n  Received: %q\n\n", expected, result)
	}
}
