package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Day1Part1 is an implementation of the Algorithm interface
// for the Day 1, Part 1 challenge.
type Day1Part1 struct{}

// Title returns the description of the Day1Part1 algorithm.
func (d Day1Part1) Title() string {
	return "(Day1Part1) Sum of Calibration Values"
}

// Compute executes the main algorithm of Day1Part1.
func (d Day1Part1) Compute(input string) string {
	sum := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		firstDigit := findFirstDigit(line)
		lastDigit := findLastDigit(line)

		// Combine the first and last digits into a two-digit number.
		calibrationValue, err := strconv.Atoi(string([]rune{firstDigit, lastDigit}))
		if err != nil {
			fmt.Println("Error converting digits to integer:", err)
			return "Error"
		}

		sum += calibrationValue
	}

	return fmt.Sprintf("%s: %d", d.Title(), sum)
}

// findFirstDigit finds the first digit in a string.
func findFirstDigit(s string) rune {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return r
		}
	}
	return '0' // Return '0' if no digit is found
}

// findLastDigit finds the last digit in a string.
func findLastDigit(s string) rune {
	var lastDigit rune = '0'
	for _, r := range s {
		if r >= '0' && r <= '9' {
			lastDigit = r
		}
	}
	return lastDigit
}


