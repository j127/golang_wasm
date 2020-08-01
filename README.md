# Golang Wasm Experiment

```text
$ make build
$ cd cmd/server
$ go run main.go
```

The `wasm_exec.js` file came from here:

```text
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ~/Documents/webassembly/assets/
```

Based on [this](https://golangbot.com/webassembly-using-go/).
