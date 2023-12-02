package main

// Algorithm defines the interface for different computations
type Algorithm interface {
    Title() string
    Compute(input string) string
}


