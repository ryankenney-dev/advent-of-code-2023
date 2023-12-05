# Reimplmented Day 3 Part 1

## Summary

I wanted to try solving the same problem with a different approach.
This time around I wrote pseudo code and instructed ChatGPT 4
to implement it.
I then instructed ChatGPT 4 to break out unit testable code for each
piece of my pseudo code so that I could verify it.

This worked fairly well at first, but it got very combersome trying
to get ChatGPT to make incremental fixes without introducing new bugs.
It also failed at generation of test data sets pretty badly.

Still, this was a decent way to work, and although the result required
much more effort on my part (vs my first Day3Part1 approach),
it includes thorough unit testing and I have a deep understading of the algorithm.

Unfortunately I ran into bugs in my underlying algorithm
that were not discoverable with the sample data set--only the real/huge dataset.
At that point I had to resort to using `grep`/`awk` commands to diagnose the
special cases that the data introduced.

One of the biggest limiting factors with ChatGPT during this was how long it
takes to generate the fairly-long responses.

## History

* Prompt:

        In GoLang, define a function "Compute" that accepts a 2-dimensional array of characters and implements the following pseudo code.

        * Define a coordinate object type, with x and y member values

        * Define variable "numbersIndex" as a map of coordinates to integers

        * Define variable "partNumbers" as a set (non-duplicate values) of integers

        * Find all, horizontal, multi-digit integers in the grid

            * For each multi-digit integer in the grid, add an entry to the "numbersIndex" map with a key of the coordinate and a value of the number

        * Find all non-numerical, non-period symbols in the grid

            * For each symbol, for each of the 8 cardinal directions, query the "numbersIndex" of the adjacent coordinate
                * If the coordinate is in the "numbersIndex", add the value of the "numbersIndex" entry to "partNumbers"

        * Return the sum of "partNumbers" entries

* Prompt:

        Rename `Compute()` to `execute()`

* Prompt:

        Extract the `// Find all horizontal multi-digit integers` block into a function and then write a unit test for it.

* Prompt:

        Why did you name the grid "rune"

* Prompt:

        Convert this grid of characters:

        467..114..
        ...*......
        ..35..633.
        ......#...
        617*......
        .....+.58.
        ..592.....
        ......755.
        ...$.*....
        .664.598..

        ... to an inline array in the format of:

            grid := [][]rune{
                {'.', '1', '2', '3', '#'},
                {'4', '5', '6', '*', '.'},
                {'.', '7', '8', '9', '0'},
            }

* Prompt:

        Please populate the list of expected values for that grid in this format:

            expected := map[Coordinate]int{
                {X: 1, Y: 0}: 123,
                {X: 0, Y: 1}: 456,
                {X: 1, Y: 2}: 7890,
            }

* Prompt:

        Modify `execute()` to collect all symbol coordinates in a set of Coordinate objects before and then iterating over that set

* Prompt:

        Extract the "// Collect symbol coordinates" block into a function and create a unit test

* Prompt:

        Convert the `// Check adjacent cells of symbols for numbers` block into a function and write a unit test

* Prompt:

        Generate the code for this block:

                expected := map[int]struct{}{
                    // ... (expected numbers adjacent to symbols) ...
                }

        Given this example sample data:

        467..114..
        ...*......
        ..35..633.
        ......#...
        617*......
        .....+.58.
        ..592.....
        ......755.
        ...$.*....
        .664.598..

* Prompt:

        Rewrite that assuming each symbol can have more than one adjacent number, and in all 8 cardinal directions

* Prompt:

        Can we print expected/actual results in an pretty format (indented)?

* Prompt:

        Update the following method to collect the coordinate of every character of every number. For example, a 3 digit number would have three coordinates in the resulting map.

        func findHorizontalMultiDigitIntegers(grid [][]rune) map[Coordinate]int {
            numbersIndex := make(map[Coordinate]int)
            for y, row := range grid {
                number := ""
                for x, char := range row {
                    if unicode.IsDigit(char) {
                        number += string(char)
                    } else if number != "" {
                        if len(number) > 1 {
                            num, _ := strconv.Atoi(number)
                            numbersIndex[Coordinate{X: x - len(number), Y: y}] = num
                        }
                        number = ""
                    }
                }
                if number != "" && len(number) > 1 {
                    num, _ := strconv.Atoi(number)
                    numbersIndex[Coordinate{X: len(row) - len(number), Y: y}] = num
                }
            }
            return numbersIndex
        }

