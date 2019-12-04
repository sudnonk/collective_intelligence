package models

import (
	"github.com/sudnonk/collective_intelligence/utils"
)

type MutexPaths struct {
	mm utils.MutexMap
}

func (ps *MutexPaths) Exists(key interface{}) bool {
	return ps.mm.Exists(key)
}

func (ps *MutexPaths) Get(key interface{}) *Path {
	return ps.mm.Get(key).(*Path)
}

func (ps *MutexPaths) Set(key string, value *Path) {
	ps.mm.Set(key, value)
}

func (ps *MutexPaths) Delete(key interface{}) {
	ps.mm.Delete(key)
}

func (ps *MutexPaths) Range(f func(key interface{}, value interface{}) bool) {
	ps.mm.Range(f)
}

func (ps *MutexPaths) Merge() {
	ps.mm.Merge()
}

func NewPaths() *MutexPaths {
	return &MutexPaths{
		utils.NewMutexMap(),
	}
}
