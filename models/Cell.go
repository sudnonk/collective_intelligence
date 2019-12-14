package models

import (
	"github.com/sudnonk/collective_intelligence/debug"
	"math/rand"

	"github.com/sudnonk/collective_intelligence/config"
	"github.com/sudnonk/collective_intelligence/utils"
)

//細胞
type Cell struct {
	Id       string    `json:"id"`
	Point    *Point    `json:"point"`
	Resource Resource  `json:"resource"`
	State    cellState `json:"state"`
	Persona  *Persona  `json:"persona"`
	IsDead   bool      `json:"is_dead"`
}

//細胞の状態
type cellState int64

const (
	needsResource cellState = 1 //資源が少ないので助けが必要
	healthy       cellState = 0 //健康
)

//その細胞の資源を回復する
func (c *Cell) recover() {
	//回復量
	c.Resource += CalcRecover(c.Point)
}

//その細胞の資源を別の細胞に渡すためにパスに渡す
func (c *Cell) give(to *Path, a Resource) {
	if c.Resource > a {
		to.in(a, c)
		c.Resource -= a
	} else {

	}
}

//Pathから資源を受け取る
func (c *Cell) receive(a Resource) {
	c.Resource += a
}

func (c *Cell) getPaths() *Paths {
	return findPaths(c)
}

//この細胞の思考回路
func (c *Cell) Brain() {
	//ダメージを受ける
	c.bombed()

	//死亡判定
	if c.deathDetermination() {
		return
	}

	//回復
	c.recoverAll()

	//助けが必要か
	if c.needsHelp() {
		//助けが必要なら何もしない
		debug.Printf("%s needs help.", c.Id)
		return
	}

	//周りを助ける
	if c.helpOthers() {
		debug.Printf("%s helped neighbors.", c.Id)

		return
	}

	//道を拡張するか、新しい細胞を作る
	if c.makeNewCell() {
		//新しい細胞を作れたら終わり
		debug.Printf("%s made new cell.", c.Id)
		return
	} else if c.connectNear() {
		//近くの細胞との間に道を作る。資源が足りなかったら何もしない。
		debug.Printf("%s connect near.", c.Id)
		return
	} else if c.upgradePath() {
		//道をアップグレードする。資源が足りなかったら何もしない
		debug.Printf("%s upgrades path.", c.Id)
		return
	}

	debug.Printf("%s did nothing.", c.Id)
	return
}

//爆撃の範囲内ならダメージを受ける
func (c *Cell) bombed() {
	if isBombed(c.Point) {
		c.Resource -= BombDamage()
	}
}

//死亡判定、死んでたらtrue
func (c *Cell) deathDetermination() bool {
	//もし死んでたらスキップ
	if c.IsDead {
		return true
	}

	//もし生命維持に必要な最低量より資源が少なくなれば
	if c.Resource < ResourceLimit() {
		//死ぬ
		c.IsDead = true
		return true
	} else {
		c.IsDead = false
		return false
	}
}

//全ての手段で資源を回復する
func (c *Cell) recoverAll() {
	//定時回復
	c.State = healthy
	c.recover()
	for _, p := range *c.getPaths() {
		//自身に向けて輸送中の資源を受け取る
		c.receive(p.out(c))
	}
	c.Resource.adjust()
}

//もし自分の資源が恐怖を感じるほど少なければ、周りに助けを求める
func (c *Cell) needsHelp() bool {
	if c.Resource.toFloat64() < ResourceMax().toFloat64()*c.Persona.Fear {
		//周りに助けを求める
		c.State = needsResource

		return true
	} else {
		c.State = healthy

		return false
	}
}

//周りに助けを求めている人が居れば助ける
func (c *Cell) helpOthers() bool {
	//接続している道
	for _, p := range *c.getPaths() {
		//道の向こうの細胞
		n := p.otherSide(c)
		if n == nil {
			continue
		}

		//もし隣人が助けを求めていたら
		if n.State == needsResource {
			//やさしさの分だけ資源を渡す
			w := utils.Round(c.Resource.toFloat64() * c.Persona.Kindness)
			a := p.in(newResource(w), c)
			c.Resource -= a

			//その人を助けたら終わり
			return true
		}
	}
	return false
}

