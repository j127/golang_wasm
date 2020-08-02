package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"
)

func prettyJson(input string, indentation int) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", strings.Repeat(" ", indentation))
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "invalid number of arguments passed"
		}
		inputJSON := args[0].String()
		// var indentation int
		// if args[1] != "" {
		// 	indentation = args[1].Int()
		// } else {
		// 	indentation = 4
		// }
		indentation := 4
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJson(inputJSON, indentation)
		if err != nil {
			fmt.Printf("unable to convert to JSON %s\n", err)
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}

func main() {
	fmt.Println("hello world from wasm")
	js.Global().Set("formatJSON", jsonWrapper())

	// If you don't block here, the Go program will exit.
	<- make(chan bool)
}
