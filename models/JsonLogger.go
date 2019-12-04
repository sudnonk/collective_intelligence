package models

import (
	"encoding/json"
	"fmt"
	"os"
)

func JsonLogger(step int64) {
	cs := map[string]Cell{}
	Cells.Range(func(key, value interface{}) bool {
		c := Cells.Get(key)
		cs[c.Id] = *c

		return true
	})

	ps := map[string]Path{}
	Roads.Range(func(key, value interface{}) bool {
		p := Roads.Get(key)
		ps[p.Id] = *p

		return true
	})

	js := map[string]interface{}{
		"Cells": cs,
		"Paths": ps,
	}

	j, err := json.Marshal(js)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(fmt.Sprintf("json/%d.json", step), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = f.Write(j)
	if err != nil {
		panic(err)
	}
}
