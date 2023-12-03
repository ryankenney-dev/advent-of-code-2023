package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day1Part2 is an implementation of the Algorithm interface
// for the Day 1, Part 2 challenge.
type Day1Part2 struct{}

// Title returns the description of the Day1Part2 algorithm.
func (d Day1Part2) Title() string {
	return "(Day1Part2) Sum of Real Calibration Values"
}

// Compute executes the main algorithm of Day1Part2.
func (d Day1Part2) Compute(input string) string {
	sum := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		firstDigit := findFirstRealDigit(line)
		lastDigit := findLastRealDigit(line)

		// Combine the first and last digits into a two-digit number.
		calibrationValue, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		if err != nil {
			fmt.Println("Error converting digits to integer:", err)
			return "Error"
		}

		sum += calibrationValue
	}

	return fmt.Sprintf("%s: %d", d.Title(), sum)
}

// digitMap maps spelled-out digits to their numerical values.
var digitMap = map[string]int{
	"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4,
	"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

// findFirstRealDigit finds the first digit (numerical or spelled-out) in a string.
func findFirstRealDigit(s string) int {
	for i := 0; i < len(s); i++ {
		for word, value := range digitMap {
			if strings.HasPrefix(s[i:], word) {
				return value
			}
		}
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}
	}
	return 0
}

// findLastRealDigit finds the last digit (numerical or spelled-out) in a string.
func findLastRealDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		for word, value := range digitMap {
			if i >= len(word)-1 && strings.HasPrefix(s[i-len(word)+1:], word) {
				return value
			}
		}
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}
	}
	return 0
}

