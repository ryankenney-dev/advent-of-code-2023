package main

import "fmt"
import "strings"

// LinesAlgorithm implements the Algorithm interface for counting lines
type LinesAlgorithm struct{}

// Title returns the title of the algorithm
func (la LinesAlgorithm) Title() string {
    return "Number of Lines"
}

// Compute performs the computation of LinesAlgorithm
func (la LinesAlgorithm) Compute(input string) string {
    return fmt.Sprintf("Number of lines: %d", len(strings.Split(input, "\n")))
}


