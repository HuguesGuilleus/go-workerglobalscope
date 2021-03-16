# go-workerglobalscope

A wrapper for the WorkerGlobalScope, for go in WebAssembly. The console package and the fetch package can be used in *main js worker*, not only in a Worker.

To minimise some init syscall, we separate the API in different package.

## Warning

- This module is based on a experimental package without compatibilty, so it can be break in future Realease of Go. We use Go 1.16
- Not all API and all properties and methods will be wrapper, but you can complete it for your own needs.
- The goal is to support modern browser that, it suppose to support fetch API...
