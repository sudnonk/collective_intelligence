package models

import (
	"github.com/sudnonk/collective_intelligence/config"
)

type Resource int64

//その資源の値が正しい範囲内か
func (r Resource) isValid() bool {
	if r < ResourceMin() || r > ResourceMax() {
		return false
	} else {
		return true
	}
}

func (r Resource) toInt64() int64 {
	return int64(r)
}

func (r Resource) toFloat64() float64 {
	return float64(r)
}

func (r *Resource) adjust() {
	if *r > ResourceMax() {
		*r = ResourceMax()
	} else if *r < ResourceMin() {
		*r = ResourceMin()
	}
}

//int64から資源にする
func newResource(a int64) Resource {
	r := Resource(a)
	r.adjust()
	return r
}

//資源の下限（定数）
func ResourceMin() Resource {
	return Resource(config.ResourceMin())
}

//資源の上限（定数）
func ResourceMax() Resource {
	return Resource(config.ResourceMax())
}

//生命維持に必要な最低量
func ResourceLimit() Resource {
	return Resource(config.ResourceLimit())
}
