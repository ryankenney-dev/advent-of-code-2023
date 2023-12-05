# Added Day 4 Part 2

## Summary

Again I tried to let ChatGPT 4 solve the whole thing, but it didn't quite get there.

I then very generically instructed it to:

* Break the solution into parts
* Generate unit tests for each part

Then I fed it unit test failures and after a couple of iterations per test, I genearlly got a fix out of it.

I did hit one bug in a unit test dataset.

By the end of generting a complete solution, I have zero understanding of the problem it was
trying to solve--I never actually read the problem.

## History

* Prompt:

		Modify the solution for part 2 of the advent of code problem:

		--- Part Two ---

		(pasted the problem in)

* Prompt:

		Generate a unit test of Compute() for part 2

* Prompt:

		Why might this fail with:

		Expected Total Number of Scratchcards: 30, got Total Number of Scratchcards: 81

* Prompt:

		You have my exact implementation of Compute() above

* Prompt:

		Now I have this error:

		Test Example Case failed. Expected Total Number of Scratchcards: 30, got Total Number of Scratchcards: 63

* Prompt:

		Break the solution into multiple steps and generate unit tests for each step.

* Prompt:

		Please generate the new Compute body that uses these new functions

* Prompt:

		Also generate the body of the new methods you are calling.

* Prompt:

		Here is the first unit test failure:

		TestParseCards (0.00s)
			day4_part2_test.go:20: parseCards() = [map[17:{} 31:{} 41:{} 48:{} 53:{} 6:{} 83:{} 86:{} 9:{}] map[13:{} 16:{} 17:{} 19:{} 20:{} 24:{} 30:{} 32:{} 61:{} 68:{} 82:{}]], want [map[17:{} 41:{} 48:{} 83:{} 86:{}] map[13:{} 16:{} 20:{} 32:{} 61:{}]]

* Prompt:

		Here is the second failing unit test:

		TestCountMatches (0.00s)
			day4_part2_test.go:34: countMatches() = 9, want 5

* Prompt:

		Please provide an updated Compute() function that calls countMatches with the proper number of arguments

* Prompt:

		Please update this unit test to pass the correct data structures:

		func TestCountMatches(t *testing.T) {
			card := map[string]struct{}{
				"41": {}, "48": {}, "83": {}, "86": {}, "17": {},
				"6": {}, "31": {}, "9": {}, "53": {},
			}

			expected := 5 // Number of matches in the card
			result := countMatches(card)

			if result != expected {
				t.Errorf("countMatches() = %d, want %d", result, expected)
			}
		}

* Prompt:

		Please update this test method to use the new data structures:

		func TestCalculateTotalCards(t *testing.T) {
			cards := []map[string]struct{}{
				{"41": {}, "48": {}, "83": {}, "86": {}, "17": {}},
				{"13": {}, "32": {}, "20": {}, "16": {}, "61": {}},
				// ... Add all cards as per your test case
			}

			expected := 30 // Expected total cards
			total := calculateTotalCards(cards)

			if total != expected {
				t.Errorf("calculateTotalCards() = %d, want %d", total, expected)
			}
		}

* Prompt:

		Update this test function with the new data structures:

		func TestParseCards(t *testing.T) {
			input := `41 48 83 86 17 | 83 86  6 31 17  9 48 53
					13 32 20 16 61 | 61 30 68 82 17 32 24 19`

			expected := []map[string]struct{}{
				{"41": {}, "48": {}, "83": {}, "86": {}, "17": {}},
				{"13": {}, "32": {}, "20": {}, "16": {}, "61": {}},
			}

			cards := parseCards(input)

			if !reflect.DeepEqual(cards, expected) {
				t.Errorf("parseCards() = %v, want %v", cards, expected)
			}
		}

* Prompt:

		Now I see this error:

		TestCountMatches (0.00s)
			day4_part2_test.go:61: countMatches() = 4, want 5

* Prompt:

		Why is "expected := 5 // Number of matches in the card" ? Are there 5 numbers that match winners

* Prompt:

		I fixed that test by setting the expectation to 4.

		This is my final unit test failure:

		TestCalculateTotalCards (0.00s)
			day4_part2_test.go:100: calculateTotalCards() = 3, want 30

* Prompt:

		Please flush out then entire `cards` structure in this test per the original Part 2 problem sample data:

		func TestCalculateTotalCards(t *testing.T) {
			cards := []struct{
				winningNumbers map[string]struct{}
				playerNumbers  map[string]struct{}
			}{
				{
					winningNumbers: map[string]struct{}{"41": {}, "48": {}, "83": {}, "86": {}, "17": {}},
					playerNumbers:  map[string]struct{}{"83": {}, "86": {}, "6": {}, "31": {}, "17": {}, "9": {}, "48": {}, "53": {}},
				},
				{
					winningNumbers: map[string]struct{}{"13": {}, "32": {}, "20": {}, "16": {}, "61": {}},
					playerNumbers:  map[string]struct{}{"61": {}, "30": {}, "68": {}, "82": {}, "17": {}, "32": {}, "24": {}, "19": {}},
				},
				// ... Add all cards as per your test case
			}

			expected := 30 // Expected total cards
			total := calculateTotalCards(cards)

			if total != expected {
				t.Errorf("calculateTotalCards() = %d, want %d", total, expected)
			}
		}
