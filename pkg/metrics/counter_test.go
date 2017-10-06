package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter_Add(t *testing.T) {
	c := Counter{}
	c.Add(5)
	c.Add(5)
	assert.Equal(t, int64(10), c.Value(), "unexpected value")
}
