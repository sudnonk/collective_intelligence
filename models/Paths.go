package models

import (
	"github.com/sudnonk/collective_intelligence/config"
	"math"
	"math/rand"
	"sort"
)

type Paths map[string]*Path

func (ps *Paths) len() int {
	s := *ps
	return len(s)
}

//Pathsの中で最も間隔があいてる二つのPathを返す
func (ps Paths) GetWidest() (p1 *Path, p2 *Path) {
	max := 0.0

	//ソートされたpsのkeys
	sks := SortPathsByArg(ps)
	l := len(ps)
	i1 := 0
	i2 := l - 1

	for i := range sks {
		if i == 0 {
			m := math.Abs(ps[sks[i]].Arg.sharpen().AddOneMore().AsFloat64() - ps[sks[l-1]].Arg.sharpen().AsFloat64())
			if m > max {
				max = m
				i1 = i
				i2 = l - 1
			}
		} else {
			m := math.Abs(ps[sks[i-1]].Arg.sharpen().AsFloat64() - ps[sks[i]].Arg.sharpen().AsFloat64())

			if m > max {
				max = m
				i1 = i - 1
				i2 = i
			}
		}
	}

	//i2の方が大きいはず=i2の方が角度が大きいはず
	return ps[sks[i2]], ps[sks[i1]]
}

//PathsをArgが小さい順に並び替える
func SortPathsByArg(ps Paths) []string {
	ks := ps.keys()
	sort.Slice(ks, func(i, j int) bool {
		return ps[ks[i]].Arg.sharpen() < ps[ks[j]].Arg.sharpen()
	})
	return ks
}

func decideNewPath(c Cell) (arg Radian, d float64) {
	//もしすでに複数本道があれば
	if c.getPaths().len() > 1 {
		//一番間隔が開いてる2つのPath
		p1, p2 := c.getPaths().GetWidest()
		//そのPathの角度の平均
		arg = (p1.Arg.sharpen() + p2.Arg.sharpen()) / 2
		//そのPathの距離の平均、小さすぎれば補正
		d = math.Max((p1.Length+p2.Length)/2, config.MinDist())
	} else if c.getPaths().len() == 1 {
		//1本なら既存のやつの逆
		k0 := c.getPaths().keys()[0]
		arg = (*c.getPaths())[k0].Arg.sharpen().AddHalf()
		d = (*c.getPaths())[k0].Length
	} else {
		//0本ならランダム
		arg, d = randomNewPath()
	}

	return arg.sharpen(), d
}

func randomNewPath() (arg Radian, d float64) {
	arg = FromDegree(float64(rand.Intn(360))).sharpen()
	d = float64(rand.Intn(10)) + config.MinDist()

	return arg, d
}

func (ps Paths) keys() []string {
	var ks []string
	for k := range ps {
		ks = append(ks, k)
	}

	return ks
}
