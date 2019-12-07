package models

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"github.com/sudnonk/collective_intelligence/utils"
	"math"
	"os"
)

const stretch = 8
const radius = 5

func Visualize(step int64) {
	fn := fmt.Sprintf("svgs/%d.svg", step)
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		panic(err)
	}

	width := 100*stretch + 100
	height := 100*stretch + 100
	canvas := svg.New(f)
	canvas.Start(width, height)

	canvas.Text(15, 15, fmt.Sprintf("個数：%d", Cells.Len()))
	canvas.Text(15, 30, fmt.Sprintf("ステップ：%d", step))

	if Cells.Len() == 0 {
		canvas.Text(50*stretch, 50+stretch, "全滅")
	}

	ps := map[string]struct{}{}

	Cells.Range(func(key, value interface{}) bool {
		c := Cells.Get(key)

		for _, p := range *c.getPaths() {
			if _, ok := ps[p.Id]; ok {
				continue
			}

			var cop string
			if p.Transit == 0 {
				cop = "blue"
			} else {
				cop = "red"
			}
			canvas.Line(p.Node1.Point.X*stretch, p.Node1.Point.Y*stretch, p.Node2.Point.X*stretch, p.Node2.Point.Y*stretch, fmt.Sprintf("id='p-%s' stroke='%s' stroke-width='%d'", p.Id, cop, p.Width*2))
			ps[p.Id] = struct{}{}
		}

		var coc string
		if c.State == needsResource {
			coc = "red"
		} else {
			coc = "black"
		}
		//最大体力が2π
		arg := (2 * math.Pi / ResourceMax().toFloat64()) * c.Resource.toFloat64()

		canvas.Group(fmt.Sprintf("id='c-%s'", c.Id))
		//扇形
		canvas.Path(makeSectorD(*c.Point, arg, radius, stretch), fmt.Sprintf("fill='%s'", coc))
		//外枠
		canvas.Circle(c.Point.X*stretch, c.Point.Y*stretch, radius*2, fmt.Sprintf("stroke='%s' stroke-width='1' fill='none'", coc))
		//クリック用の透明な円
		canvas.Circle(c.Point.X*stretch, c.Point.Y*stretch, radius*2, fmt.Sprintf("fill='none'"))
		canvas.Gend()

		//爆撃範囲
		canvas.Rect(BombPoint.X*stretch, BombPoint.Y*stretch, 10*stretch, 10*stretch, fmt.Sprintf("fill='yellow'"))

		return true
	})

	canvas.End()
}

//SVGで扇形を書く時のdを生成する
func makeSectorD(center Point, arg float64, radius int, stretch int) string {
	center.X *= stretch
	center.Y *= stretch

	start := Point{
		X: center.X,
		Y: center.Y + radius,
	}

	end := Point{
		X: int(utils.Round(float64(center.X) + math.Sin(arg)*float64(radius))),
		Y: int(utils.Round(float64(center.Y) - math.Cos(arg)*float64(radius))),
	}

	var pattern string
	if arg > math.Pi {
		pattern = fmt.Sprintf("%d %d", 1, 1)
	} else {
		pattern = fmt.Sprintf("%d %d", 0, 1)
	}

	//M中心座標 L始まり座標 A半径 0 パターン 終わり座標z
	return fmt.Sprintf("M %d,%d L %d,%d A %d %d 0 %s %d %d z", center.X, center.Y, start.X, start.Y, radius, radius, pattern, end.X, end.Y)
}
