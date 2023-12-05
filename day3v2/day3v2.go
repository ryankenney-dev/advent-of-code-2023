package day3v2

import (
    "sort"
    "strconv"
    "strings"
    "unicode"
)

type Day3V2Part1 struct{}

func (d Day3V2Part1) Title() string {
    return "(Day3V2Part1) Sum of Part Numbers"
}

type Coordinate struct {
    X, Y int
}

type PartNumber struct {
    StartCoordinate Coordinate
    Value           int
}

func (d Day3V2Part1) Compute(input string) string {
    grid := parseInput(input)
    sum := execute(grid)
    return d.Title() + ": " + strconv.Itoa(sum)
}

func execute(grid [][]rune) int {
    numbersIndex := findIntegersAsPotentialPartNumbers(grid)
    symbolCoordinates := collectSymbolCoordinates(grid)
    partNumbers := getNumbersAdjacentToSymbols(symbolCoordinates, numbersIndex)

    // Sum the partNumbers
    sum := 0
    for _, partNumber := range partNumbers {
        sum += partNumber.Value
    }

    return sum
}

func parseInput(input string) [][]rune {
    lines := strings.Split(input, "\n")
    grid := make([][]rune, len(lines))
    for i, line := range lines {
        grid[i] = []rune(line)
    }
    return grid
}

func findIntegersAsPotentialPartNumbers(grid [][]rune) map[Coordinate]*PartNumber {
    numbersIndex := make(map[Coordinate]*PartNumber)
    for y, row := range grid {
        number := ""
        startX := 0 // To keep track of the start of the number
        for x, char := range row {
            if unicode.IsDigit(char) {
                if number == "" {
                    startX = x
                }
                number += string(char)
            } else if number != "" {
                if len(number) >= 1 {
                    num, _ := strconv.Atoi(number)
                    partNumber := &PartNumber{
                        StartCoordinate: Coordinate{X: startX, Y: y},
                        Value:           num,
                    }
                    for i := 0; i < len(number); i++ {
                        numbersIndex[Coordinate{X: startX + i, Y: y}] = partNumber
                    }
                }
                number = ""
            }
        }
        if number != "" {
            num, _ := strconv.Atoi(number)
            partNumber := &PartNumber{
                StartCoordinate: Coordinate{X: startX, Y: y},
                Value:           num,
            }
            for i := 0; i < len(number); i++ {
                numbersIndex[Coordinate{X: startX + i, Y: y}] = partNumber
            }
        }
    }
    return numbersIndex
}

func collectSymbolCoordinates(grid [][]rune) map[Coordinate]struct{} {
    symbolCoordinates := make(map[Coordinate]struct{})
    for y, row := range grid {
        for x, char := range row {
            if !unicode.IsDigit(char) && char != '.' {
                symbolCoordinates[Coordinate{X: x, Y: y}] = struct{}{}
            }
        }
    }
    return symbolCoordinates
}

func getNumbersAdjacentToSymbols(symbolCoordinates map[Coordinate]struct{}, numbersIndex map[Coordinate]*PartNumber) []*PartNumber {
    partNumbersSet := make(map[*PartNumber]struct{})
    var partNumbers []*PartNumber

    directions := []Coordinate{
        {0, 1}, {1, 0}, {0, -1}, {-1, 0},
        {1, 1}, {1, -1}, {-1, 1}, {-1, -1},
    }

    for coord := range symbolCoordinates {
        for _, dir := range directions {
            adjCoord := Coordinate{X: coord.X + dir.X, Y: coord.Y + dir.Y}
            if partNumber, exists := numbersIndex[adjCoord]; exists {
                if _, alreadyAdded := partNumbersSet[partNumber]; !alreadyAdded {
                    partNumbers = append(partNumbers, partNumber)
                    partNumbersSet[partNumber] = struct{}{}
                }
            }
        }
    }

    // Sort partNumbers by their StartCoordinate
    sort.Slice(partNumbers, func(i, j int) bool {
        if partNumbers[i].StartCoordinate.Y == partNumbers[j].StartCoordinate.Y {
            return partNumbers[i].StartCoordinate.X < partNumbers[j].StartCoordinate.X
        }
        return partNumbers[i].StartCoordinate.Y < partNumbers[j].StartCoordinate.Y
    })

    return partNumbers
}