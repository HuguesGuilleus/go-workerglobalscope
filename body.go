// Copyright (c) 2021, Hugues Guilleus. All rights reserved.
// Use of this source code is governed by a BBSD 3-Clause License
// that can be found in the LICENSE file.

package ws

import (
	"io"
	"syscall/js"
)

// ReadBody creates a io.Reader with the Body.arrayBuffer().
// The error implement js.JSValue
//
// Body is a mixin for a Blob (File inherit of a Blob) or a fetch response.
// https://fetch.spec.whatwg.org/#body
func ReadBody(body js.Value) (io.Reader, error) {
	arrayBuffer, err := AwaitError(body.Call("arrayBuffer"))
	if err != nil {
		return nil, err
	}
	uint8array := Uint8Array.New(arrayBuffer)

	return &bodyReader{
		array: uint8array,
		size:  uint8array.Get("byteLength").Int(),
		pos:   0,
	}, nil
}

type bodyReader struct {
	array js.Value
	size  int
	pos   int
}

func (b *bodyReader) Read(dst []byte) (int, error) {
	var readed int
	if b.pos >= b.size {
		return 0, io.EOF
	} else if b.pos == 0 {
		readed = js.CopyBytesToGo(dst, b.array)
	} else {
		readed = js.CopyBytesToGo(dst, b.array.Call("subarray", b.pos))
	}
	b.pos += readed
	return readed, nil
}
