package main

import (
	"github.com/sudnonk/collective_intelligence/config"
	"github.com/sudnonk/collective_intelligence/models"
)

func main() {
	config.Set(config.Config{
		ResourceMax:   314,
		ResourceMin:   0,
		ResourceLimit: 10,
		RecoverNormal: 100,
		MaxWidth:      10,
		MinWidth:      1,
		WidthCost:     20,
		CellCost:      100,
		MutationRate:  0.1,
		WorldSizeX:    100,
		WorldSizeY:    100,
		MaxStep:       1000,
		MinDist:       5,
		EffectDist:    40,
	})

	models.Run()
}
