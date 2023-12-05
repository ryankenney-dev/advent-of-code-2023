package day4_part2

import (
	"reflect"
	"testing"
)

func TestParseCards(t *testing.T) {
    input := `41 48 83 86 17 | 83 86  6 31 17  9 48 53
              13 32 20 16 61 | 61 30 68 82 17 32 24 19`

    expected := []struct {
        winningNumbers map[string]struct{}
        playerNumbers  map[string]struct{}
    }{
        {
            winningNumbers: map[string]struct{}{"41": {}, "48": {}, "83": {}, "86": {}, "17": {}},
            playerNumbers:  map[string]struct{}{"83": {}, "86": {}, "6": {}, "31": {}, "17": {}, "9": {}, "48": {}, "53": {}},
        },
        {
            winningNumbers: map[string]struct{}{"13": {}, "32": {}, "20": {}, "16": {}, "61": {}},
            playerNumbers:  map[string]struct{}{"61": {}, "30": {}, "68": {}, "82": {}, "17": {}, "32": {}, "24": {}, "19": {}},
        },
    }

    cards := parseCards(input)

    if !reflect.DeepEqual(cards, expected) {
        t.Errorf("parseCards() = %v, want %v", cards, expected)
    }
}

func TestCountMatches(t *testing.T) {
	winningNumbers := map[string]struct{}{
		"41": {}, "48": {}, "83": {}, "86": {}, "17": {},
	}
	playerNumbers := map[string]struct{}{
		"83": {}, "86": {}, "6": {}, "31": {}, "17": {}, "9": {}, "48": {}, "53": {},
	}

	expected := 4 // Number of matches in the card
	result := countMatches(winningNumbers, playerNumbers)

	if result != expected {
		t.Errorf("countMatches() = %d, want %d", result, expected)
	}
}

func TestCalculateTotalCards(t *testing.T) {
    cards := []struct {
        winningNumbers map[string]struct{}
        playerNumbers  map[string]struct{}
    }{
        {
            winningNumbers: map[string]struct{}{"41": {}, "48": {}, "83": {}, "86": {}, "17": {}},
            playerNumbers:  map[string]struct{}{"83": {}, "86": {}, "6": {}, "31": {}, "17": {}, "9": {}, "48": {}, "53": {}},
        },
        {
            winningNumbers: map[string]struct{}{"13": {}, "32": {}, "20": {}, "16": {}, "61": {}},
            playerNumbers:  map[string]struct{}{"61": {}, "30": {}, "68": {}, "82": {}, "17": {}, "32": {}, "24": {}, "19": {}},
        },
        {
            winningNumbers: map[string]struct{}{"1": {}, "21": {}, "53": {}, "59": {}, "44": {}},
            playerNumbers:  map[string]struct{}{"69": {}, "82": {}, "63": {}, "72": {}, "16": {}, "21": {}, "14": {}, "1": {}},
        },
        {
            winningNumbers: map[string]struct{}{"41": {}, "92": {}, "73": {}, "84": {}, "69": {}},
            playerNumbers:  map[string]struct{}{"59": {}, "84": {}, "76": {}, "51": {}, "58": {}, "5": {}, "54": {}, "83": {}},
        },
        {
            winningNumbers: map[string]struct{}{"87": {}, "83": {}, "26": {}, "28": {}, "32": {}},
            playerNumbers:  map[string]struct{}{"88": {}, "30": {}, "70": {}, "12": {}, "93": {}, "22": {}, "82": {}, "36": {}},
        },
        {
            winningNumbers: map[string]struct{}{"31": {}, "18": {}, "13": {}, "56": {}, "72": {}},
            playerNumbers:  map[string]struct{}{"74": {}, "77": {}, "10": {}, "23": {}, "35": {}, "67": {}, "36": {}, "11": {}},
        },
    }

    expected := 30 // Expected total cards
    total := calculateTotalCards(cards)

    if total != expected {
        t.Errorf("calculateTotalCards() = %d, want %d", total, expected)
    }
}

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
			expected: "Total Number of Scratchcards: 30",
		},
	}

	algorithm := Day4Part2{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := algorithm.Compute(tc.input)
			if result != tc.expected {
				t.Errorf("Test %s failed. Expected %s, got %s", tc.name, tc.expected, result)
			}
		})
	}
}
