package main

import (
	"aoc_go/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

type round struct {
	Red   int
	Green int
	Blue  int
}

type game struct {
	Id     int
	Rounds []round
}

// 3 blue, 4 red
// 1 red, 2 green, 6 blue
// 12 green
func roundFromString(s string) round {
	r := round{0, 0, 0}

	for _, cubeString := range strings.Split(s, ", ") {
		countString, color, _ := strings.Cut(cubeString, " ")

		count := utils.Unwrap(strconv.Atoi(countString))

		switch color[0] {
		case 'r':
			r.Red = count
		case 'g':
			r.Green = count
		case 'b':
			r.Blue = count
		}
	}

	return r
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func gameFromString(s string) game {
	g := game{0, []round{}}

	left, right, _ := strings.Cut(s, ": ")

	_, idString, _ := strings.Cut(left, " ")

	g.Id = utils.Unwrap(strconv.Atoi(idString))

	for _, roundString := range strings.Split(right, "; ") {
		g.Rounds = append(g.Rounds, roundFromString(roundString))
	}

	return g
}

func part1(input string) string {
	maxRound := round{12, 13, 14}
	result := 0

	for _, line := range strings.Split(input, "\n") {
		g := gameFromString(line)

		if utils.Every(g.Rounds, func(r round) bool {
			return r.Red <= maxRound.Red && r.Green <= maxRound.Green && r.Blue <= maxRound.Blue
		}) {
			result += g.Id
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	result := 0

	for _, line := range strings.Split(input, "\n") {
		g := gameFromString(line)

		maxRound := round{
			Red:   slices.Max(utils.Map(g.Rounds, func(r round) int { return r.Red })),
			Green: slices.Max(utils.Map(g.Rounds, func(r round) int { return r.Green })),
			Blue:  slices.Max(utils.Map(g.Rounds, func(r round) int { return r.Blue })),
		}

		result += maxRound.Red * maxRound.Green * maxRound.Blue
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
