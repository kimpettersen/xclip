package clipboard

import (
	"sort"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

// Clipboard is the struct that contains a history of copied/cut items
type Clipboard struct {
	Items       []string
	Storage     Storager
	HistorySize int
}

// New returns a new Clipboard with a storage engine
func New(storage Storager, historySize int) Clipboard {
	c := Clipboard{
		Storage:     storage,
		HistorySize: historySize,
	}
	return c
}

// Add adds a new item to the clipboard
func (c Clipboard) Add(s string) {
	c.Storage.Save(s)
}

// GetHistory returns a list of all items in the clipboard
func (c Clipboard) GetHistory() ([]string, error) {
	items, err := c.Storage.GetHistory(c.HistorySize)

	if err != nil {
		return nil, err
	}
	return items, nil
}

// Find is a case insensitive search of the clipboard
func (c Clipboard) Find(s string) ([]string, error) {
	history, err := c.GetHistory()
	if err != nil {
		return nil, err
	}
	findResult := fuzzy.RankFindFold(s, history)
	sort.Sort(findResult)
	var rankedResult []string
	for _, i := range findResult {
		rankedResult = append(rankedResult, i.Target)
	}
	return rankedResult, nil
}
