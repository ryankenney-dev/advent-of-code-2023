# Added Day 3 Part 1

## Summary

For this problem I tried to use my brain as little as possible and just feed ChatGPT 4
outputs that indicated a problem. In the end:

* ChatGPT 4 was able to solve at least one bug with a simple restatement of the problem
* ChatGPT 4 was able to add useful debugging info (which allowed me to leave my brain off)
* I did have to diagnose a bug myself, but ChatGPT 4 was able to generate a fix from there

## History

I stated with my normal request:

* Prompt:

        Given this Go interface:

        type Algorithm interface {
            Title() string
            Compute(input string) string
        }

        ... where:

        * Title describes the meaning of the resulting numberic value
        * Compute executes the main algorithm and returns a string in the format "<Title>: <Resulting Number>"

        Solve the following Advent of Code problem as a Go library that implements the inteface. Name the resulting class "Day3Part1".

        === Begin Advent of Code Problem ===

        (pasted content of the problem)

Then I tried using 0% of my brain and just tossing the wrong answer back at it:

* Prompt:

        For the following input, I am getting a result of 64, but I should be getting 4361.

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

        Did the body of isAdacentToSymbol change since the signature did?

With the reponses at least one bug was fixed.

I tried the same approach again:

* Prompt:

        I am now getting a value of 2325 for the sample input. I'm not sure your solution contains an error or if I entered something incorrectly. Here is my current file:

        (pasted file)

... but this did not necessarily improve things.

Then I went on an info gathering cycle:

* Prompt:

        Please add logging to the file that indicates each number parsed and whether or not it is being included in the sum.

* Prompt:

        I notice that number 617 is being excluded from the sum but is adjacent to "*"

* Prompt:

        Please add logging to isAdjacentToSymbol that logs the value of every individual condition tested

* Prompt:

        Remove those changes and add a single log line before the if statement, that follows this pseudo code:

        log.Printf("Conditions: %s %s %s ...", newX >= 0, newX < len(grid), newY >= 0, ...)

* Prompt:

        I notice that 617 has no conditions checked.

* Prompt:

        This condition is failing for 467:
        
        "if _, found := visited[endIndex]; !found"


At this point I accidentally jumped into a new thread, but I got it back on track fairly quickly...

* Prompt:

        If I am reading the code correctly, the `visited` array should store found numbers by endindex but it appears the value being used is actually the endindex within a single column, so columns are colliding in that array.

* Prompt:

        Please update your code to properly deconflict rows in the visited array. Let's use this copy of your code to start:
        
        (pasted code)

At this point I had a working solution.
