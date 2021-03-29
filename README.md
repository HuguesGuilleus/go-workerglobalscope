# go-workerglobalscope

[![Go Reference](https://pkg.go.dev/badge/github.com/HuguesGuilleus/go-workerglobalscope.svg)](https://pkg.go.dev/github.com/HuguesGuilleus/go-workerglobalscope)

A wrapper for go in WebAssembly in a Worker context to expose some API (console, fetch) and some Js value (Date, Uint8Array). The console package and the fetch package can be used in *main js worker*, not only in a Worker.

## Warning

- This module is based on a experimental package without compatibilty, so it can be break in future Realease of Go. We use Go 1.16
- Not all API and all properties and methods will be wrapper, but you can complete it for your own needs.
- The goal is to support modern browser that, it suppose to support fetch API...
