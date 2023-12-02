package main

import "fmt"

// CharactersAlgorithm implements the Algorithm interface for counting characters
type CharactersAlgorithm struct{}

// Title returns the title of the algorithm
func (ca CharactersAlgorithm) Title() string {
    return "Number of Characters"
}

// Compute performs the computation of CharactersAlgorithm
func (ca CharactersAlgorithm) Compute(input string) string {
    return fmt.Sprintf("Number of characters: %d", len(input))
}

