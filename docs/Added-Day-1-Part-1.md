# Added Part 1 Day 1

I tried to see if ChatGPT 4 would just lookup the answer with this prompt:

    Solve advent of code 2023 question 1 part 1 using go

... but it doesn't appear to do that (good!). It responded with:

> To provide a solution for Advent of Code 2023's Day 1, Part 1 using Go, I would need the specific problem statement. Advent of Code presents a variety of coding challenges, and each year the problems are unique. Could you please provide the details of the question you'd like me to solve?

So here is how I go to an answer...

* Prompt:

        (entire body of Advent of Code Question)

* Prompt:

        Remove the main function so it does not conflict with other code.

* Prompt:

        Refactor this file to implement this interface:

        // Algorithm defines the interface for different computations
        type Algorithm interface {
        Title() string
        Compute(input string) string
        }

        ... where:

        * Title describes the meaning of the resulting value
        * Compute executes the main algorithm and returns a string in the format "<Title>: <Resulting Number>"

* Prompt:

        Keep the printable title "Sum of Calibration Values", but name the class Day1Part1
