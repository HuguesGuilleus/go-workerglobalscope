// Copyright (c) 2021, Hugues Guilleus. All rights reserved.
// Use of this source code is governed by a BBSD 3-Clause License
// that can be found in the LICENSE file.

package ws

import (
	"syscall/js"
	"time"
)

var (
	Uint8Array js.Value = js.Global().Get("Uint8Array")
	Blob       js.Value = js.Global().Get("Blob")
	Object     js.Value = js.Global().Get("Object")
	Array      js.Value = js.Global().Get("Array")
	Date       js.Value = js.Global().Get("Date")
)

func NewUint8Array(b []byte) js.Value {
	a := Uint8Array.New(len(b))
	js.CopyBytesToJS(a, b)
	return a
}

// NewBlob create a new blob with the mime type mime and the blob content is b.
func NewBlob(mime string, b []byte) js.Value {
	opt := Object.New()
	opt.Set("type", mime)
	return Blob.New(Array.Call("of", NewUint8Array(b)), opt)
}

// NewDate create a js.Value of type Date from t.
func NewDate(t time.Time) js.Value {
	if t.IsZero() {
		return Date.New(0)
	}
	j, _ := t.MarshalText()
	return Date.New(string(j))
}
