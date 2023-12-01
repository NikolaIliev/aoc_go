package main

import (
	"aoc_go/utils"
	"fmt"
	"time"
)

func part1(input string) string {
	return ""
}

func part2(input string) string {
	return ""
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