* Prompt:

        Generate a unit test for the execute function

* Prompt:

        Given a numbersIndex map, such as the following:

            numbersIndex := map[Coordinate]int{
                {X: 0, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 1, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 2, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 5, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 6, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 7, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 2, Y: 2}: 35,    // "35" starting at (2,2)
                {X: 3, Y: 2}: 35,    // "35" starting at (2,2)
                {X: 6, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 7, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 8, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 0, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 1, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 2, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 7, Y: 5}: 58,    // "58" starting at (6,5)
                {X: 8, Y: 5}: 58,    // "58" starting at (6,5)
                {X: 2, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 3, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 4, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 6, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 7, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 8, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 1, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 2, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 3, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 5, Y: 9}: 598,   // "598" starting at (4,9)
                {X: 6, Y: 9}: 598,   // "598" starting at (4,9)
                {X: 7, Y: 9}: 598,   // "598" starting at (4,9)
            }

        Write code to generate a text block of the coordinates, with every coordinate from the map marked with an "X". Here is an example output for the above data set:

        XXX..XXX..
        ..........
        ..XX.XXX.
        ..........
        XXX.......
        .......XX.
        ..XXX.....
        ......XXX.
        ..........
        .XXX.XXX..

* Prompt:

        Make that function return a string

* Prompt:

        findHorizontalMultiDigitIntegers is not finding numbers shorter than 2 digits

* Prompt:

        Update the logic to store partNumbers in a list, not a unique set.

* Prompt:

        Update the following file so that:

        * `findHorizontalMultiDigitIntegers()` returns a map of coordinates to integer objects, which should allow us to identify unique instances of the same integer when they are in different locations in the overall grid
        * `checkAdjacentCellsForNumbers()` returns a list of numbers (not a set of unique numbers), where duplicate numbers are identified as unique integer objects referenced by the `numbersIndex` values

        Here is the current file:

        (pasted file contents)

* Prompt:

        That wasn't quite right. Let's try this again...

        Create a new object type `PartNumber` that contains two fields:

        * `StartCoordinate`: The coordinate of the first character of a number found in the grid
        * `Value`: The integer value of the PartNumber

        Update the following file so that:

        * `findHorizontalMultiDigitIntegers()` returns a map of coordinates to PartNumber objects, where each multi-character integer in the input grid will generate multiple key entries in the return, but exactly one value object (of type PartNumber)
        * `checkAdjacentCellsForNumbers()` returns a list of PartNumber objects, which may include duplicate part number values, but exactly one PartNumber object per integer in the original grid

        Modify the  file I provided in my previous request.

* Prompt:

        Update this unit test to use the new function interface

        func TestCheckAdjacentCellsForNumbers(t *testing.T) {
            symbolCoordinates := map[Coordinate]struct{}{
                {X: 3, Y: 1}: {},
                {X: 7, Y: 3}: {},
                {X: 3, Y: 4}: {},
                {X: 5, Y: 5}: {},
                {X: 3, Y: 8}: {},
                {X: 5, Y: 8}: {},
            }

            numbersIndex := map[Coordinate]int{
                {X: 0, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 1, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 2, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 5, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 6, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 7, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 2, Y: 2}: 35,    // "35" starting at (2,2)
                {X: 3, Y: 2}: 35,    // "35" starting at (2,2)
                {X: 6, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 7, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 8, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 0, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 1, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 2, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 7, Y: 5}: 58,    // "58" starting at (6,5)
                {X: 8, Y: 5}: 58,    // "58" starting at (6,5)
                {X: 2, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 3, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 4, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 6, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 7, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 8, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 1, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 2, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 3, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 5, Y: 9}: 598,   // "598" starting at (4,9)
                {X: 6, Y: 9}: 598,   // "598" starting at (4,9)
                {X: 7, Y: 9}: 598,   // "598" starting at (4,9)
            }

            expected := []int{
                467,
                35,
                633,
                617,
                592,
                755,
                664,
                598,
            }

            result := checkAdjacentCellsForNumbers(symbolCoordinates, numbersIndex)

            if !reflect.DeepEqual(result, expected) {
                t.Errorf("EXPECTED\n%s\nACTUAL\n%s\n", prettyFormat(expected), prettyFormat(result))
            }
        }

