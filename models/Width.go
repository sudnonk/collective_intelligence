package models

import (
	"fmt"
	"github.com/sudnonk/collective_intelligence/config"
)

type Width int64

//道幅を1広げる
func (w *Width) expand() error {
	e := *w + 1
	if e > MaxWidth() {
		return fmt.Errorf("width already maximum")
	} else {
		*w = e

		return nil
	}
}

func (w Width) isValid() bool {
	if w < MinWidth() || w > MaxWidth() {
		return false
	} else {
		return true
	}
}

func (w Width) toInt64() int64 {
	return int64(w)
}

func MaxWidth() Width {
	return Width(config.MaxWidth())
}

func MinWidth() Width {
	return Width(config.MinWidth())
}

func WidthCost() Resource {
	return Resource(config.WidthCost())
}
