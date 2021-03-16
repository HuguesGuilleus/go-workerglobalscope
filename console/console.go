// Copyright (c) 2021, Hugues Guilleus. All rights reserved.
// Use of this source code is governed by a BBSD 3-Clause License
// that can be found in the LICENSE file.

// A part of the operation over the console.
package console

import (
	"syscall/js"
)

var console = js.Global().Get("console")

func Clear() { console.Call("clear") }

func Log(args ...interface{})   { console.Call("log", args...) }
func Warn(args ...interface{})  { console.Call("warn", args...) }
func Error(args ...interface{}) { console.Call("error", args...) }
func Info(args ...interface{})  { console.Call("info", args...) }
func Table(args ...interface{}) { console.Call("table", args...) }

func Time(id string)    { console.Call("time", id) }
func TimeEnd(id string) { console.Call("timeEnd", id) }
func TimeLog(id string, args ...interface{}) {
	console.Call("timeLog", append([]interface{}{id}, args...)...)
}

// If condition is false, print the args with Error.
func Assert(condition bool, args ...interface{}) {
	if condition == false {
		Error(args...)
	}
}
