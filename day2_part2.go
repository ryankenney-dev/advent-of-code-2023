package main

import (
    "fmt"
    "strconv"
    "strings"
)

type Day2Part2 struct{}

func (d Day2Part2) Title() string {
    return "(Day2Part2) Sum of Power of Minimum Sets"
}

func (d Day2Part2) Compute(input string) string {
    sum := 0
    games := strings.Split(input, "\n")
    for _, game := range games {
        _, cubes := parseGame2(game)
        sum += powerOfSet(cubes)
    }

    return fmt.Sprintf("%s: %d", d.Title(), sum)
}

// Helper function to parse each game's data
func parseGame2(game string) (int, map[string]int) {
    parts := strings.Split(game, ": ")
    gameID, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
    if err != nil {
        return 0, nil
    }

    cubeSets := strings.Split(parts[1], "; ")
    cubes := map[string]int{
        "red":   0,
        "green": 0,
        "blue":  0,
    }
    for _, set := range cubeSets {
        for _, cube := range strings.Split(set, ", ") {
            details := strings.Split(cube, " ")
            count, _ := strconv.Atoi(details[0])
            color := details[1]
            if count > cubes[color] {
                cubes[color] = count
            }
        }
    }

    return gameID, cubes
}

// Helper function to calculate the power of a set of cubes
func powerOfSet(cubes map[string]int) int {
    red := cubes["red"]
    green := cubes["green"]
    blue := cubes["blue"]
    return red * green * blue
}