* Prompt:

        Please flush out the full unit test function

* Prompt:

        Update this unit test for the new function signatures. Please include the complete fuction (no elipses)

        func TestFindHorizontalMultiDigitIntegers(t *testing.T) {
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
            
            expected := map[Coordinate]int{
                {X: 0, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 1, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 2, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 5, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 6, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 7, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 2, Y: 2}: 35,    // "35" starting at (2,2)
                {X: 3, Y: 2}: 35,    // "35" starting at (2,2)
                {X: 6, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 7, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 8, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 0, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 1, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 2, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 7, Y: 5}: 58,    // "58" starting at (6,5)
                {X: 8, Y: 5}: 58,    // "58" starting at (6,5)
                {X: 2, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 3, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 4, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 6, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 7, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 8, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 1, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 2, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 3, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 5, Y: 9}: 598,   // "598" starting at (4,9)
                {X: 6, Y: 9}: 598,   // "598" starting at (4,9)
                {X: 7, Y: 9}: 598,   // "598" starting at (4,9)
            }

            result := findHorizontalMultiDigitIntegers(grid)

            if !reflect.DeepEqual(result, expected) {
                t.Errorf("findHorizontalMultiDigitIntegers() = %v, want %v", result, expected)
            }
        }

* Prompt:

        Please complete the entire  test function... including all parts commented with "... repeat for other"

* Prompt:

        Why is the unit test comparison showing an odd encoding for the PartNumber objects

        day3_test.go:66: findHorizontalMultiDigitIntegers() = map[{0 0}:0xc00017e078 {0 4}:0xc00017e0d8 {1 0}:0xc00017e078 ...], want map[{0 0}:467 {0 4}:617 {1 0}:467  ...]

* Prompt:

        I'm seeing this error now:

        invalid operation: cannot indirect expectedPartNum (variable of type int)

        I assume my `expected` definition needs an update:

            expected := map[Coordinate]int{
                {X: 0, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 1, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 2, Y: 0}: 467,   // "467" starting at (0,0)
            ...

* Prompt:

        Can you rewrite this entire block in that new format? 

            expected := map[Coordinate]int{
                {X: 0, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 1, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 2, Y: 0}: 467,   // "467" starting at (0,0)
                {X: 5, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 6, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 7, Y: 0}: 114,   // "114" starting at (5,0)
                {X: 2, Y: 2}: 35,    // "35" starting at (2,2)
                {X: 3, Y: 2}: 35,    // "35" starting at (2,2)
                {X: 6, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 7, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 8, Y: 2}: 633,   // "633" starting at (5,2)
                {X: 0, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 1, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 2, Y: 4}: 617,   // "617" starting at (0,4)
                {X: 7, Y: 5}: 58,    // "58" starting at (6,5)
                {X: 8, Y: 5}: 58,    // "58" starting at (6,5)
                {X: 2, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 3, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 4, Y: 6}: 592,   // "592" starting at (2,6)
                {X: 6, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 7, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 8, Y: 7}: 755,   // "755" starting at (6,7)
                {X: 1, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 2, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 3, Y: 9}: 664,   // "664" starting at (1,9)
                {X: 5, Y: 9}: 598,   // "598" starting at (4,9)
                {X: 6, Y: 9}: 598,   // "598" starting at (4,9)
                {X: 7, Y: 9}: 598,   // "598" starting at (4,9)
            }

* Prompt:

        Can we make checkAdjacentCellsForNumbers return a set of PartNumbers, so each PartNumber is included once, but if the part number appears in the grid more than once, then it appears twice 

* Prompt:

        I do not want to include a partnumber object for each part number character in the return. I want to include exactly one partnumber object per part number with a unique starting location within the gride

* Prompt:

        Can you sort the result by coordinates (x, then y)?

* Prompt:

        Reverse x and y as to which sorts first

* Prompt:

        Generate an example comment for the unit test function
