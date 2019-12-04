package models

import (
	crand "crypto/rand"
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

func cutMatrix(p *Point, r int) mat.Matrix {
	x1 := utils.Min(utils.Max(p.X-r, 0), config.WorldSizeX()-r*2)
	x2 := utils.Min(utils.Max(p.X+r, r*2), config.WorldSizeX())
	y1 := utils.Min(utils.Max(p.Y-r, 0), config.WorldSizeY()-r*2)
	y2 := utils.Min(utils.Max(p.Y+r, r*2), config.WorldSizeY())
	return Grid.Slice(x1, x2, y1, y2)
}

func countMatrix(size int, s mat.Matrix) int {
	r := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			r += int(s.At(i, j))
		}
	}
	return r
}

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
