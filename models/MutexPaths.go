package models

import "sync"

type MutexPaths struct {
	sm sync.Map
}

func (ps *MutexPaths) Exists(key interface{}) bool {
	_, ok := ps.sm.Load(key)
	return ok
}

func (ps *MutexPaths) Load(key interface{}) *Path {
	val, ok := ps.sm.Load(key)
	if !ok {
		return nil
	}

	return val.(*Path)
}

func (ps *MutexPaths) Store(key string, value *Path) {
	ps.sm.Store(key, value)
}

func (ps *MutexPaths) Delete(key interface{}) {
	ps.sm.Delete(key)
}

func (ps *MutexPaths) Len() int {
	i := 0
	ps.sm.Range(func(key, value interface{}) bool {
		i++
		return true
	})
	return i
}

func NewPaths() *MutexPaths {
	return &MutexPaths{}
}
