package day5

import (
    "fmt"
    "strconv"
    "strings"
)

type Day5Part1 struct {
    mappings map[string]map[int]int
    seeds    []int
}

func NewDay5Part1() *Day5Part1 {
    return &Day5Part1{
        mappings: make(map[string]map[int]int),
    }
}

func (d *Day5Part1) Title() string {
    return "(Day5Part1) Lowest Location Number for Seeds"
}

func (d *Day5Part1) Compute(input string) string {
    d.parseAlmanac(input)
    lowestLocation := d.findLowestLocation()
    return fmt.Sprintf("%s: %d", d.Title(), lowestLocation)
}

func (d *Day5Part1) parseAlmanac(almanac string) {
    d.mappings = make(map[string]map[int]int)
    d.seeds = nil
    sections := strings.Split(almanac, "\n\n") // Assuming each section is separated by a blank line

    for _, section := range sections {
        lines := strings.Split(section, "\n")
        if strings.Contains(lines[0], "seeds:") {
            // Parse initial seeds
            seedStrs := strings.Fields(lines[0])[1:]
            for _, s := range seedStrs {
                seed, err := strconv.Atoi(s)
                if err != nil {
                    fmt.Printf("Error parsing seed: %s\n", err)
                    continue
                }
                d.seeds = append(d.seeds, seed)
            }
        } else {
            // Parse mapping table
            categoryName := strings.TrimSuffix(lines[0], " map:")
            d.mappings[categoryName] = make(map[int]int)
            for _, line := range lines[1:] {
                parts := strings.Fields(line)
                if len(parts) != 3 {
                    fmt.Printf("Invalid mapping line: %s\n", line)
                    continue
                }
                destStart, _ := strconv.Atoi(parts[0])
                sourceStart, _ := strconv.Atoi(parts[1])
                rangeLength, _ := strconv.Atoi(parts[2])

                for i := 0; i < rangeLength; i++ {
                    d.mappings[categoryName][sourceStart+i] = destStart + i
                }
            }
        }
    }
}

func (d *Day5Part1) findLowestLocation() int {
    lowestLocation := -1

    for _, seed := range d.seeds {
        location := d.convertThroughStages(seed)
        if lowestLocation == -1 || location < lowestLocation {
            lowestLocation = location
        }
    }

    return lowestLocation
}

func (d *Day5Part1) convertThroughStages(initial int) int {
    stages := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
    current := initial

    for _, stage := range stages {
        current = d.convertNumber(current, stage)
    }

    return current
}

func (d *Day5Part1) convertNumber(number int, stage string) int {
    if mapping, ok := d.mappings[stage][number]; ok {
        return mapping
    }
    return number // If no mapping, the number remains the same
}

// Additional methods to handle mapping logic
