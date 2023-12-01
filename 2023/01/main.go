package main

import (
	"aoc_go/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func part1(input string) string {
	result := 0

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		num := ""
		runes := []rune(line)

		for i := 0; i < len(runes); i++ {
			rune := runes[i]

			if rune >= '0' && rune <= '9' {
				num += string(rune)
				break
			}

		}

		for i := len(runes) - 1; i >= 0; i-- {
			rune := runes[i]

			if rune >= '0' && rune <= '9' {
				num += string(rune)
				break
			}

		}

		result += utils.Unwrap(strconv.Atoi(num))
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	inverse_digits := []string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}

	result := 0

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		num := ""
		buf := ""
		runes := []rune(line)

		for i := 0; i < len(runes); i++ {
			rune := runes[i]

			if rune >= '0' && rune <= '9' {
				num += string(rune)
				break
			}

			buf += string(rune)

			idx := utils.FindIndex(digits, func(digit string) bool {
				return strings.HasSuffix(buf, digit)
			})

			if idx >= 0 {
				num += strconv.Itoa(idx + 1)
				break
			}
		}

		buf = ""

		for i := len(runes) - 1; i >= 0; i-- {
			rune := runes[i]

			if rune >= '0' && rune <= '9' {
				num += string(rune)
				break
			}

			buf += string(rune)

			idx := utils.FindIndex(inverse_digits, func(digit string) bool {
				return strings.HasSuffix(buf, digit)
			})

			if idx >= 0 {
				num += strconv.Itoa(idx + 1)
				break
			}
		}

		result += utils.Unwrap(strconv.Atoi(num))
	}

	return strconv.Itoa(result)
}

func main() {
	input := utils.ReadInput()

	part1_start := time.Now()
	part1_result := part1(input)
	part1_duration := time.Since(part1_start)

	part2_start := time.Now()
	part2_result := part2(input)
	part2_duration := time.Since(part2_start)

	fmt.Println()
	fmt.Printf("Part 1: %s (%v)", part1_result, part1_duration)
	fmt.Println()
	fmt.Printf("Part 2: %s (%v)", part2_result, part2_duration)
}