//新しい細胞とその間の道を作る。作れるだけの資源があって、作れたらtrue
func (c *Cell) makeNewCell() bool {
	cost := WidthCost() + CellCost()

	//もし必要なコストを賄えるだけの資源があれば
	if c.Resource-ResourceLimit() > cost {
		//新しい細胞に渡す資源
		a := (c.Resource - cost) / 2

		c2 := newCell(c, a)
		//適切な場所に細胞を配置できない場合はランダムに配置する
		i := 0
		for {
			i++
			if c2.isValid() {
				break
			}

			c2 = newCellRandom(c, a)
			if i > 100 {
				debug.Printf("no space around %s.", c.Id)

				return false
			}
		}

		//それぞれを道でつなぐ
		p1 := connect(c, c2)
		p2 := connect(c2, c)
		Roads.Set(p1.Id, p1)
		Roads.Set(p2.Id, p2)

		Cells.Set(c2.Id, c2)

		c.Resource = c.Resource - cost - a

		return true
	} else {
		debug.Printf("not enough resource %s", c.Id)
		return false
	}
}

func newCell(c *Cell, a Resource) *Cell {
	arg, d := decideNewPath(*c)
	po := calcNewPoint(arg, d, c.Point)

	return &Cell{
		Id:       po.identify(),
		Point:    po,
		Resource: a,
		State:    healthy,
		Persona:  newPersona(c.Persona),
		IsDead:   false,
	}
}

func newCellRandom(c *Cell, a Resource) *Cell {
	arg, d := randomNewPath()
	po := calcNewPoint(arg, d, c.Point)

	return &Cell{
		Id:       po.identify(),
		Point:    po,
		Resource: a,
		State:    healthy,
		Persona:  newPersona(c.Persona),
		IsDead:   false,
	}
}

//道幅を広げる
func (c *Cell) upgradePath() bool {
	//もし道を広げられるだけの資源があって
	if c.Resource > WidthCost() {
		for _, p := range *c.getPaths() {
			//広げて欲しがっている道があれば
			if p.WantExpand > 0.8 {
				if p.canExpand() {
					//広げて終わり
					c.Resource -= WidthCost()
					p.expand()
					return true
				}
			}
		}
	}

	return false
}

//近くの細胞との間に道を作る
func (c *Cell) connectNear() bool {
	//もし資源が十分あって
	if c.Resource-ResourceLimit() > 2*WidthCost() {
		n, f := searchNear(c)
		//近くに繋がってない細胞があれば
		if f == false {
			debug.Printf("no cells near %s", c.Id)
			return false
		}

		//繋げる
		connect(c, n)
		c.Resource -= 2 * WidthCost()
		return true
	} else {
		debug.Printf("not enough resource in %s", c.Id)
		return false
	}
}

func (c *Cell) isValid() bool {
	if c.Point.X >= config.WorldSizeX() || c.Point.Y >= config.WorldSizeY() {
		return false
	}
	if c.Point.X < 0 || c.Point.Y < 0 {
		return false
	}
	if !canPut(c.Point) {
		return false
	}

	return true
}

//新しい細胞を作るのに必要なコスト（定数）
func CellCost() Resource {
	return Resource(config.CellCost())
}

//ランダムに細胞を生成する
func GenerateRandom() *Cell {
	for {
		po := randomPoint()
		c := &Cell{
			Id:       po.identify(),
			Point:    po,
			Resource: newResource(rand.Int63n(ResourceMax().toInt64()-config.ResourceLimit()) + config.ResourceLimit()),
			State:    healthy,
			Persona:  randomPersona(),
			IsDead:   false,
		}

		if c.isValid() {
			return c
		}
	}
}
