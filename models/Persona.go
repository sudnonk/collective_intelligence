package models

import (
	"github.com/sudnonk/collective_intelligence/config"
	"github.com/sudnonk/collective_intelligence/utils"
	"math"
)

//ある細胞の性格
type Persona struct {
	Kindness float64 `json:"kindness"` //やさしさ。助けを求められたときにどれぐらい与えるか
	Fear     float64 `json:"fear"`     //臆病さ。資源がどれぐらい減ったら周りに助けを求めるか
}

//ある性格を基に新しい性格を作る
func newPersona(p *Persona) *Persona {
	return &Persona{
		Kindness: calcMutation(p.Kindness),
		Fear:     calcMutation(p.Fear),
	}
}

//性格をランダムに変動させる
func calcMutation(f float64) float64 {
	r := config.MutationRate()

	max := math.Min(f*(1+r), 1.0)
	min := math.Max(f*(1-r), 0.0)

	return utils.RandFloat64n(min, max)
}

//性格をランダムに作る
func randomPersona() *Persona {
	return &Persona{
		Kindness: utils.RandFloat64n(0, 1),
		Fear:     utils.RandFloat64n(0, 1),
	}
}
