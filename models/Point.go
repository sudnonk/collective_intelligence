package models

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/sudnonk/collective_intelligence/config"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//p1とp2の間の距離
func calcDistance(p1 *Point, p2 *Point) float64 {
	x := p1.X - p2.X
	y := p1.Y - p2.Y
	return math.Sqrt(float64(x*x + y*y))
}

//p1を基準としたp2の角度
func calcArg(p1 *Point, p2 *Point) Radian {
	x := p2.X - p1.X
	y := p2.Y - p1.Y

	return Radian(math.Atan2(float64(y), float64(x))).AsPositive()
}

//起点とそこからの距離・角度から終点を計算する
func calcNewPoint(a Radian, d float64, fp *Point) *Point {
	c := math.Cos(a.AsFloat64())
	s := math.Sin(a.AsFloat64())

	return &Point{
		X: int(math.Round(c*d + float64(fp.X))),
		Y: int(math.Round(s*d + float64(fp.Y))),
	}
}

//ランダムな点
func randomPoint() *Point {
	return &Point{
		X: rand.Intn(config.WorldSizeX()),
		Y: rand.Intn(config.WorldSizeY()),
	}
}

func isSame(p1 *Point, p2 *Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func (p Point) identify() string {
	ps := []int{p.X, p.Y}
	sort.Ints(ps)
	return fmt.Sprintf("%d,%d", ps[0], ps[1])
}
