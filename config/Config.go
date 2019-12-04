package config

type Config struct {
	ResourceMax   int64   //資源の上限
	ResourceMin   int64   //資源の下限
	ResourceLimit int64   //細胞が生存しているのに必要な資源の下限
	RecoverNormal int64   //標準回復量
	MaxWidth      int64   //道幅の上限
	MinWidth      int64   //道幅の下限
	WidthCost     int64   //道幅を広げるのに必要なコスト
	CellCost      int64   //新しい細胞を作るのに必要なコスト
	MutationRate  float64 //性格の変動量
	WorldSizeX    int     //世界のサイズ（横）
	WorldSizeY    int     //世界のサイズ（盾）
	MaxStep       int64   //世界の寿命
	MinDist       float64 //細胞同士が最低限離れている距離
	EffectDist    int     //栄養計算に使う細胞の範囲の距離
}

var config Config

func Set(c Config) {
	config = c
	if ResourceMax() < CellCost()+WidthCost() {
		panic("parameter invalid: cells cannot create any cells due to high cost.")
	}
}

func ResourceMax() int64 {
	return config.ResourceMax
}

func ResourceMin() int64 {
	return config.ResourceMin
}

func ResourceLimit() int64 {
	return config.ResourceLimit
}

func RecoverNormal() int64 {
	return config.RecoverNormal
}

func MaxWidth() int64 {
	return config.MaxWidth
}

func MinWidth() int64 {
	return config.MinWidth
}

func WidthCost() int64 {
	return config.WidthCost
}

func CellCost() int64 {
	return config.CellCost
}

func MutationRate() float64 {
	return config.MutationRate
}

func WorldSizeX() int {
	return config.WorldSizeX
}

func WorldSizeY() int {
	return config.WorldSizeY
}

func MaxStep() int64 {
	return config.MaxStep
}

func MinDist() float64 {
	return config.MinDist
}

func EffectDist() int {
	if config.EffectDist%2 != 0 {
		panic("config.EffectDist must be 2の倍数")
	}
	return config.EffectDist
}
