// Copyright (c) 2021, Hugues Guilleus. All rights reserved.
// Use of this source code is governed by a BBSD 3-Clause License
// that can be found in the LICENSE file.

package fetch

import (
	"io"
	"syscall/js"
)

var Uint8Array = js.Global().Get("Uint8Array")

// You can use it with a response, a Blob for example.
//
// https://fetch.spec.whatwg.org/#body
type Body struct {
	js.Value
}

// Call text method and resolve the promise.
func (b *Body) Text() string {
	return Await(b.Call("text")).String()
}

// Get an array buffer of the Body, and to each call of Read copy the bytes
// into the destination []byte.
func (b *Body) Reader() io.Reader {
	a := Uint8Array.New(Await(b.Call("arrayBuffer")))
	return &bodyReader{
		array: a,
		size:  a.Get("byteLength").Int(),
		pos:   0,
	}
}

type bodyReader struct {
	array js.Value
	size  int
	pos   int
}

func (b *bodyReader) Read(dst []byte) (int, error) {
	if b.pos >= b.size {
		return 0, io.EOF
	}
	readed := js.CopyBytesToGo(dst, b.array)
	b.pos += readed
	return readed, nil
}
