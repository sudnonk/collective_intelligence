package models

import (
	"github.com/sudnonk/collective_intelligence/utils"
)

type Nodes struct {
	mm utils.MutexMap
}

func (ns *Nodes) Exists(key interface{}) bool {
	return ns.mm.Exists(key)
}

func (ns *Nodes) Get(key interface{}) *Cell {
	return ns.mm.Get(key).(*Cell)
}

func (ns *Nodes) Set(key string, value *Cell) {
	ns.mm.Set(key, value)
}

func (ns *Nodes) Delete(key string) {
	ns.mm.Delete(key)
}

func (ns *Nodes) Len() int {
	i := 0
	ns.mm.Range(func(key, value interface{}) bool {
		i++
		return true
	})
	return i
}

func (ns *Nodes) Range(f func(key interface{}, value interface{}) bool) {
	ns.mm.Range(f)
}

func (ns *Nodes) Merge() {
	ns.mm.Merge()
}

func NewCells() *Nodes {
	return &Nodes{}
}
