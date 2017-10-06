package metrics

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimer_Update(t *testing.T) {
	c := Timer{}
	c.Update(100 * time.Millisecond)
	assert.Equal(t, 100*time.Millisecond, c.Value(), "unexpected value")
}
