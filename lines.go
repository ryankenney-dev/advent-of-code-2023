package main

import "strings"

// CountLines counts the number of lines in the given text.
func CountLines(text string) int {
    return len(strings.Split(text, "\n"))
}


