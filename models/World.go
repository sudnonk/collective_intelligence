package models

import (
	crand "crypto/rand"
	"fmt"
	"github.com/sudnonk/collective_intelligence/config"
	"github.com/sudnonk/collective_intelligence/utils"
	"gonum.org/v1/gonum/mat"
	"log"
	"math"
	"math/big"
	"math/rand"
	"sync"
	"time"
)

//その時点で世界にある細胞
var Cells *Nodes

//その時点で世界にあるパス
var Roads *MutexPaths

//その時点で世界座標上のどこに細胞があるか
var Grid *mat.Dense

//その時点で爆撃された地点
var BombPoint *Point

//世界を実行する
func Run() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	//初期化
	Cells = NewCells()
	Roads = NewPaths()
	UpdateGrid()

	//一つ目の細胞を配置
	init := GenerateRandom()
	Cells.Set(init.Id, init)
	UpdateGrid()

	for step := int64(0); step < config.MaxStep(); step++ {
		s := time.Now().Nanosecond()
		UpdateGrid()
		decideBombArea(step)

		var wg sync.WaitGroup

		Cells.Range(func(k interface{}, v interface{}) bool {
			c := Cells.Get(k)
			wg.Add(1)
			go func(c *Cell) {
				c.Brain()
				wg.Done()
			}(c)
			return true
		})
		wg.Wait()

		Cells.Merge()
		Roads.Merge()

		removeDead()

		Cells.Merge()
		Roads.Merge()

		t := time.Now().Nanosecond() - s
		log.Printf("step %d took %f msec.", step, float64(t)/1e+6)

		//ログ処理はここ
		s = time.Now().Nanosecond()
		Visualize(step)
		JsonLogger(step)
		log.Printf("viaualization took %f msec.", float64(time.Now().Nanosecond()-s)/1e+6)

		if Cells.Len() == 0 {
			panic("全滅")
		}
	}
}

//その場所に細胞を置けるか
func canPut(p *Point) bool {
	//その点の周辺半径2の行列内に
	s := cutMatrix(p, 5)
	//1つ以上あれば
	if countMatrix(10, s) > 0 {
		return false
	}

	return true
}

//死んだ細胞は取り除く
func removeDead() {
	Cells.Range(func(key, value interface{}) bool {
		c := Cells.Get(key)
		if c.IsDead {
			//繋がっている道をすべて消す
			Roads.Range(func(key, value interface{}) bool {
				p := Roads.Get(key)
				if p.Node1.Id == c.Id || p.Node2.Id == c.Id {
					Roads.Delete(key)
				}

				return true
			})

			//細胞を消す
			Cells.Delete(c.Id)
		}

		return true
	})
}

//細胞の場所行列を更新する
func UpdateGrid() {
	Grid = mat.NewDense(config.WorldSizeX(), config.WorldSizeY(), nil)

	Cells.Range(func(key, value interface{}) bool {
		c := Cells.Get(key)
		Grid.Set(c.Point.X, c.Point.Y, 1)

		return true
	})
}

//そのセルがそのターンに回復できる量を計算する
func CalcRecover(p *Point) Resource {
	//その点の周りnマスを行列として切り出す
	s := cutMatrix(p, config.EffectDist()/2)
	//その行列にある1の数を数える
	r := countMatrix(config.EffectDist(), s)
	if r < 1 {
		panic("at least 1 Point must be exist in the area")
	}

	return newResource(int64(math.Round(float64(config.RecoverNormal()) / float64(r))))
}

//場所行列から一部を抜き出す
func cutMatrix(p *Point, r int) mat.Matrix {
	// |<---->|  |
	x1 := right(p.X-r, config.WorldSizeX())
	// |  |<---->|
	x2 := left(p.X+r, config.WorldSizeX())
	y1 := right(p.Y-r, config.WorldSizeY())
	y2 := left(p.Y+r, config.WorldSizeY())
	return Grid.Slice(x1, x2, y1, y2)
}

func right(x int, r int) int {
	if x < 0 {
		return 0
	}
	if x > r {
		return r
	}
	return x
}

func left(x int, l int) int {
	if x < 0 {
		return 0
	}
	if x > l {
		return l
	}
	return x
}

//場所行列内にある細胞の数を数える
func countMatrix(size int, s mat.Matrix) int {
	t := 0
	r, c := s.Dims()
	im := size
	if size > r {
		im = r
	}
	jm := size
	if size > c {
		jm = c
	}

	for i := 0; i < im; i++ {
		for j := 0; j < jm; j++ {
			t += int(s.At(i, j))
		}
	}
	return t
}

//Pathsからその細胞に繋がっているPathsを探して返す
func findPaths(c *Cell) *Paths {
	ps := Paths{}

	Roads.Range(func(key, value interface{}) bool {
		p := Roads.Get(key)
		if p.Node1.Id == c.Id || p.Node2.Id == c.Id {
			ps[p.Id] = p
		}

		return true
	})

	return &ps
}

//爆撃影響範囲を決める
func decideBombArea(step int64) {
	if config.BombFrequency() == 0.0 {
		BombPoint = nil
		return
	}
	hz := utils.Round(1 / config.BombFrequency())
	if step%hz == 0 {
		BombPoint = randomPoint()
	} else {
		BombPoint = nil
	}
}

//その地点が爆撃範囲かを返す
func isBombed(p *Point) bool {
	if BombPoint == nil {
		return false
	}
	//Xが範囲内
	if p.X >= BombPoint.X-config.BombRadius() && p.X <= BombPoint.X+config.BombRadius() {
		//Yが範囲内
		if p.Y >= BombPoint.Y-config.BombRadius() && p.Y <= BombPoint.Y+config.BombRadius() {
			return true
		}
	}

	return false
}

func searchNear(c *Cell) (*Cell, bool) {
	m := cutMatrix(c.Point, config.SearchRadius())
	for i := 0; i < config.SearchRadius(); i++ {
		for j := 0; j < config.SearchRadius(); j++ {
			//もし近くに細胞があって
			if int(m.At(i, j)) == 1 {
				id := fmt.Sprintf("%d,%d", i, j)
				//本当に細胞があって
				if Cells.Exists(id) {
					t := Cells.Get(fmt.Sprintf("%d,%d", i, j))
					ip := false
					for _, p := range *c.getPaths() {
						//そことの間に道が繋がってなかったら
						if p.Node1.Id == t.Id || p.Node2.Id == t.Id {
							ip = true
							break
						}
					}
					if ip == false {
						return t, true
					}
				}
			}
		}
	}
	return c, false
}
