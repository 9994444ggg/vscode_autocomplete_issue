package main

import (
	"fmt"

	"github.com/sethgrid/vscode_autocomplete_issue/subone"
)

func main() {
	// Open in vs code
	// Below here, start to type "subone" - you will not get autocomplete.
	// Also, you cannot cmd+p "> Go: Add Import" > "subone", it is not found
	fmt.Println(subone.SubOne(5))
}
