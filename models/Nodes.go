package models

import "sync"

type Nodes struct {
	sm sync.Map
}

func (ns *Nodes) Exists(key interface{}) bool {
	_, ok := ns.sm.Load(key)
	return ok
}

func (ns *Nodes) Load(key interface{}) (*Cell, bool) {
	val, ok := ns.sm.Load(key)
	if !ok {
		return nil, false
	}

	return val.(*Cell), true
}

func (ns *Nodes) Store(key string, value *Cell) {
	ns.sm.Store(key, value)
}

func (ns *Nodes) Delete(key string) {
	ns.sm.Delete(key)
}

func (ns *Nodes) Len() int {
	i := 0
	ns.sm.Range(func(key, value interface{}) bool {
		i++
		return true
	})
	return i
}

func NewCells() *Nodes {
	return &Nodes{}
}
