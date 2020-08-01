build:
	cd ./cmd/wasm && GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm && cd ../..
