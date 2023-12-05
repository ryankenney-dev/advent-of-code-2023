package day4

import (
	"testing"
)

func TestCompute(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Example Case",
			input: `41 48 83 86 17 | 83 86  6 31 17  9 48 53
					13 32 20 16 61 | 61 30 68 82 17 32 24 19
					1 21 53 59 44 | 69 82 63 72 16 21 14  1
					41 92 73 84 69 | 59 84 76 51 58  5 54 83
					87 83 26 28 32 | 88 30 70 12 93 22 82 36
					31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
			expected: "(Day4Part1) Total Points of Scratchcards: 13",
		},
	}

	algorithm := Day4Part1{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := algorithm.Compute(tc.input)
			if result != tc.expected {
				t.Errorf("Test %s failed. Expected %s, got %s", tc.name, tc.expected, result)
			}
		})
	}
}
