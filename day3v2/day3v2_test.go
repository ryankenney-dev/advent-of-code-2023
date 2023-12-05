package day3v2

import (
    "reflect"
    "testing"
)

func TestFindIntegersAsPotentialPartNumbers(t *testing.T) {
    // Setup
    // ----
	grid := [][]rune{
		{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
		{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
		{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
		{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
		{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
	}

    // Execute
    // ----
    result := findIntegersAsPotentialPartNumbers(grid)

    // Verify
    // ----
    pn467 := &PartNumber{StartCoordinate: Coordinate{X: 0, Y: 0}, Value: 467}
    pn114 := &PartNumber{StartCoordinate: Coordinate{X: 5, Y: 0}, Value: 114}
    pn35 := &PartNumber{StartCoordinate: Coordinate{X: 2, Y: 2}, Value: 35}
    pn633 := &PartNumber{StartCoordinate: Coordinate{X: 6, Y: 2}, Value: 633}
    pn617 := &PartNumber{StartCoordinate: Coordinate{X: 0, Y: 4}, Value: 617}
    pn58 := &PartNumber{StartCoordinate: Coordinate{X: 7, Y: 5}, Value: 58}
    pn592 := &PartNumber{StartCoordinate: Coordinate{X: 2, Y: 6}, Value: 592}
    pn755 := &PartNumber{StartCoordinate: Coordinate{X: 6, Y: 7}, Value: 755}
    pn664 := &PartNumber{StartCoordinate: Coordinate{X: 1, Y: 9}, Value: 664}
    pn598 := &PartNumber{StartCoordinate: Coordinate{X: 5, Y: 9}, Value: 598}
    expected := map[Coordinate]*PartNumber{
        {X: 0, Y: 0}: pn467, {X: 1, Y: 0}: pn467, {X: 2, Y: 0}: pn467,
        {X: 5, Y: 0}: pn114, {X: 6, Y: 0}: pn114, {X: 7, Y: 0}: pn114,
        {X: 2, Y: 2}: pn35,  {X: 3, Y: 2}: pn35,  {X: 6, Y: 2}: pn633,
        {X: 7, Y: 2}: pn633, {X: 8, Y: 2}: pn633, {X: 0, Y: 4}: pn617,
        {X: 1, Y: 4}: pn617, {X: 2, Y: 4}: pn617, {X: 7, Y: 5}: pn58,
        {X: 8, Y: 5}: pn58,  {X: 2, Y: 6}: pn592, {X: 3, Y: 6}: pn592,
        {X: 4, Y: 6}: pn592, {X: 6, Y: 7}: pn755, {X: 7, Y: 7}: pn755,
        {X: 8, Y: 7}: pn755, {X: 1, Y: 9}: pn664, {X: 2, Y: 9}: pn664,
        {X: 3, Y: 9}: pn664, {X: 5, Y: 9}: pn598, {X: 6, Y: 9}: pn598,
        {X: 7, Y: 9}: pn598,
    }
    if len(result) != len(expected) {
        t.Errorf("Length mismatch: got %d, want %d", len(result), len(expected))
    }
    for coord, partNum := range result {
        expectedPartNum, ok := expected[coord]
        if !ok {
            t.Errorf("Unexpected coordinate in result: %v", coord)
            continue
        }
        if *partNum != *expectedPartNum {
            t.Errorf("Mismatch at %v: got %+v, want %+v", coord, *partNum, *expectedPartNum)
        }
    }
}

// TestFindIntegersAsPotentialPartNumbers_WithSingleLength tests the
// findIntegersAsPotentialPartNumbers function, similarly to 
// TestFindIntegersAsPotentialPartNumbers_WithSingleLength, but it includes the special case
// of single-digit numbers, which did not apear in the original test
// dataset.
func TestFindIntegersAsPotentialPartNumbers_WithSingleLength(t *testing.T) {
    // Setup
    // ----
	grid := [][]rune{
		{'4', '6', '.', '4'},
		{'+', '.', '.', '5'},
		{'.', '2', '.', '*'},
	}

    // Execute
    // ----
    result := findIntegersAsPotentialPartNumbers(grid)

    // Verify
    // ----
    pn46 := &PartNumber{StartCoordinate: Coordinate{X: 0, Y: 0}, Value: 46}
    pn4 := &PartNumber{StartCoordinate: Coordinate{X: 3, Y: 0}, Value: 4}
    pn5 := &PartNumber{StartCoordinate: Coordinate{X: 3, Y: 1}, Value: 5}
    pn2 := &PartNumber{StartCoordinate: Coordinate{X: 1, Y: 2}, Value: 2}
    expected := map[Coordinate]*PartNumber{
        {X: 0, Y: 0}: pn46, {X: 1, Y: 0}: pn46,
        {X: 3, Y: 0}: pn4,
        {X: 3, Y: 1}: pn5,
        {X: 1, Y: 2}: pn2,
    }
    if len(result) != len(expected) {
        t.Errorf("Length mismatch: got %d, want %d", len(result), len(expected))
    }
    for coord, partNum := range result {
        expectedPartNum, ok := expected[coord]
        if !ok {
            t.Errorf("Unexpected coordinate in result: %v", coord)
            continue
        }
        if *partNum != *expectedPartNum {
            t.Errorf("Mismatch at %v: got %+v, want %+v", coord, *partNum, *expectedPartNum)
        }
    }
}

func TestCollectSymbolCoordinates(t *testing.T) {
    // Setup
    // ----
    grid := [][]rune{
		{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
		{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
		{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
		{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
		{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
    }

    // Execute
    // ----
    result := collectSymbolCoordinates(grid)

    // Verify
    // ----
    expected := map[Coordinate]struct{}{
        {X: 3, Y: 1}: {},
        {X: 7, Y: 3}: {},
        {X: 3, Y: 4}: {},
        {X: 5, Y: 5}: {},
        {X: 3, Y: 8}: {},
        {X: 5, Y: 8}: {},
    }
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("collectSymbolCoordinates() = %v, want %v", result, expected)
    }
}

func TestGetNumbersAdjacentToSymbols(t *testing.T) {
    // Setup
    // ----
    symbolCoordinates := map[Coordinate]struct{}{
        {X: 3, Y: 1}: {},
        {X: 7, Y: 3}: {},
        {X: 3, Y: 4}: {},
        {X: 5, Y: 5}: {},
        {X: 3, Y: 8}: {},
        {X: 5, Y: 8}: {},
    }
    pn467 := &PartNumber{StartCoordinate: Coordinate{X: 0, Y: 0}, Value: 467}
    pn114 := &PartNumber{StartCoordinate: Coordinate{X: 5, Y: 0}, Value: 114}
    pn35 := &PartNumber{StartCoordinate: Coordinate{X: 2, Y: 2}, Value: 35}
    pn633 := &PartNumber{StartCoordinate: Coordinate{X: 6, Y: 2}, Value: 633}
    pn617 := &PartNumber{StartCoordinate: Coordinate{X: 0, Y: 4}, Value: 617}
    pn58 := &PartNumber{StartCoordinate: Coordinate{X: 7, Y: 5}, Value: 58}
    pn592 := &PartNumber{StartCoordinate: Coordinate{X: 2, Y: 6}, Value: 592}
    pn755 := &PartNumber{StartCoordinate: Coordinate{X: 6, Y: 7}, Value: 755}
    pn664 := &PartNumber{StartCoordinate: Coordinate{X: 1, Y: 9}, Value: 664}
    pn598 := &PartNumber{StartCoordinate: Coordinate{X: 5, Y: 9}, Value: 598}
    numbersIndex := map[Coordinate]*PartNumber{
        {X: 0, Y: 0}: pn467, {X: 1, Y: 0}: pn467, {X: 2, Y: 0}: pn467,
        {X: 5, Y: 0}: pn114, {X: 6, Y: 0}: pn114, {X: 7, Y: 0}: pn114,
        {X: 2, Y: 2}: pn35,  {X: 3, Y: 2}: pn35,  {X: 6, Y: 2}: pn633,
        {X: 7, Y: 2}: pn633, {X: 8, Y: 2}: pn633, {X: 0, Y: 4}: pn617,
        {X: 1, Y: 4}: pn617, {X: 2, Y: 4}: pn617, {X: 7, Y: 5}: pn58,
        {X: 8, Y: 5}: pn58,  {X: 2, Y: 6}: pn592, {X: 3, Y: 6}: pn592,
        {X: 4, Y: 6}: pn592, {X: 6, Y: 7}: pn755, {X: 7, Y: 7}: pn755,
        {X: 8, Y: 7}: pn755, {X: 1, Y: 9}: pn664, {X: 2, Y: 9}: pn664,
        {X: 3, Y: 9}: pn664, {X: 5, Y: 9}: pn598, {X: 6, Y: 9}: pn598,
        {X: 7, Y: 9}: pn598,
    }

    // Execute
    // ----
    result := getNumbersAdjacentToSymbols(symbolCoordinates, numbersIndex)

    // Verify
    // ----
    expected := []*PartNumber{
        pn467, pn35, pn633, pn617, pn592, pn755, pn664, pn598,
    }
    if len(result) != len(expected) {
        t.Errorf("Expected and actual results differ in length: expected %d, got %d", len(expected), len(result))
        return
    }
    for i, expectedPart := range expected {
        if result[i] != expectedPart {
            t.Errorf("Mismatch at index %d: expected %+v, got %+v", i, expectedPart, result[i])
        }
    }
}

// TestGetNumbersAdjacentToSymbols_DuplicateNumbers test the
// getNumbersAdjacentToSymbols_DuplicateNumbers function in a similar way to
// TestGetNumbersAdjacentToSymbols, but it includes the special case
// of duplicate numbers in the grid, which did not apear in the original test
// dataset.
func TestGetNumbersAdjacentToSymbols_DuplicateNumbers(t *testing.T) {
    // Setup
    // ----
    symbolCoordinates := map[Coordinate]struct{}{
        {X: 3, Y: 1}: {},
        {X: 7, Y: 1}: {},
    }
    pn467 := &PartNumber{StartCoordinate: Coordinate{X: 0, Y: 0}, Value: 467}
    pn467b := &PartNumber{StartCoordinate: Coordinate{X: 5, Y: 0}, Value: 467}
    numbersIndex := map[Coordinate]*PartNumber{
        {X: 0, Y: 0}: pn467, {X: 1, Y: 0}: pn467, {X: 2, Y: 0}: pn467,
        {X: 5, Y: 0}: pn467b, {X: 6, Y: 0}: pn467b, {X: 7, Y: 0}: pn467b,
    }

    // Execute
    // ----
    result := getNumbersAdjacentToSymbols(symbolCoordinates, numbersIndex)

    // Verify
    // ----
    expected := []*PartNumber{
        pn467, pn467b,
    }
    if len(result) != len(expected) {
        t.Errorf("Expected and actual results differ in length: expected %d, got %d", len(expected), len(result))
        return
    }
    for i, expectedPart := range expected {
        if result[i] != expectedPart {
            t.Errorf("Mismatch at index %d: expected %+v, got %+v", i, expectedPart, result[i])
        }
    }
}

func TestExecute(t *testing.T) {
    // Setup
    // ----
	grid := [][]rune{
		{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
		{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
		{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
		{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
		{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
    }

    // Execute
    // ----
	result := execute(grid)

    // Verify
    // ----
    expectedSum := 4361
	if result != expectedSum {
        t.Errorf("execute() = %d, want %d", result, expectedSum)
    }
}