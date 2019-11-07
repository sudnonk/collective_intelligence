package utils

import "math/rand"

//minからmaxの間で一様分布にランダムな値を返す
func RandFloat64n(min float64, max float64) float64 {
	return (rand.Float64() * (max - min)) + min
}
