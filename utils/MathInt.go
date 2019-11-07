package utils

import "math"

func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func Round(n float64) int64 {
	return int64(math.Round(n))
}

func Min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min64(a int64, b int64) int64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max64(a int64, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}
