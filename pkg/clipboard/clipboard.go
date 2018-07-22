package clipboard

import (
	"time"
)

type Item struct {
	Text  string
	Added time.Time
}

type Clipboard struct {
	Items   []Item
	Storage Storager
}

func NewItem(text string) Item {
	return Item{
		Text:  text,
		Added: time.Now(),
	}
}

func New(storage Storager) Clipboard {
	c := Clipboard{
		Storage: storage,
	}
	return c
}

func (c Clipboard) Add(i *Item) {
	c.Storage.Save(i)
}

func (c Clipboard) GetAll() ([]*Item, error) {
	items, err := c.Storage.GetAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}
