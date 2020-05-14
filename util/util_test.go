package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage(t *testing.T) {

	var sl = []string{"1", "2", "3"}

	assert.True(t, InSlice(sl, "2"), "Item found in the slice")
	assert.False(t, InSlice(sl, "5"), "Item not found in the slice")
}