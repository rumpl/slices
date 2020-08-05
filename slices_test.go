package slices

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestContains(t *testing.T) {
	assert.Equal(t, Contains([]string{"a", "b", "c"}, "a"), true)
	assert.Equal(t, Contains([]string{"a", "b", "c"}, "d"), false)
}
