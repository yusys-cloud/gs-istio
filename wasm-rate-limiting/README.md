# wasm-rate-limiting

A sample Istio WebAssembly plugin for rate limiting.

Istio 1.12 release introduces new Wasm Extension API. This folder contains a sample application of
implementing rate limiting in Golang, and deploy the Wasm Plugin using Istio API. The plugin will
enforce rate limiting for 2 request per second. Extra request beyond the limit will be rejected.

<!-- TODO(incfly): provide the actual link once the blog is ready. -->
For detailed instructions, checkout tetrate.io/blog.
## build
``` 
go mod edit -require=github.com/tetratelabs/proxy-wasm-go-sdk@main
go mod download github.com/tetratelabs/proxy-wasm-go-sdk
tinygo build -o main.wasm -scheduler=none -target=wasi ./main.go
python3 -m http.server 8088
```
