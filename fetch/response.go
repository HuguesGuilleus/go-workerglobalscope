// Copyright (c) 2021, Hugues Guilleus. All rights reserved.
// Use of this source code is governed by a BBSD 3-Clause License
// that can be found in the LICENSE file.

package fetch

import (
	"syscall/js"
)

func Fetch(url string) Response {
	rep := Await(js.Global().Call("fetch", url))
	status := rep.Get("status").Int()
	return Response{
		Status:     status,
		StatusText: rep.Get("statusText").String(),
		Ok:         200 <= status && status < 300,
		Body:       Body{Value: rep},
		Headers:    Headers{Value: rep.Get("headers")},
	}
}

// https://fetch.spec.whatwg.org/#response-class
type Response struct {
	Body
	Headers
	Status     int
	StatusText string
	Ok         bool
}

func Await(promise js.Value) (resolve js.Value) {
	c := make(chan struct{})

	promise.Call("then", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		resolve = args[0]
		c <- struct{}{}
		return nil
	}))

	<-c

	return
}
