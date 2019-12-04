package models

import (
	"fmt"
	"sort"
)

type Path struct {
	Id     string  `json:"id"`
	Width  Width   `json:"width"`
	Length float64 `json:"length"`
	//そのPathの角度
	Arg   Radian `json:"arg"`
	Node1 *Cell  `json:"node_1"`
	Node2 *Cell  `json:"node_2"`
	//輸送中の資源
	Transit Resource `json:"transit"`
	//輸送中の資源の出所のノード
	From *Cell `json:"from"`
	//この道の広げて欲しさ
	WantExpand float64 `json:"want_expand"`
}

//細胞から受け取った資源を輸送中にして、実際に受け取った量を返す
func (p *Path) in(a Resource, c *Cell) Resource {
	var r Resource
	//受け取った量より道幅の方が狭ければ
	if a.toInt64() > p.Width.toInt64() {
		//道幅だけ受け入れる
		r = newResource(p.Width.toInt64())
		//広げて欲しさを少し上げる
		if p.canExpand() {
			p.WantExpand += 0.1
		}
	} else {
		r = a
	}

	p.Transit = r

	if isSame(p.Node1.Point, c.Point) {
		p.From = p.Node1
	} else {
		p.From = p.Node2
	}

	return r
}

//輸送中の資源を細胞に渡す
func (p *Path) out(c *Cell) Resource {
	//もし輸送中じゃないなら
	if p.Transit == 0 {
		return 0
	} else {
		//もし渡す先が受け取った先と同じなら
		if isSame(c.Point, p.From.Point) {
			//何も渡さない
			return 0
		} else {
			r := p.Transit
			p.Transit = 0
			p.From = nil

			return r
		}
	}
}

//cとは逆側の細胞を返す
func (p *Path) otherSide(c *Cell) *Cell {
	if isSame(p.Node1.Point, c.Point) {
		return p.Node2
	} else {
		return p.Node1
	}
}

//道をこれ以上広げることができるか
func (p *Path) canExpand() bool {
	return p.Width < MaxWidth()
}

//道を広げて、広げて欲しさをリセットする
func (p *Path) expand() {
	p.Width += 1
	p.WantExpand = 0
}

func connect(from *Cell, to *Cell) *Path {
	return &Path{
		Id:         identify(from.Point, to.Point),
		Width:      1,
		Length:     calcDistance(from.Point, to.Point),
		Arg:        calcArg(from.Point, to.Point),
		Node1:      from,
		Node2:      to,
		Transit:    0,
		From:       nil,
		WantExpand: 0,
	}
}

func identify(from *Point, to *Point) string {
	ps := []int{from.X, from.Y, to.X, to.Y}
	sort.Ints(ps)
	return fmt.Sprintf("%d,%d,%d,%d", ps[0], ps[1], ps[2], ps[3])
}
