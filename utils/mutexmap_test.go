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

func TestEqualsAfterMerge(t *testing.T) {
	mm := NewMutexMap()
	mm.Set("po", "hoge")
	mm.Merge()

	assert.Equal(t, mm.dirty["po"], mm.readonly["po"])
}
