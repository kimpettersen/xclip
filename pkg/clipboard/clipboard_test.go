package clipboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToClipboard(t *testing.T) {
	storage := InMemStorage{}
	c := New(&storage, 10)
	c.Add("This is the first item")
	all, err := c.GetHistory()
	assert.NoError(t, err)
	assert.Len(t, all, 1)
	assert.Equal(t, "This is the first item", all[0])
}

func TestHistorySize(t *testing.T) {
	storage := InMemStorage{}
	c := New(&storage, 3)
	c.Add("sam")
	c.Add("sample")
	c.Add("samurai")
	c.Add("samosa")
	hist, _ := c.GetHistory()
	assert.Len(t, hist, 3)
	assert.Equal(t, "sample", hist[0])
	assert.Equal(t, "samurai", hist[1])
	assert.Equal(t, "samosa", hist[2])

	res, _ := c.Find("sam")
	assert.NotContains(t, res, "sam")
}

func TestSearch(t *testing.T) {
	storage := InMemStorage{}
	c := New(&storage, 10)

	c.Add("Cow")
	c.Add("Cowboy")
	c.Add("OWL")
	c.Add("Horse")

	res, err := c.Find("ow")
	assert.NoError(t, err)
	assert.Len(t, res, 3)
	assert.Equal(t, "Cow", res[0])
	assert.Equal(t, "OWL", res[1])
	assert.Equal(t, "Cowboy", res[2])
}
