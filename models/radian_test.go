package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsPositive(t *testing.T) {
	//-45°
	r := FromDegree(-45)
	assert.Equal(t, float64(315), r.AsPositive().AsDegree())
}
