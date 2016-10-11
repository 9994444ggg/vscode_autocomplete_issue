package subone

import "github.com/sethgrid/vscode_autocomplete_issue/subtwo"

// SubOne is an example method that leverages a different subpackage to show VSCode's Go plugin failing to work as expected.
func SubOne(i int) int {
	// Open in vs code
	// Below here, start to type "subtwo" - you will not get autocomplete.
	// Also, you cannot cmd+p "> Go: Add Import" > "subtwo", it is not found
	return subtwo.SubTwo(i) + 1
}
