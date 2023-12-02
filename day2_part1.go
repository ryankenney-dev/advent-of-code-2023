package main

import (
    "fmt"
    "strconv"
    "strings"
)

type Day2Part1 struct{}

func (d Day2Part1) Title() string {
    return "Sum of IDs of Possible Games"
}

func (d Day2Part1) Compute(input string) string {
    // Parse the input and apply game logic
    sum := 0
    games := strings.Split(input, "\n")
    for _, game := range games {
        gameID, cubes, valid := parseGame(game)
        if valid && isGamePossible(cubes) {
            sum += gameID
        }
    }

    return fmt.Sprintf("%s: %d", d.Title(), sum)
}

// Helper function to parse each game's data
func parseGame(game string) (int, map[string]int, bool) {
    parts := strings.Split(game, ": ")
    if len(parts) != 2 {
        return 0, nil, false
    }

    gameID, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
    if err != nil {
        return 0, nil, false
    }

    cubeSets := strings.Split(parts[1], "; ")
    cubes := make(map[string]int)
    for _, set := range cubeSets {
        for _, cube := range strings.Split(set, ", ") {
            details := strings.Split(cube, " ")
            count, _ := strconv.Atoi(details[0])
            color := details[1]
            if current, ok := cubes[color]; !ok || count > current {
                cubes[color] = count
            }
        }
    }

    return gameID, cubes, true
}

// Helper function to check if a game is possible
func isGamePossible(cubes map[string]int) bool {
    return cubes["red"] <= 12 && cubes["green"] <= 13 && cubes["blue"] <= 14
}
