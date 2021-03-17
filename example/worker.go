package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/HuguesGuilleus/go-workerglobalscope/console"
	"github.com/HuguesGuilleus/go-workerglobalscope/fetch"
	"github.com/HuguesGuilleus/go-workerglobalscope/message"
	"io"
	"time"
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

	s := &struct {
		Array      [2]time.Time
		Bool       bool
		BigInt     int64
		Float      float64
		Set        map[string]bool
		Int        int
		Lang       map[string]string
		String     string
		Slice      []time.Time
		Time       time.Time
		unexported string
	}{
		Array:  [2]time.Time{time.Now()},
		Bool:   true,
		BigInt: int64(0xFFFF_FFFF),
		Float:  56.89,
		Set: map[string]bool{
			"Hello": true,
			"Hola":  true,
			"Salut": true,
		},
		Int: 42,
		Lang: map[string]string{
			"en": "English",
			"es": "Español",
			"fr": "Français",
		},
		String:     "Yolo!",
		Slice:      []time.Time{time.Now().Add(-24 * time.Hour)},
		Time:       time.Now(),
		unexported: "Hello World",
	}

	message.Post(&s)

	for m := range message.Listen() {
		console.Log(m)
	}
}
