package main

import (
    "strconv"
    "strings"
)

type Day3Part1 struct{}

func (d Day3Part1) Title() string {
    return "Sum of Part Numbers"
}

// Position represents a unique position in the grid
type Position struct {
    Row, Col int
}

func (d Day3Part1) Compute(input string) string {
    grid := parseInput(input)
    sum := 0
    visited := make(map[Position]bool) // To track visited multi-digit numbers uniquely by position
    for i := range grid {
        for j := 0; j < len(grid[i]); {
            if isNumber(grid[i][j]) {
                num, endIndex := extractNumber(grid[i], j)
                pos := Position{Row: i, Col: endIndex}
                if _, found := visited[pos]; !found {
                    if isAdjacentToSymbol(grid, i, j, endIndex) {
                        number, _ := strconv.Atoi(num)
                        sum += number
                        visited[pos] = true
                    }
                }
                j = endIndex + 1 // Skip to the end of the current number
            } else {
                j++ // Move to next cell
            }
        }
    }
    return d.Title() + ": " + strconv.Itoa(sum)
}

func parseInput(input string) [][]rune {
    lines := strings.Split(input, "\n")
    grid := make([][]rune, len(lines))
    for i, line := range lines {
        grid[i] = []rune(line)
    }
    return grid
}

func isNumber(c rune) bool {
    return c >= '0' && c <= '9'
}

func extractNumber(row []rune, startIndex int) (string, int) {
    var number strings.Builder
    i := startIndex
    for i < len(row) && isNumber(row[i]) {
        number.WriteRune(row[i])
        i++
    }
    return number.String(), i - 1
}

func isAdjacentToSymbol(grid [][]rune, i, startJ, endJ int) bool {
    directions := [...]struct{ x, y int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
    for _, dir := range directions {
        for j := startJ; j <= endJ; j++ {
            newX, newY := i+dir.x, j+dir.y
            inBounds := newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[newX])
            adjacentToSymbol := inBounds && grid[newX][newY] != '.' && !isNumber(grid[newX][newY])
            if adjacentToSymbol {
                return true
            }
        }
    }
    return false
}
