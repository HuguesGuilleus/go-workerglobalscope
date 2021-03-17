// Copyright (c) 2021, Hugues Guilleus. All rights reserved.
// Use of this source code is governed by a BBSD 3-Clause License
// that can be found in the LICENSE file.

package message

import (
	"syscall/js"
)

var postMessage = js.Global().Get("postMessage")

// Send a message to the parent of this worker.
//
//	self.postMessage(m) // in Javascript
func Post(m interface{}) {
	postMessage.Invoke(m)
}

// Send the field data of the message.
//
// It use self.addEventListener method, so several eventhandler can be used.
var Event <-chan js.Value = func() <-chan js.Value {
	c := make(chan js.Value)

	js.Global().Call("addEventListener", "message", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		c <- args[0].Get("data")
		return nil
	}))

	return c
}()
