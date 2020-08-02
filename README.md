# Golang Wasm Experiment

Problems: this JSON prettifier creates a 2.7MB wasm file just to prettify JSON. I tried to compile it with tinygo, but that doesn't support all the needed libraries.

```text
$ make build
$ cd cmd/server
$ go run main.go
```

The `wasm_exec.js` file came from here:

```text
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ~/Documents/webassembly/assets/
```

Based on [this](https://golangbot.com/webassembly-using-go/) and [this](https://marianogappa.github.io/software/2020/04/01/webassembly-tinygo-cheesse/).
