package day4_part1

import (
	"strconv"
	"strings"
)

type Day4Part1 struct{}

func (d Day4Part1) Title() string {
	return "(Day4Part1) Total Points of Scratchcards"
}

func (d Day4Part1) Compute(input string) string {
	totalPoints := calculateTotalPoints(input)
	return d.Title() + ": " + strconv.Itoa(totalPoints)
}

func calculateTotalPoints(input string) int {
	lines := strings.Split(input, "\n")
	totalPoints := 0

	for _, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue
		}

		winningNumbers := toIntSet(strings.Fields(parts[0]))
		playerNumbers := strings.Fields(parts[1])
		points := 0
		found := false

		for _, num := range playerNumbers {
			if _, ok := winningNumbers[num]; ok {
				if found {
					points *= 2
				} else {
					points = 1
					found = true
				}
			}
		}

		totalPoints += points
	}

	return totalPoints
}

func toIntSet(numbers []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, num := range numbers {
		set[num] = struct{}{}
	}
	return set
}
