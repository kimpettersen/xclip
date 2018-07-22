package clipboard

type Storager interface {
	Save(i *Item) error
	GetAll() ([]*Item, error)
}

type InMemStorage struct {
	Items []*Item
}

func (im *InMemStorage) Save(i *Item) error {
	im.Items = append(im.Items, i)
	return nil
}

func (im *InMemStorage) GetAll() ([]*Item, error) {
	return im.Items, nil
}
