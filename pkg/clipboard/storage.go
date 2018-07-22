package clipboard

type Storager interface {
	Save(string) error
	GetHistory(int) ([]string, error)
}

type InMemStorage struct {
	Items []string
}

func (im *InMemStorage) Save(i string) error {
	im.Items = append(im.Items, i)
	return nil
}

func (im *InMemStorage) GetHistory(size int) ([]string, error) {
	position := len(im.Items) - size
	if position < 0 {
		position = 0
	}
	res := im.Items[position:]
	return res, nil
}
