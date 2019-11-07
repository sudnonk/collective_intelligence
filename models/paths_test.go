package models

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSort(t *testing.T) {
	ps := Paths{
		"testCase_1": &Path{
			Arg: FromDegree(45),
		},
		"testCase_2": &Path{
			Arg: FromDegree(-45),
		},
	}

	sks := SortPathsByArg(ps)
	assert.Equal(t, float64(45), ps[sks[0]].Arg.AsDegree())
	assert.Equal(t, float64(315), ps[sks[0]].Arg.AsDegree())
}

func TestGetWidest(t *testing.T) {
	ps := Paths{
		"testCase_1": &Path{
			Arg: FromDegree(45),
		},
		"testCase_2": &Path{
			Arg: FromDegree(315),
		},
		"testCase_3": &Path{
			Arg: FromDegree(170),
		},
	}

	p1, p2 := ps.GetWidest()

	assert.Equal(t, float64(145), math.Abs(p1.Arg.AsDegree()-p2.Arg.AsDegree()))
}

func TestGetWidest2(t *testing.T) {
	ps := Paths{
		"testCase_1": &Path{
			Arg: FromDegree(45),
		},
		"testCase_2": &Path{
			Arg: FromDegree(180),
		},
		"testCase_3": &Path{
			Arg: FromDegree(315),
		},
	}

	p1, p2 := ps.GetWidest()

	assert.Equal(t, float64(135), math.Abs(p1.Arg.AsDegree()-p2.Arg.AsDegree()))
}
