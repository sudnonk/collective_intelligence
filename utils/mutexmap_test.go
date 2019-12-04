package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangesOnlyAfterMerge(t *testing.T) {
	mm := NewMutexMap()
	mm.Set("po", "hoge")
	assert.Nil(t, mm.Get("po"))

	mm.Merge()
	assert.Equal(t, "hoge", mm.Get("po"))
}
