package day3

import (
    "strconv"
    "strings"
)

type Day3Part1 struct{}

type Day3Part2 struct{}

func (d Day3Part1) Title() string {
    return "(Day3Part1) Sum of Part Numbers"
}

func (d Day3Part2) Title() string {
    return "(Day3Part2) Sum of Gear Ratios"
}

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

func (d Day3Part2) Compute(input string) string {
    grid := parseInput(input)
    sum := 0
    for i, row := range grid {
        for j, cell := range row {
            if cell == '*' {
                ratios, ok := findGearRatios(grid, i, j)
                if ok {
                    product := 1
                    for _, ratio := range ratios {
                        product *= ratio
                    }
                    sum += product
                }
            }
        }
    }
    return d.Title() + ": " + strconv.Itoa(sum)
}

func findGearRatios(grid [][]rune, i, j int) ([]int, bool) {
    var ratios []int
    var processedPositions map[Position]bool = make(map[Position]bool) // To track processed numbers
    directions := [...]struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
    for _, dir := range directions {
        newX, newY := i+dir.x, j+dir.y
        if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[newX]) && isNumber(grid[newX][newY]) {
            startY := findStartOfHorizontalNumber(grid[newX], newY)
            pos := Position{Row: newX, Col: startY}
            if !processedPositions[pos] {
                num, _ := extractNumber(grid[newX], startY)
                number, _ := strconv.Atoi(num)
                ratios = append(ratios, number)
                processedPositions[pos] = true
            }
        }
    }
    return ratios, len(ratios) == 2
}

func findStartOfHorizontalNumber(row []rune, endIndex int) int {
    startIndex := endIndex
    for startIndex > 0 && isNumber(row[startIndex-1]) {
        startIndex--
    }
    return startIndex
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
