package models

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"os"
)

const stretch = 8
const diameter = 5

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
			canvas.Line(p.Node1.Point.X*stretch, p.Node1.Point.Y*stretch, p.Node2.Point.X*stretch, p.Node2.Point.Y*stretch, fmt.Sprintf("stroke='%s' stroke-width='%d'", cop, p.Width*2))
			ps[p.Id] = struct{}{}
		}

		var coc string
		if c.State == needsResource {
			coc = "red"
		} else {
			coc = "black"
		}
		rate := (c.Resource.toFloat64() / ResourceMax().toFloat64()) * 100

		//扇形
		canvas.Circle(c.Point.X*stretch, c.Point.Y*stretch, diameter, fmt.Sprintf("stroke='%s' stroke-width='%d' stroke-dasharray='%f,%f' fill='none'", coc, diameter*2, rate, 100-rate))
		//外枠
		canvas.Circle(c.Point.X*stretch, c.Point.Y*stretch, diameter*2, fmt.Sprintf("stroke='%s' stroke-width='1' fill='none'", coc))

		return true
	})

	canvas.End()
}
