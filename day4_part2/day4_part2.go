package day4_part2

import (
	"strconv"
	"strings"
)

type Day4Part2 struct{}

func (d Day4Part2) Title() string {
	return "(Day4Part2) Total Number of Scratchcards"
}

func (d Day4Part2) Compute(input string) string {
    cards := parseCards(input)
    totalCards := calculateTotalCards(cards)
    return d.Title() + ": " + strconv.Itoa(totalCards)
}

func calculateTotalCards(cards []struct{
    winningNumbers map[string]struct{}
    playerNumbers  map[string]struct{}
}) int {
    cardCounts := make([]int, len(cards))
    for i := range cardCounts {
        cardCounts[i] = 1 // Initialize with the original card
    }

    for i := 0; i < len(cards); i++ {
        matches := countMatches(cards[i].winningNumbers, cards[i].playerNumbers)

        for j := 1; j <= matches; j++ {
            if i+j < len(cards) {
                cardCounts[i+j] += cardCounts[i]
            }
        }
    }

    totalCards := 0
    for _, count := range cardCounts {
        totalCards += count
    }

    return totalCards
}

func parseCards(input string) []struct {
    winningNumbers map[string]struct{}
    playerNumbers  map[string]struct{}
} {
    lines := strings.Split(input, "\n")
    var cards []struct {
        winningNumbers map[string]struct{}
        playerNumbers  map[string]struct{}
    }

    for _, line := range lines {
        parts := strings.Split(line, "|")
        if len(parts) != 2 {
            continue
        }

        winningNumbers := strings.Fields(parts[0])
        playerNumbers := strings.Fields(parts[1])

        winningMap := make(map[string]struct{})
        playerMap := make(map[string]struct{})

        for _, num := range winningNumbers {
            winningMap[num] = struct{}{}
        }
        for _, num := range playerNumbers {
            playerMap[num] = struct{}{}
        }

        cards = append(cards, struct {
            winningNumbers map[string]struct{}
            playerNumbers  map[string]struct{}
        }{winningNumbers: winningMap, playerNumbers: playerMap})
    }

    return cards
}

func countMatches(winningNumbers, playerNumbers map[string]struct{}) int {
    count := 0
    for number := range playerNumbers {
        if _, found := winningNumbers[number]; found {
            count++
        }
    }
    return count
}

func toIntSet(numbers []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, num := range numbers {
		set[num] = struct{}{}
	}
	return set
}
