package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/HuguesGuilleus/go-workerglobalscope/console"
	"github.com/HuguesGuilleus/go-workerglobalscope/fetch"
	"github.com/HuguesGuilleus/go-workerglobalscope/message"
	"io"
)

func main() {
	// Console
	console.Log("Hidden message")
	console.Clear()

	console.Group("Console")

	console.Log("Log", 42, 36.125, true)
	console.Warn("Warn", 42, 36.125, true)
	console.Error("Error", 42, 36.125, true)
	console.Info("Info", 42, 36.125, true)
	console.Table("Table", 42, 36.125, true)

	console.Assert(true, "assert true", 686)
	console.Assert(false, "assert false", 686)

	console.Time("timer")
	console.TimeLog("timer", "here")
	console.TimeEnd("timer")

	console.GroupEnd()

	// Fetch
	rep := fetch.Fetch("README.md")

	// Print the headers
	console.Group("Headers:")
	for k, v := range rep.Headers.Map() {
		console.Log(k, v)
	}
	console.GroupEnd()

	// Hash the body and print it in hexadecimal.
	h := sha256.New()
	io.Copy(h, rep.Reader())
	console.Log("body sha256 hash:", hex.EncodeToString(h.Sum(nil)))

	message.Post(42)

	// message.Post(js.ValueOf(42))
	for m := range message.Event {
		console.Log(m)
	}
}
