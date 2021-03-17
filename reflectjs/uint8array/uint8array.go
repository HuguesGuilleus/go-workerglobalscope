// Copyright (c) 2021, Hugues Guilleus. All rights reserved.
// Use of this source code is governed by a BBSD 3-Clause License
// that can be found in the LICENSE file.

package uint8array

import "syscall/js"

var Uint8Array js.Value = js.Global().Get("Uint8Array")

func New(b []byte) js.Value {
	a := Uint8Array.New(len(b))
	js.CopyBytesToJS(a, b)
	return a
}
