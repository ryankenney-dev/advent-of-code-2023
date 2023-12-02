# Refactored for Ease of Algorithm Addition

Here are my prompts:

* Refactor the interface between main.go and the other files so that the algorithm files so we call the same method on each algorithm to get:
    * The algorithm title
    * The algorithm computation result

* Update main.go to read the drop-down titles from the algorithm Title function

And now I have this nice interface implemented across the algorithms:

    // Algorithm defines the interface for different computations
    type Algorithm interface {
        Title() string
        Compute(input string) string
    }
