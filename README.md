# Advent of Code 2023

## Usage

Build and execute unit tests (via rootful docker):

    ./build-test.sh

Build and launch webserver:

    ./build-run.sh

... and connect at: `http://127.0.0.1:8080`

## Experience with ChatGPT

This was all done in a programming language (Go Lang) that I would still say I have never used.

* [Initial Go Webserver Creation](docs/Initial-Go-Webserver-Creation.md)
  (~15 minutes with ChatGPT 4)
* [Refactored for Ease of Algorithm Addition](docs/Refactored-for-Ease-of-Algorithm-Addition.md)
  (~10 minutes with ChatGPT 4 and debugging copy/paste issues)
* [Added Day 1 Part 1](docs/Added-Day-1-Part-1.md) (~10 minutes with ChatGPT 4)
* [Added Day 1 Part 2](docs/Added-Day-1-Part-2.md) (~5 minutes with ChatGPT 4)
* [Added Day 2 Part 1](docs/Added-Day-2-Part-1.md) (~5 minutes with ChatGPT 4)
* [Added Day 2 Part 2](docs/Added-Day-2-Part-2.md) (~5 minutes with ChatGPT 4)
* [Updated CSS Styling](docs/Updated-CSS-Styling.md) (~5 minutes with ChatGPT 4)
* [Added Day 3 Part 1](docs/Added-Day-3-Part-1.md) (~1 hour with ChatGPT 4)
    * I let it generate a full solution then I helped debug the issues (mostly feeding it debug outputs)
* [Added Day 3 Part 2](docs/Added-Day-3-Part-2.md) (~1 hour with ChatGPT 4)
    * Similar experience to part 1. I was surprised to see it repeat errors from part 1.
* [Reimplmented Day 3 Part 1 (Day3V2Part1)](docs/Added-Day-3-Part-1.md) ("hours" with ChatGPT 4)
    * I wrote pseudo code for a solution and ChatGPT implemented it
    * I described where to refactor for testability and ChatGPT generated biolerplate unit tests (with often bad sample data)
    * In the end I had bugs in the underlying pseduo code that only appeared in the large dataset,
      which made debugging difficult
