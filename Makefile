# unfortunately, tinygo doesn't support some of the required features
# cd ./cmd/wasm && tinygo build -o  ../../assets/json.wasm -target wasm . && cd ../..
build:
	cd ./cmd/wasm && GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm && cd ../..
