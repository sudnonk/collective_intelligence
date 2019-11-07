package models

import "math"

type Radian float64

func (r Radian) sharpen() Radian {
	for {
		if r > 2*math.Pi {
			r -= 2 * math.Pi
		} else {
			return r
		}
	}
}

//負のラジアンを正にする
func (r Radian) AsPositive() Radian {
	if r < 0 {
		return r.sharpen() + 2*math.Pi
	} else {
		return r.sharpen()
	}
}

//もう一周した値を返す
func (r Radian) AddOneMore() Radian {
	return r.sharpen() + 2*math.Pi
}

//もう半周した値を返す
func (r Radian) AddHalf() Radian {
	return r.sharpen() + math.Pi
}

//float64にする
func (r Radian) AsFloat64() float64 {
	return float64(r.sharpen())
}

//ラジアンを弧度法にする（デバッグ用）
func (r Radian) AsDegree() float64 {
	return float64(r.sharpen() * 180 / math.Pi)
}

//弧度法からラジアンにする（デバッグ用）
func FromDegree(d float64) Radian {
	return Radian(d * math.Pi / 180).sharpen().AsPositive()
}
