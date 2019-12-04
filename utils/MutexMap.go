package utils

type MutexMap struct {
	readonly map[interface{}]interface{}
	dirty    map[interface{}]interface{}
}

// ---- readonly ----

//readonlyに値があるか調べる
func (mm MutexMap) Exists(key interface{}) bool {
	_, ok := mm.readonly[key]
	return ok
}

//readonlyから値を取得する
func (mm MutexMap) Get(key interface{}) interface{} {
	if mm.Exists(key) {
		return mm.readonly[key]
	} else {
		return nil
	}
}

//全要素に対して引数の関数を実行する
func (mm MutexMap) Range(f func(key interface{}, value interface{}) bool) {
	for key := range mm.readonly {
		if !f(key, mm.Get(key)) {
			break
		}
	}
}

// ---- dirty -----

//dirtyに値を入れる
func (mm *MutexMap) Set(key interface{}, value interface{}) {
	mm.dirty[key] = value
}

func (mm *MutexMap) Delete(key interface{}) {
	if mm.Exists(key) {
		delete(mm.dirty, key)
	}
}

//dirtyの内容をreadonlyに反映する
func (mm *MutexMap) Merge() {
	for key := range mm.readonly {
		delete(mm.readonly, key)
	}
	for key, value := range mm.dirty {
		mm.readonly[key] = value
	}
}
