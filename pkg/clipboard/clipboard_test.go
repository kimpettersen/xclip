package clipboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToClipboard(t *testing.T) {
	storage := InMemStorage{}
	c := New(&storage)
	item := NewItem("This is the first item")
	c.Add(&item)
	all, err := c.GetAll()
	assert.NoError(t, err)
	assert.Len(t, all, 1)
	assert.Equal(t, "This is the first item", all[0].Text)
	assert.NotEmpty(t, all[0].Added)
}
