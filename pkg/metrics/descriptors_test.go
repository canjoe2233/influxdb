package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID_newID(t *testing.T) {
	var id = newID(0xff, 0xff0f0fff)
	assert.Equal(t, ID(0xff0f0fff000000ff), id)
	assert.Equal(t, 0xff, id.id())
	assert.Equal(t, 0xff0f0fff, id.gid())
}

func TestID_setGID(t *testing.T) {
	var id = ID(1)
	assert.Equal(t, 0, id.gid())
	id.setGID(1)
	assert.Equal(t, 1, id.gid())
}
