# Added Day 3 Part 2

## Summary

This experience was very similar to part 1:

* It generated similar bugs to part 1 but was able to mostly generate its own logging and fixes through prompts
* I can probably speed up the process a bit by actually analyzing its solution sooner

## History

* Prompt:

        Please update the go file for part 2 of the advent of code problem that follows

        --- Part Two ---

        (pasted content of the problem)

* Prompt:

        I am getting a result of 0 from the sample input. What might be wrong?

* Prompt:

        It appears that the code is not capturing the start of numbers that are adjacent to "*" in later digits

* Prompt:

        It appears you've included support for "vertical numbers" but they are not part of the problem space

* Prompt:

        The number of return values from findStartOfHorizontalNumber is incorrect

* Prompt:

        Now a different problem with the solution.. it appears to only find the first adjacent number to each "*" instead of all adjacent numbers.

* Prompt:

        This update limits the matches to 2 adjacent numbers, but the problem was that I was only seeing 1 adjacent number for each, so that change has no impact.

* Prompt:

        Based upon the logging, it appears that it's failing to identify adjacent numbers in one of these two cases:

        * Diagonally adjacent
        * At a lower valued y direction

* Prompt:

        The direction of the search should not affect the direction of the search for numbers. Again, all numbers are horizontal. It appears the `// If checking up or down` clause is completely unecessary

* Prompt:

        Please update directions to include all 8 cardinal directions.

* Prompt:

        Please refer to this updated source code. No need to respond.

        (pasted current code)

* Prompt:

        Two issues:

        * Update the code to limit results to gears with exactly 2 adjacent numbers
        * Identify why the sum of products is always zero

* Prompt:

        Update the code to prevent duplicate numbers from being included in the sets of adjacent numbers
