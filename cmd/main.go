package main

import (
	"fmt"

	"github.com/kimpettersen/clipboard/pkg/clipboard"
)

func main() {
	storage := clipboard.InMemStorage{}
	board := clipboard.New(&storage)
	fmt.Printf("Created board %s", board)
	// for {
	// Listen for events
	// If something is added to PRIMARY or CLIPBOARD, Add to clipboard
	//
	// board.Add(clipboard.NewItem(event.Text))
	//
	//

	// If it is key events, look for the configured combination that triggers listing of content
	//
	// board.GetAll()
	//
	// Continue to listen for key release. Figure out how to paste the selected element
	// }
}
