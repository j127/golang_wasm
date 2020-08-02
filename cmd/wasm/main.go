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

// TODO: this function updates the DOM, but shouldn't this program also export
// the prettify function so that other JS could use it directly?
func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			result := map[string]interface{}{
				"error": "invalid number of arguments passed",
			}
			return result
		}

		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			result := map[string]interface{}{
				"error": "unable to find the DOM",
			}
			return result
		}

		jsonOutputTextArea := jsDoc.Call("getElementById", "jsonoutput")
		if !jsonOutputTextArea.Truthy() {
			result := map[string]interface{}{
				"error": "unable to get output textarea",
			}
			return result
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
			errStr := fmt.Sprintf("unable to parse JSON. Error\n", err)
			result := map[string]interface{}{
				"error": errStr,
			}
			return result
		}

		jsonOutputTextArea.Set("value", pretty)
		return nil
	})
	return jsonFunc
}

func main() {
	fmt.Println("hello world from wasm")
	js.Global().Set("formatJSON", jsonWrapper())

	// If you don't block here, the Go program will exit.
	<- make(chan bool)
}
