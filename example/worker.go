package main

import (
	"github.com/HuguesGuilleus/go-workerglobalscope/console"
)

func main() {
	// Console
	console.Log("Hidden message")
	console.Clear()

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
}
