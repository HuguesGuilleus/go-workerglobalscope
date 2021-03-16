// Copyright (c) 2021, Hugues Guilleus. All rights reserved.
// Use of this source code is governed by a BBSD 3-Clause License
// that can be found in the LICENSE file.

package fetch

import (
	"syscall/js"
)

type Headers struct {
	js.Value
}

func (h Headers) Get(name string) string {
	v := h.Value.Get(name)
	if !v.Truthy() {
		return ""
	}
	return v.String()
}

func (h Headers) Map() (m map[string]string) {
	m = make(map[string]string)
	entries := h.Value.Call("entries")

	for {
		n := entries.Call("next")
		if n.Get("done").Bool() {
			return
		}
		v := n.Get("value")
		m[v.Index(0).String()] = v.Index(1).String()
	}
}
