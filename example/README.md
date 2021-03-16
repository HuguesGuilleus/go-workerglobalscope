# Build

```
GOOS=js GOARCH=wasm go build -o worker.wasm worker.go
```

# HTTP

Usage of WebAssmbly API in js need a secure context (HTTPS of localhost). Example of a HTTP server for localhost in GO:

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
```
