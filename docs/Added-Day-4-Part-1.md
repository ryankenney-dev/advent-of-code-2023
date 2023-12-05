# Added Day 4 Part 1

I only submitted a couple of requests to ChatGPT 4 for this one and I had a complete solution:

* Prompt:

        Given this Go interface:

        type Algorithm interface {
            Title() string
            Compute(input string) string
        }

        ... where:

        * Title describes the meaning of the resulting numberic value
        * Compute executes the main algorithm and returns a string in the format "<Title>: <Resulting Number>"

        Solve the following Advent of Code problem as a Go library that implements the inteface. Name the resulting class "Day4Part1".

        (entire body of Advent of Code question)

* Prompt:

        Wite a unit test for Compute()
