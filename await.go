// Copyright (c) 2021, Hugues Guilleus. All rights reserved.
// Use of this source code is governed by a BBSD 3-Clause License
// that can be found in the LICENSE file.

package ws

import (
	"syscall/js"
)

// Resolve the promise.
func Await(promise js.Value) (resolve, reject js.Value) {
	c := make(chan struct{})
	defer close(c)

	then := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		resolve = args[0]
		c <- struct{}{}
		return nil
	})
	defer then.Release()
	promise.Call("then", then)

	catch := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		reject = args[0]
		c <- struct{}{}
		return nil
	})
	defer catch.Release()
	promise.Call("catch", catch)

	<-c
	return
}

// AwaitError resolve the promise. If fail the error contain the js object and
// his string value.
func AwaitError(promise js.Value) (resolve js.Value, err error) {
	var reject js.Value
	resolve, reject = Await(promise)
	if reject.Truthy() {
		err = promiseError{
			Value: reject,
			err:   reject.String(),
		}
	}
	return
}

type promiseError struct {
	js.Value
	err string
}

func (pe promiseError) Error() string {
	return pe.err
}
